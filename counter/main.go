package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("<<<==========COUNTING START=========>>>")
	start := time.Now()

	counter, err := CounterConstructor()
	Check(err)
	counter.Count()

	WorkingTime(start)
	fmt.Println("<<<===========COUNTING END==========>>>")
}
