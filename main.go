package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const cDir string = "notabar"
const defaultConfig string = "config"

// find config file
func xdg() string {
	return filepath.Join(os.Getenv("HOME"), ".config")
}

// wrap parseConf to read paths
func parsePath(path string) map[int][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Printf("Failed opening config file: %v", err)
	}
	defer f.Close()
	return parseConf(f)
}

// read config file and return map of string arrays
func parseConf(f *os.File) map[int][]string {
	entries := make(map[int][]string)
	conf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("Failed reading config file: %v", err)
	}
	r := csv.NewReader(strings.NewReader(string(conf)))
	r.Comment = '#'        // allows for comment lines starting with #
	r.FieldsPerRecord = -1 // allows for variable number of fields
	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		entries[i] = record
		i++
	}
	return entries
}

// take txt string array and return finalized string
func txt(entry []string) string {
	s := ""
	switch entry[1] {
	case "\\n":
		s = s + "\n"
	default:
		s = s + entry[1]
	}
	return s
}

// take cmd string array and return finalized string
func cmd(entry []string) string {
	c := entry[1]
	a := entry[2:]
	out, err := exec.Command(c, a...).Output()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	s := strings.TrimSuffix(string(out), "\n")
	return s
}

// take map of string arrays and return notification string
func build(entries map[int][]string) string {
	s := ""
	for e := 0; e < len(entries); e++ { // e = entry
		switch entries[e][0] {
		case "txt":
			s = s + txt(entries[e])
		case "cmd":
			s = s + cmd(entries[e])
		default:
			// error
			fmt.Println("Error: Unkown config entry type")
		}
	}
	return s
}

// run the notification string
func notify(s string) {
	s = strings.TrimSuffix(s, "\n")
	c := exec.Command("notify-send", s)
	err := c.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}

func main() {
	// check if stdin was passed
	if _, err := os.Stdin.Stat(); err != nil {
		// if stdin was not passed
		arg := filepath.Join(xdg(), cDir, defaultConfig)
		args := os.Args[1:]
		if len(args) > 1 {
			// if there was an argument passed
			arg = filepath.Join(xdg(), cDir, os.Args[1])
		}
		// open the config file
		entries := parsePath(arg)
		note := build(entries)
		notify(note)
	} else {
		// stdin is available
		entries := parseConf(os.Stdin)
		note := build(entries)
		notify(note)
	}
}
