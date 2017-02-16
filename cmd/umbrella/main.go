package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

var dir = flag.String("path", "", "")

func abs(name string) (string, error) {
	if name == "" {
		return os.Getwd()
	}

	if !path.IsAbs(name) {
		abs, err := filepath.Abs(name)
		if err != nil {
			return "", err
		}
		name = abs
	}

	_, err := os.Stat(name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func init() {
	flag.Parse()

	abs, err := abs(*dir)
	if err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
	*dir = abs
}

func main() {
	fmt.Println(*dir)
}
