package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"zzzz/myfile"
)

func main() {
	myfile.FileTest()
	mp := make(map[int]int)
	for i := 0; i < 10; i++ {
		mp[i] += 10 - i
	}
	for k, v := range mp {
		fmt.Println(k, v)
	}
	ofile, _ := os.Create("tmp.txt")

	for k, v := range mp {
		fmt.Fprintf(ofile, "%v %v\n", k, v)
	}
	ofile.Close()
	switch x := 0; x {
	case 0:
		fmt.Println("here")
	case 1:
		fmt.Println("there")
	}
	file, _ := ioutil.TempFile("", "mr-out")
	fmt.Println(file.Name())
	for k, v := range mp {
		fmt.Fprintf(file, "%v %v\n", k, v)
	}
	file.Close()
	os.Rename(file.Name(), "./mr-out1")
}
