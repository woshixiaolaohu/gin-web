package cmd

import "fmt"

type server interface {
	ListenAndServe() error
}

func Execute() {
	//s := initServer("8888")
	fmt.Println("Execute")
}
