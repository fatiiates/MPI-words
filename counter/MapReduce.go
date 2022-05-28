package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"sync"
)

type MapReduce struct {
	process_number int
	filename       string
	filesize       int
	wait_group     *sync.WaitGroup
	recv_buffer    map[int]*Result
	results_map    map[string]int
	mutex          sync.Mutex
}

func Constructor() (*MapReduce, error) {
	mr := MapReduce{
		0,
		"",
		0,
		&sync.WaitGroup{},
		make(map[int]*Result),
		make(map[string]int),
		sync.Mutex{},
	}

	err := mr.validate()
	Check(err)

	return &mr, nil
}

func (mr *MapReduce) validate() error {
	if len(os.Args) < 1 {
		return errors.New("CLI arguments must have 1 arguments at least")
	}

	num, err := strconv.Atoi(os.Args[1])
	Check(err)

	mr.process_number = num

	if mr.process_number > 100 {
		return errors.New("WORLD_SIZE can be 100 at max")
	}
	if mr.process_number < 1 {
		return errors.New("WORLD_SIZE cannot be lowest from 1")
	}

	if len(os.Args) > 2 {
		mr.filename = os.Args[2]
	} else {
		mr.filename = strings.Join([]string{RESULTS_PATH, GetLastCreatedFile()}, "/")
	}

	if mr.filename == "" {
		return errors.New("filename cannot be empty")
	}

	file, err := os.Stat(mr.filename)
	if errors.Is(err, os.ErrNotExist) {
		return err
	}
	if file.IsDir() {
		return errors.New(mr.filename + " is not file")
	}

	if file.Size() == 0 {
		return errors.New(mr.filename + " is not contain anything")
	}

	if file.Size() < int64(mr.process_number) {
		return errors.New(mr.filename + " size cannot be lower from WORLD_SIZE")
	}

	if file.Size() > 2000000 {
		return errors.New(mr.filename + " size can be 2M at max")
	}

	mr.filesize = int(file.Size())

	return nil
}

func (mr *MapReduce) Map() {
	mod := mr.filesize % mr.process_number
	bytes_per_proc := mr.filesize / mr.process_number
	offset := 0
	for i := 0; i < mr.process_number; i++ {
		mr.wait_group.Add(1)
		i := i
		size := bytes_per_proc
		os := offset
		if mod > 0 {
			size++
			mod--
		}
		go func() {
			defer mr.wait_group.Done()
			res, err := CountWords(size, mr.filename, i, os)
			Check(err)
			mr.mutex.Lock()
			mr.recv_buffer[i] = res
			mr.mutex.Unlock()
		}()
		offset += size
	}
}

func (mr *MapReduce) Reduce() {
	mr.wait_group.Wait()
	mr.mergeUncompletedWords()
	for i := 0; i < mr.process_number; i++ {
		for k, v := range mr.recv_buffer[i].words {
			mr.results_map[k] += v
		}
	}
}

func (mr *MapReduce) mergeUncompletedWords() {

	tmp := ""
	for i := 0; i < mr.process_number; i++ {
		for _, v := range mr.recv_buffer[i].uncompleted_words {
			if tmp != "" && v.isLeftComplete {
				mr.results_map[tmp]++
				tmp = v.word
				continue
			}
			if v.isRightComplete {
				mr.results_map[tmp+v.word]++
				tmp = ""
				continue
			}
			tmp += v.word
		}
	}
	if tmp != "" {
		mr.results_map[tmp]++
	}
}
