package tasks

import "os"

//Doc aaa
type Doc struct {
	ID            string
	G             string
	A             string
	LatestVersion string
	RepositoryID  string
	P             string
	Timestamp     int
	VersionCount  int
	V             string
	Text          []string
	Ex            []string
}

//Response aaa
type Response struct {
	NumFound int
	Start    int
	Docs     []Doc
}

//SearchResult aaa
type SearchResult struct {
	Response Response
}

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
