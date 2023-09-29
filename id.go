package main

import (
	"github.com/teris-io/shortid"
)

var sid *shortid.Shortid

func SetupShortId() {
	sid = shortid.MustNew(1, shortid.DefaultABC, 2342)
}
