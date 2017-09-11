package runner

import (
	"io/ioutil"
	"encoding/json"
	"bytes"
	"regexp"
	"os/exec"
)

// Composer.json top level
type Manifest struct {
	Require map[string]string
	Repositories map[string]string
}

type Dependency struct {
	Repo string
}

func FindDependenciesInFile(fileName string) (deps []Dependency) {
	data, err := ioutil.ReadFile(fileName)
	check(err)
	dec := json.NewDecoder(bytes.NewReader(data))
	var d Manifest
	dec.Decode(&d)

	for k, v := range d.Require {
		pattern := "^dev\\-(.+)"
		isMatch,_ := regexp.MatchString(pattern, v)
		if isMatch {
			r, _ := regexp.Compile(pattern)
			_ = r.FindStringSubmatch(v)
			dependency := Dependency{
				k,
			}
			deps = append(deps, dependency)
		}
	}
	return deps
}

func UpdateDependencies(deps []Dependency) (result bool, err error) {
	args := []string{"update"}
	for _, v := range deps {
		args = append(args, v.Repo)
	}
	cmd := exec.Command("composer", args...)
	_, err = cmd.Output()
	result = true
	if err != nil {
		result = false
	}
	return result, err
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}