package main

import (
	"context"
	"fmt"

	"github.com/SsSJKK/search/pkg/search"
)

func main() {

	res := search.All(context.Background(), "a", []string{"./data/test.txt", "./data/test2.txt "})

	r, ok := <-res
	if !ok {
		fmt.Println("error ok =>", ok)
	}

	for _, r := range r {

		fmt.Println("---------------")
		fmt.Println("res.Phrase) => ", r.Phrase)
		fmt.Println("res.Line) => ", r.Line)
		fmt.Println("res.LineNum) => ", r.LineNum)
		fmt.Println("res.ColNum) => ", r.ColNum)
		fmt.Println("---------------")
	}

}
