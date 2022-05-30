package main

type MapReduce struct {
	counter *Counter
}

func MRConstructor(c *Counter) *MapReduce {
	mr := MapReduce{
		c,
	}

	return &mr
}

func (mr *MapReduce) Map() {
	mod := mr.counter.filesize % mr.counter.process_number
	bytes_per_proc := mr.counter.filesize / mr.counter.process_number
	offset := 0
	for i := 0; i < mr.counter.process_number; i++ {
		mr.counter.wait_group.Add(1)
		i := i
		size := bytes_per_proc
		os := offset
		if mod > 0 {
			size++
			mod--
		}
		go func() {
			defer mr.counter.wait_group.Done()
			res, err := CountWords(size, mr.counter.filename, i, os)
			Check(err)
			mr.counter.mutex.Lock()
			mr.counter.recv_buffer[i] = res
			mr.counter.mutex.Unlock()
		}()
		offset += size
	}
}

func (mr *MapReduce) Reduce() {
	mr.counter.wait_group.Wait()
	mr.mergeUncompletedWords()
	for i := 0; i < mr.counter.process_number; i++ {
		for k, v := range mr.counter.recv_buffer[i].words {
			mr.increaseKeyAndTotalWord(k, v)
		}
	}
}

func (mr *MapReduce) mergeUncompletedWords() {

	tmp := ""
	for i := 0; i < mr.counter.process_number; i++ {
		for _, v := range mr.counter.recv_buffer[i].uncompleted_words {
			if tmp != "" && v.isLeftComplete {
				mr.increaseKeyAndTotalWord(tmp, 1)
				tmp = ""
			}
			if v.isRightComplete {
				mr.increaseKeyAndTotalWord(tmp+v.word, 1)
				tmp = ""
				continue
			}
			tmp += v.word
		}
	}
	if tmp != "" {
		mr.increaseKeyAndTotalWord(tmp, 1)
	}
}

func (mr *MapReduce) increaseKeyAndTotalWord(key string, v int) {
	mr.counter.results_map[key] += v
	mr.counter.total_word_count += v
}
