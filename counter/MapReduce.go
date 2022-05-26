package main

import (
	"errors"
	"os"
	"strconv"
)

type MapReduce[T any] struct {
	send_buffer    *[]T
	recv_buffer    *[]T
	process_number int
}

func Constructor[T comparable](send_buffer *[]T, recv_buffer *[]T, process_number int) MapReduce[T] {
	mr := MapReduce[T]{}

	err := mr.validate()
	if err != nil {
		panic(err)
	}

	return mr
}

func (mr *MapReduce[T]) validate() error {
	if len(os.Args) < 1 {
		return errors.New("CLI arguments must have 1 arguments at least")
	}

	num, err := strconv.Atoi(os.Args[1])
	Check(err)

	mr.process_number = num

	if mr.process_number < 100 {
		return errors.New("WORLD_SIZE can be 100 at max")
	}
	if mr.process_number < 100 {
		return errors.New("WORLD_SIZE can't be lowest from 1")
	}

	return nil
}

func (mr *MapReduce[T]) Map(size int) {

}
