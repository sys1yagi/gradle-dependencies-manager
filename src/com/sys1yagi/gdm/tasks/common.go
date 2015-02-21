package tasks

import "os"

//Exists check file exists.
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func getArg(index int) string {
	if len(os.Args) <= index {
		return ""
	}
	return os.Args[index]
}
