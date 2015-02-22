package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
		tasks.Search(func(i int, doc tasks.Doc) {
			fmt.Println(strconv.Itoa(i) + ".  " + doc.G + ":" + doc.A + ":" + doc.LatestVersion)

			fmt.Print("  ")
			for j := range doc.Text {
				text := doc.Text[j]
				if strings.HasPrefix(text, ".") {
					fmt.Print(text + ", ")
				}
			}
			fmt.Println()
		})
	case "installSearch":
		tasks.Search(func(i int, doc tasks.Doc) {
			fmt.Println(doc.G + ":" + doc.A + ":" + doc.LatestVersion)
		})
	case "install":
		tasks.Install()
	case "versions":
		fmt.Println("versions")

	default:
		fmt.Println("usage:")
		fmt.Println("  gdm gen [dir_name]")
		fmt.Println("  gdm search $query")
		fmt.Println("  gdm install $dependency $target [type=$type]")
		fmt.Println("  gdm versions $group_id $artifact_id")
	}

}
