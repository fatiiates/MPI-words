package main

import "fmt"

func main() {
	//counter := Counter{}
	// util := Util{}

	// dat, err := os.ReadFile(latest_file)
	// Check(err)
	// fmt.Print(string(dat))

	// fmt.Println("")
	// fmt.Println("testk")
	// fmt.Println("")

	// rf, err := CountWords(5, latest_file, 0)
	// Check(err)

	// fmt.Println("WORD RF => ", len(rf.words))
	// for i := 0; i < len(rf.uncompleted_words); i++ {
	// 	fmt.Println("word => ", rf.uncompleted_words[i], " isfirst => ", i == 0, " isLast => ", len(rf.uncompleted_words) == i+1)
	// }

	// rf1, err := CountWords(5, latest_file, 1)
	// Check(err)

	// fmt.Println("WORD RF1 => ", len(rf1.words))
	// for i := 0; i < len(rf1.uncompleted_words); i++ {
	// 	fmt.Println("word => ", rf1.uncompleted_words[i], " isfirst => ", i == 0, " isLast => ", len(rf1.uncompleted_words) == i+1)
	// }

	mr, err := Constructor()
	Check(err)
	mr.Map()
	mr.Reduce()
	fmt.Println(mr.results_map)
}
