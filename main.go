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

// placeholder for XDG config implementation
// see : https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
func xdg() string {
	return filepath.Join(os.Getenv("HOME"), ".config")
}

// placeholder for default config generation
func makeConf(p string, d string) {
}

// read a config file and return a
func readConf(path string) {
	conf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "notabar: reading config")
	}

	r := csv.NewReader(strings.NewReader(string(conf)))
	r.Comment = '#'        // allows for comment lines starting with #
	r.FieldsPerRecord = -1 // allows for variable number of fields
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
		fmt.Printf("%T\n", record)
	}
}

// placeholder send the notification
func notify(s string) {
	cmd := exec.Command("notify-send", s)
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}

func main() {
	config := filepath.Join(xdg(), cDir, cFile)
	fmt.Println(config)
	readConf(config)
	notify("test")
}
