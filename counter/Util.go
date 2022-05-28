package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	RESULTS_PATH = "../generator/results"
	BUFF_SIZE    = 8
)

type Util struct{}

type File struct {
	name         string
	created_date time.Time
}

type Word struct {
	word            string
	isLeftComplete  bool
	isRightComplete bool
}

type Result struct {
	words             map[string]int
	uncompleted_words []Word
}

func discoverFiles() ([]File, error) {
	fs, err := ioutil.ReadDir(RESULTS_PATH)
	if err != nil {
		Check(err)
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
		return nil, errors.New("There is no generated file in " + RESULTS_PATH + ". Please give a specific file with COUNT_FILE variable")
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

func CountWords(size int, filename string, process_number int, offset int) (*Result, error) {

	var res Result = Result{
		make(map[string]int),
		[]Word{},
	}

	var counter int = size

	f, err := os.Open(filename)
	Check(err)
	_, err = f.Seek(int64(offset), 0)
	Check(err)

	buff := make([]byte, BUFF_SIZE)
	var tmp string

	for counter != 0 {
		bytes, err := f.Read(buff)
		Check(err)

		if counter < BUFF_SIZE {
			bytes = counter
		}

		index := 0
		for i := 0; i < bytes; i++ {
			if buff[i] == ' ' {
				tmp += string(buff[:i])
				if len(res.uncompleted_words) == 0 {
					AppendToArray(&res.uncompleted_words, Word{
						tmp,
						index != 0,
						true,
					})
				} else {
					res.words[tmp]++
				}
				tmp = ""
				index = i + 1
			}
		}
		counter -= bytes
		tmp += string(buff[index:bytes])
	}

	if tmp != "" {
		AppendToArray(&res.uncompleted_words, Word{
			tmp,
			len(res.uncompleted_words) > 0 || len(res.words) > 0,
			false,
		})
	}

	if len(res.uncompleted_words) == 0 && len(res.words) == 0 {
		return nil, errors.New("there is no value in process " + fmt.Sprint(process_number))
	}

	return &res, nil
}

func AppendToArray[T comparable](arr *[]T, val T) {
	*arr = append(*arr, val)
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
