package main

import (
	"log"

	"github.com/jlintvedt/bargi/util"
)

func main() {
	// n := bargi.NewNode()
	// n.Test()
	sn, err := util.NewSnapLoader("H:\\Repos\\go\\src\\github.com\\jlintvedt\\testdata")
	if err != nil {
		log.Fatal(err)
	}
	sn.Test()
}
