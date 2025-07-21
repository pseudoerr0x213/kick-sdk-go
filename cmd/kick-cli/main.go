package main

import (
	"context"
	"fmt"

	"github.com/pseudoerr/kick-sdk-go/pkg/kick"
)

func main() {

	// the main point is to provide easy to use abstractions
	// initialize the backrground context
	ctx := context.Background()
	// simply pass your id and secret token to the constructor
	client := kick.NewClient("yourID", "yourSecret")
	// call a desired method
	categories, _ := client.GetCategories(ctx)
	fmt.Println(categories)
}
