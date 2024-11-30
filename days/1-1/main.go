package main

import (
	"fmt"

	"github.com/robryanx/adventofcode2024/util"
)

func main() {
	strs, err := util.ReadBytes("1", false)
	if err != nil {
		panic(err)
	}

	for _, str := range strs {
		fmt.Println(str)
	}
}
