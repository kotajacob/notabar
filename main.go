package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

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
func build(entries map[int][]string) (string, string) {
	s := ""
	u := "normal"
	for e := 0; e < len(entries); e++ { // e = entry
		switch entries[e][0] {
		case "txt":
			s = s + txt(entries[e])
		case "cmd":
			s = s + cmd(entries[e])
		case "urgency":
			u = entries[e][1]
		default:
			// error
			fmt.Println("Error: Unkown config entry type")
		}
	}
	return s, u
}

// run the notification string
func notify(s string, u string) {
	s = strings.TrimSuffix(s, "\n")
	c := exec.Command("notify-send", "-u", u, s)
	err := c.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}

func main() {
	entries := parseConf(os.Stdin)
	note, urgency := build(entries)
	notify(note, urgency)
}
