package main

import (
	"fmt"

	"github.com/nutchy/go-mock-test/pkg/foo"
	"github.com/nutchy/go-mock-test/pkg/repository"
)

func main() {
	fmt.Println("hello world!")

	repo := repository.NewRepository()

	foo := foo.NewFoo(repo)

	resp, err := foo.DoSometingSimple()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v \n", resp)
}
