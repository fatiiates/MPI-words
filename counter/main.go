package main

import (
	"fmt"
	"os"
)

func main() {
	//counter := Counter{}
	// util := Util{}
	latest_file := "2022_25_05-17_42_11.txt" //util.GetLastCreatedFile()

	dat, err := os.ReadFile(latest_file)
    Check(err)
    fmt.Print(string(dat))

	fmt.Println("")
	fmt.Println("testk")
	fmt.Println("")

	f, err := os.Open(latest_file)
	Check(err)
	inf, err := f.Stat()
	fmt.Println("size => ", inf.Size())
	Check(err)

	fmt.Println("")
	fmt.Println("testk")
	fmt.Println("")

	tmp := ""

	b1 := make([]byte, 8)
	for i := 0; i < 6; i++ {
		n1, err := f.Read(b1)
		Check(err)
		index := 0
		for i := 0; i < 8; i++ {
			if b1[i] == ' ' {
				index = i+1
				tmp += string(b1[:i])
				fmt.Println("kelime ",i," => ", tmp)
				tmp = ""
			}
		}
		tmp += string(b1[index:])
		fmt.Println("son kelime  => ", tmp)
		fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	}

	fmt.Println("")
	fmt.Println("testk")
	fmt.Println("")

	o2, err := f.Seek(6, 0)
	Check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	Check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

}
