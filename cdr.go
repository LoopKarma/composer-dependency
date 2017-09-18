package main

import (
	"flag"
	"fmt"
	"github.com/LoopKarma/composer-dependency/runner"
	"os"
)

var (
	PATH string = "composer.json" //Path to composer.json file
)

func main() {
	flag.StringVar(&PATH, "p", PATH, "Path to composer.json file")
	flag.Parse()

	deps, err := runner.FindDependenciesInFile(PATH)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if len(deps) == 0 {
		fmt.Println("Result: Nothing to update")
		os.Exit(0)
	}
	result, err := runner.UpdateDependencies(deps)

	if result != true {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
