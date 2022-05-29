package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Counter struct {
	process_number   int
	filename         string
	filesize         int
	wait_group       *sync.WaitGroup
	recv_buffer      map[int]*Result
	results_map      map[string]int
	total_word_count int
	mutex            sync.Mutex
}

func CounterConstructor() (*Counter, error) {
	c := Counter{
		0,
		"",
		0,
		&sync.WaitGroup{},
		make(map[int]*Result),
		make(map[string]int),
		0,
		sync.Mutex{},
	}

	err := c.validate()
	Check(err)

	return &c, nil
}

func (c *Counter) validate() error {
	if len(os.Args) < 2 {
		return errors.New("CLI arguments must have 1 arguments at least")
	}

	num, err := strconv.Atoi(os.Args[1])
	Check(err)

	c.process_number = num

	if c.process_number > 100 {
		return errors.New("WORLD_SIZE can be 100 at max")
	}
	if c.process_number < 1 {
		return errors.New("WORLD_SIZE cannot be lowest from 1")
	}

	if len(os.Args) > 2 {
		c.filename = os.Args[2]
	} else {
		c.filename = strings.Join([]string{GENERATION_RESULTS_PATH, GetLastCreatedFile()}, "/")
	}

	if c.filename == "" {
		return errors.New("filename cannot be empty")
	}

	file, err := os.Stat(c.filename)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}
	if file.IsDir() {
		return errors.New(c.filename + " is not file")
	}

	if file.Size() == 0 {
		return errors.New(c.filename + " is not contain anything")
	}

	if file.Size() < int64(c.process_number) {
		return errors.New(c.filename + " size cannot be lower from WORLD_SIZE")
	}

	if file.Size() > 110000000 {
		return errors.New(c.filename + " size can be 110M at max")
	}

	c.filesize = int(file.Size())

	return nil
}

func (c *Counter) Count(){
	mr := MRConstructor(c)
	mr.Map()
	mr.Reduce()
	WriteResultsToFile(c)
}