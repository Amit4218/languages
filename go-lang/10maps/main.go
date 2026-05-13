package main

import "fmt"

func main() {
	fmt.Println("__________Maps____________")

	var languages = make(map[string]string)

	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["JS"] = "JavaScript"

	fmt.Println("All Languages: ", languages)
	fmt.Println("Py stands for: ", languages["PY"])

	delete(languages, "RB")
	fmt.Println("All Languages after delete: ", languages)

}
