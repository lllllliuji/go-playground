package myfile

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)
type KeyValue struct {
	X string 
	Y string
}
func FileTest() {
	filename := "myfiletest"
	ofile, _ := os.Create(filename)
	enc := json.NewEncoder(ofile)
	for i := 0; i < 10; i ++ {
		kv := KeyValue {
			X: strconv.Itoa(rand.Int()),
			Y: strconv.Itoa(rand.Int()),
		}
		enc.Encode(&kv)
	}
	ofile.Close()
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
