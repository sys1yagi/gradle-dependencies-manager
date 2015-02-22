package tasks

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strings"
)

const N = "\n"

func Install() {
	usage := "USAGE: gdm install $dependency $target [type=$type]"
	dependency := getArg(2)
	target := getArg(3)
	types := getArg(4)
	if dependency == "" {
		fmt.Println("ERROR: dependency must not be empty.")
		fmt.Println(usage)
		return
	}
	if target == "" {
		fmt.Println("ERROR: target must not be empty.")
		fmt.Println(usage)
		return
	}
	if types == "" {
		types = "compile"
	}

	file := "dependencies.gradle"
	if target != "" {
		file = target + "/" + file
	}
	if !Exists(file) {
		fmt.Println("target not found. " + file)
		return
	}
	//open dependencies.gradle
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	dependencies := string(bytes)
	newDependencies := list.New()
	var pos *list.Element
	for _, line := range strings.Split(dependencies, N) {
		line = strings.Trim(line, " ")
		switch {
		case strings.HasPrefix(line, types):
			pos = newDependencies.PushBack("  " + line)
		case strings.HasPrefix(line, "dependencies"):
			newDependencies.PushBack(line)
		case strings.HasPrefix(line, "}"):
		case line == "":
		default:
			newDependencies.PushBack("  " + line)
		}
	}

	if pos == nil {
		newDependencies.PushBack("  " + types + " '" + dependency + "'")
	} else {
		newDependencies.InsertAfter("  "+types+" '"+dependency+"'", pos)
	}

	newDependencies.PushBack("}")

	data := ""
	for v := newDependencies.Front(); v != nil; v = v.Next() {
		data += v.Value.(string) + N
	}

	err = ioutil.WriteFile(file, []byte(data), 0755)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("install success.")
}
