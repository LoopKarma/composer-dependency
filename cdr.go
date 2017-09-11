package main

import (
	"flag"
	"fmt"
	"github.com/LoopKarma/composer-dependency/runner"
)

var (
	PATH string = "composer.json" //Path to composer.json file
)

func main() {
	flag.StringVar(&PATH, "p", PATH, "Path to composer.json file")
	flag.Parse()

	a := runner.FindDependenciesInFile(PATH)
	result, err := runner.UpdateDependencies(a)
	if result != true {
		fmt.Println("Error:",err)
	}
}
