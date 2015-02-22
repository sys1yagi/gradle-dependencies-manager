package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

//Search dependencies.gradle.
func Search(printer func(int, Doc)) {
	query := getArg(2)
	if query == "" {
		fmt.Println("query is empty.")
		fmt.Println("USAGE: gdm search $query")
		return
	}
	fmt.Println("search... " + query)

	url := "http://search.maven.org/solrsearch/select?q=" + query + "&rows=20&wt=json"
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	var result SearchResult
	err := json.Unmarshal(byteArray, &result)
	if err != nil {
		panic(err)
		return
	}
	for i, doc := range result.Response.Docs {
		printer(i, doc)
	}
}
