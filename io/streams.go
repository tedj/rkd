package io

import (
	"os"
	"io"
	"io/ioutil"
)

// Struct representing IO for a
// command to be executed on the
// system
type IO struct {
  Stdin, Stdout, Stderr io.Writer
}

// Creates an IO struct geared towards
// ignoring IO at all.
//
//  This is useful when you need to run
// a command ignoring its output, like
// when you need to test that a command
// runs / exists, without caring too much
// about its output.
func NewDevNullIO() IO {
	return IO{ioutil.Discard, ioutil.Discard, ioutil.Discard}
}

// Creates a new IO struct that relies on
// the OS' standard IO streams.
func NewStdIO() IO {
	return IO{os.Stdin, os.Stdout, os.Stderr}
}
