package runner

import (
	"io/ioutil"
	"encoding/json"
	"bytes"
	"regexp"
	"os/exec"
	"bufio"
	"fmt"
)

// Composer.json top level
type Manifest struct {
	Require map[string]string
	Repositories map[string]string
}

type Dependency struct {
	Repo string
}

func FindDependenciesInFile(fileName string) (deps []Dependency, err error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return deps, err
	}

	fmt.Printf("Composer file '%s' found\n", fileName)

	dec := json.NewDecoder(bytes.NewReader(data))
	var d Manifest
	dec.Decode(&d)

	for k, v := range d.Require {
		pattern := "^dev\\-(.+)"
		isMatch,_ := regexp.MatchString(pattern, v)
		if isMatch {
			fmt.Printf("Found dependency [library:'%s' - version:'%s']\n", k, v)
			r, _ := regexp.Compile(pattern)
			_ = r.FindStringSubmatch(v)
			dependency := Dependency{
				k,
			}
			deps = append(deps, dependency)
		}
	}
	return deps, err
}

func UpdateDependencies(deps []Dependency) (result bool, err error) {
	//prepare command
	args := []string{"update", "--no-interaction", "--no-suggest"}
	for _, v := range deps {
		args = append(args, v.Repo)
	}
	cmd := exec.Command("composer", args...)

	fmt.Printf("Running command: %s\n", cmd.Args)

	//composer update uses Stderr for output
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Command finished with error: %v", err)
		return false, err
	}
	return true, err
}
