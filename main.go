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
const cFile string = "config"

// find config file
// see : https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
func xdg() string {
	return filepath.Join(os.Getenv("HOME"), ".config")
}

// save default config file
func makeConf(p string, d string) {
}

// read config file from path and return map of string arrays
func readConf(path string) map[int][]string {
	entries := make(map[int][]string)
	conf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "notabar: reading config ")
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
	// default config
	config := filepath.Join(xdg(), cDir, cFile)
	// read arguments
	if len(os.Args) > 1 {
		// use argument config file
		config = filepath.Join(xdg(), cDir, os.Args[1])
	}
	entries := readConf(config)
	note := build(entries)
	notify(note)
}
