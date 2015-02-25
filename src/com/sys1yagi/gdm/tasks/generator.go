package tasks

import (
	"fmt"
	"io/ioutil"
)

//Generate dependencies.gradle.
func Generate() {
	target := getArg(2)
	file := "dependencies.gradle"
	if target != "" {
		file = target + "/" + file
	}
	if Exists(file) {
		fmt.Println("already exists. " + file)
		return
	}
	data := `dependencies {
    compile fileTree(dir: 'libs', include: ['*.jar'])
    compile 'com.android.support:appcompat-v7:21.0.3'
}
`
	err := ioutil.WriteFile(file, []byte(data), 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println("generate succeed. ->" + file)
}
