package main

import (
	"fmt"
	"github.com/wdvxdr1123/go-silk/silk"
	"io/ioutil"
)

func main() {
	silkEncoder := &silk.SilkEncoder{}
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
