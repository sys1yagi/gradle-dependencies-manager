package main

import (
	"fmt"
	"os"

	"./tasks"
)

func getTask() string {
	if len(os.Args) <= 1 {
		return ""
	}
	return os.Args[1]
}

func main() {

	switch getTask() {
	case "gen":
		tasks.Generate()
	case "search":
		tasks.Search()
	case "install":
		fmt.Println("install")
	case "versions":
		fmt.Println("versions")

	default:
		fmt.Println("usage:")
		fmt.Println("  gdm gen [dir_name]")
		fmt.Println("  gdm search $query")
		fmt.Println("  gdm install $group_id $artifact_id $version $target [type=$type]")
		fmt.Println("  gdm versions $group_id $artifact_id")
	}

}
