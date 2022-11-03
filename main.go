package main

import (
	"fmt"
	"io/ioutil"

	"github.com/eric2788/go-silk/silk"
)

func main() {
	silkEncoder := &silk.Encoder{}
	err := silkEncoder.Init("cache", "codec")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadFile("test.mp3")
	if err != nil {
		fmt.Println(err)
	}
	_, err = silkEncoder.EncodeToSilk(data, "test", true)
	if err != nil {
		fmt.Println(err)
	}
}
