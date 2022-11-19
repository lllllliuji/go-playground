package myfile

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"math/rand"
	"os"
	"strconv"
)

type KeyValue struct {
	X string
	Y string
}

func FileTest() {
	tmpfile, _ := os.CreateTemp("", "tmpfile")
	// tmpfile, _ := ioutil.TempFile("", "tmpfile")
	enc := json.NewEncoder(tmpfile)
	for i := 0; i < 10; i++ {
		kv := KeyValue{
			X: strconv.Itoa(rand.Int()),
			Y: strconv.Itoa(rand.Int()),
		}
		enc.Encode(&kv)
	}
	tmpfile.Close()
	filename := "myfiletest"
	os.Rename(tmpfile.Name(), "./"+filename)
	infile, _ := os.Open(filename)
	dec := json.NewDecoder(infile)
	for {
		var kv KeyValue
		if err := dec.Decode(&kv); err != nil {
			break
		}
		fmt.Printf("%v %v\n", kv.X, kv.Y)
	}
}
