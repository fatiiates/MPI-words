package main

import (
	"errors"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"
)

var (
	RESULTS_PATH = "../generator/results"
)

type Util struct{}

type File struct {
	name         string
	created_date time.Time
}

func discoverFiles() ([]File, error) {
	fs, err := ioutil.ReadDir(RESULTS_PATH)
	if err != nil {
		log.Fatal(err)
	}

	files := []File{}

	for _, file := range fs {
		splitted_name := strings.Split(file.Name(), ".")
		r, _ := regexp.Compile(`^(\d{4}_\d{2}_\d{2}-\d{2}_\d{2}_\d{2}\.txt)$`)

		if file.IsDir() ||
			splitted_name[len(splitted_name)-1] != "txt" ||
			!r.MatchString(file.Name()) {
			continue
		}
		file.ModTime()
		files = append(files, File{
			file.Name(),
			file.ModTime(),
		})
	}

	if len(files) == 0 {
		return nil, errors.New("There is no generated file in " + RESULTS_PATH)
	}

	return files, nil
}

func GetLastCreatedFile() string {
	files, err := discoverFiles()
	if err != nil {
		panic(err)
	}
	var latest_file File
	for _, f := range files {
		if latest_file.created_date.Before(f.created_date) {
			latest_file = f
		}
	}

	return latest_file.name
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
