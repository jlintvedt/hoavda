package main

import (
	"path/filepath"

	"github.com/jlintvedt/bargi/util"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	// n := bargi.NewNode()
	// n.Test()

	path, err := filepath.Abs("..\\testdata")
	if err != nil {
		log.Fatal(err)
	}
	sn, err := util.NewSnapLoader(path)
	if err != nil {
		log.Fatal(err)
	}
	sn.Test()
}
