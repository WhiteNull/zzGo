package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("change to sit start")
	b, err := ioutil.ReadFile("/etc/woda/GlobalConfigure_sit.yaml")
	if err != nil {
		fmt.Print(err)
	}
	err = ioutil.WriteFile("/etc/woda/GlobalConfigure.yaml", b, 0644)
	if err != nil {
		fmt.Print("fail:")
		fmt.Print(err)
	} else {
		fmt.Print("change to alpha success")
	}
	fmt.Println("change to sit start")
}
