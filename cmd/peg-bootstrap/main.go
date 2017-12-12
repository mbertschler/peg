// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/pointlander/peg/tree"
)

func main() {
	buffer, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	p := &Peg{Tree: tree.New(false, false, false), Buffer: string(buffer), Pretty: true}
	p.Init()
	if err := p.Parse(); err != nil {
		log.Fatal(err)
	}
	p.Execute()
	p.Compile("boot.peg.go", os.Args, os.Stdout)
}