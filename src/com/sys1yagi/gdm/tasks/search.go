package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
	}
	for i, doc := range result.Response.Docs {
		printer(i, doc)
	}
}
