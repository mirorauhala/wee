package main

import (
	"fmt"
	"os"

	"github.com/teris-io/shortid"
)

func setup() *shortid.Shortid {
	return shortid.MustNew(1, shortid.DefaultABC, 2342)
}

func id() (string, error) {

	sid := setup()

	id, err := sid.Generate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return id, nil

}
