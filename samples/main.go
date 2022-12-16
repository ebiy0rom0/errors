package main

import (
	"fmt"

	myerr "github.com/ebiy0rom0/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error   { return wrap1() }
func wrap1() error { return wrap2() }
func wrap2() error { return wrap3() }
func wrap3() error { return wrap4() }
func wrap4() error { return wrap5() }
func wrap5() error { return wrap6() }
func wrap6() error { return myerr.New("error") }
