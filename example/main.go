package main

import (
	"fmt"

	"github.com/mix3/reisen"
)

func returnError() error {
	return fmt.Errorf("error")
}

func wrapError1() error {
	if err := returnError(); err != nil {
		return reisen.Error(err)
	}
	return nil
}

func wrapError2() error {
	if err := wrapError1(); err != nil {
		return reisen.Error(err)
	}
	return nil
}

func main() {
	if err := returnError(); err != nil {
		fmt.Println(reisen.Error(err))
	}
	if err := wrapError1(); err != nil {
		fmt.Println(reisen.Error(err))
	}
	if err := wrapError2(); err != nil {
		fmt.Println(reisen.Error(err))
	}
}
