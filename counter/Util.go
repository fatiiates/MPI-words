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

	"gopkg.in/yaml.v3"
)

const (
	GENERATION_RESULTS_PATH = "../generator/results"
	COUNTING_RESULTS_PATH   = "./results"
)

var (
	BUFF_SIZE = 64
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

type YamlFormat struct {
	Result         *map[string]int `yaml:"words"`
	TotalWordCount int             `yaml:"total_word_count"`
}

func discoverFiles() ([]File, error) {
	fs, err := ioutil.ReadDir(GENERATION_RESULTS_PATH)
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
		return nil, errors.New("there is no generated file in " + GENERATION_RESULTS_PATH + ". Please give a specific file with COUNT_FILE variable")
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

	f, err := os.Open(filename)
	Check(err)
	defer f.Close()

	_, err = f.Seek(int64(offset), 0)
	Check(err)

	var res Result = Result{
		make(map[string]int),
		[]Word{},
	}

	var counter int = size

	buff := make([]byte, BUFF_SIZE)
	var tmp string
	var leftIndex int
	var leftControl bool
	for counter != 0 {
		bytes, err := f.Read(buff)
		Check(err)
		if counter < BUFF_SIZE {
			bytes = counter
		}
		leftIndex = 0
		for i := 0; i < bytes; i++ {
			if buff[i] == ' ' {
				tmp += string(buff[leftIndex:i])
				if tmp == "" {
					leftIndex = i + 1
					continue
				}
				if len(res.uncompleted_words) == 0 {
					AppendToArray(&res.uncompleted_words, Word{
						tmp,
						leftControl,
						true,
					})
				} else {
					res.words[tmp]++
				}
				leftIndex = i + 1
				tmp = ""
			}
		}
		counter -= bytes
		tmp += string(buff[leftIndex:bytes])
		leftControl = leftIndex != 0
	}

	if tmp != "" {
		AppendToArray(&res.uncompleted_words, Word{
			tmp,
			len(res.uncompleted_words) > 0 || len(res.words) > 0 || leftControl,
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

func WriteResultsToFile(c *Counter) {
	dt := time.Now()
	filename := dt.Format("2006_02_01-15_04_05") + ".yaml"

	yaml_result := YamlFormat{
		&c.results_map,
		c.total_word_count,
	}
	yamlData, err := yaml.Marshal(&yaml_result)
	Check(err)
	err = os.WriteFile(strings.Join([]string{COUNTING_RESULTS_PATH, filename}, "/"), yamlData, 0644)
	Check(err)
	fmt.Println("Words written as counted to file", filename)
}

func WorkingTime(start_time time.Time) {
	out := time.Since(start_time)
	fmt.Println("That counting took", out.Seconds(), "seconds.")
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
