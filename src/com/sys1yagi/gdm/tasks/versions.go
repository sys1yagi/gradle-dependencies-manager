package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Versions aaa
func Versions(printer func(int, Doc)) {
	groupID := getArg(2)
	artifactID := getArg(3)

	if groupID == "" {
		fmt.Println("groupID is empty.")
		fmt.Println("USAGE: gdm version $groupID $artifactID")
		return
	}
	if artifactID == "" {
		fmt.Println("artifactID is empty.")
		fmt.Println("USAGE: gdm version $groupID $artifactID")
		return
	}

	fmt.Println("search... " + groupID + "," + artifactID)
	url := "http://search.maven.org/solrsearch/select?q=g:%22" + groupID + "%22+AND+a:%22" + artifactID + "%22&core=gav&rows=20&wt=json"
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
