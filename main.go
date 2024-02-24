// Copyright © 2024- Luka Ivanović
// This code is licensed under the terms of a 2-clause BSD licence
// (see LICENCE for details)

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const FORMAT = "2006-01-02T150405"

const DAY = time.Hour * 24

func ReadDirs(filename string) ([]string, error) {
	entries, err := os.ReadDir(filename)
	if err != nil {
		return nil, err
	}
	dirs := make([]string, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}
	return dirs, nil
}

func IsOlder(date time.Time, dur time.Duration) bool {
	return date.Add(dur).Before(time.Now())
}

var (
	months = flag.Int64("m", 0, "number of months")
	weeks  = flag.Int64("w", 0, "number of weeks")
	days   = flag.Int64("d", 0, "number of days")
)

func main() {
	flag.Parse()
	duration := (*months)*30 + (*weeks)*7 + (*days)
	args := flag.Args()
	for _, arg := range args {
		path, err := filepath.Abs(arg)
		if err != nil {
			log.Println(err)
			continue
		}
		dirs, err := ReadDirs(path)
		if err != nil {
			log.Println(err)
			continue
		}
		for _, dir := range dirs {
			date, err := time.Parse(FORMAT, dir)
			if err != nil {
				log.Println(err)
				continue
			}
			if IsOlder(date, time.Duration(int64(DAY)*duration)) {
				fmt.Printf("%s\n", dir)
			}
		}
	}
}
