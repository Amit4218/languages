package main

import (
	"fmt"
	"io"
	http "net/http"
)

const URL = "https://amit4218.fun"

func main() {
	fmt.Println("_________WebRequests_________")

	res, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Content: %v", string(data))

}
