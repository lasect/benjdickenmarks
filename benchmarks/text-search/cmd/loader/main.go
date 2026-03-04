package main

import (
	"fmt"
	"os"

	"github.com/benjdickenmarks/text-search/internal/loader"
)

func main() {
	dataset := os.Getenv("DATASET")
	size := 100000

	fmt.Printf("Generating %d %s records...\n", size, dataset)

	switch dataset {
	case "wikipedia":
		data := loader.GenerateWikipediaArticles(size)
		fmt.Printf("Generated %d Wikipedia articles\n", len(data))
	case "ecommerce":
		data := loader.GenerateProducts(size)
		fmt.Printf("Generated %d products\n", len(data))
	case "qa":
		data := loader.GenerateQA(size)
		fmt.Printf("Generated %d Q&A pairs\n", len(data))
	default:
		data := loader.GenerateProducts(size)
		fmt.Printf("Generated %d products (default)\n", len(data))
	}
}
