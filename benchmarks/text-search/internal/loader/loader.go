package loader

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// LoadWikipediaFromJSON loads Wikipedia articles from JSON file
// The file should have the format: [{"title": "...", "text": "..."}, ...]
func LoadWikipediaFromJSON(filename string) ([]Article, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	var articles []struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}

	if err := json.Unmarshal(data, &articles); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	result := make([]Article, len(articles))
	for i, a := range articles {
		result[i] = Article{
			ID:      int64(i + 1),
			Title:   a.Title,
			Content: a.Text,
		}
	}

	return result, nil
}

// ConvertParquetToJSON demonstrates how to convert parquet to JSON
// Run this with Python: python3 -c "import pandas as pd; df = pd.read_parquet('data/wikipedia.parquet'); df.to_json('data/wikipedia.json', orient='records')"
func ConvertParquetNote() {
	fmt.Println("To use real Wikipedia data:")
	fmt.Println("1. Install Python dependencies: pip install pandas pyarrow")
	fmt.Println("2. Convert: python3 -c \"import pandas as pd; df = pd.read_parquet('data/wikipedia.parquet'); df.to_json('data/wikipedia.json', orient='records')\"")
	fmt.Println("3. The JSON file will be used automatically if present")
}

// GetContent combines multiple text fields for search indexing
func GetSearchContent(article Article) string {
	return article.Title + " " + article.Content
}

// GetSearchContentForProduct combines product fields
func GetSearchContentForProduct(p Product) string {
	return p.Name + " " + p.Description + " " + p.Category + " " + p.Brand
}

// GetSearchContentForQA combines Q&A fields
func GetSearchContentForQA(q QAPair) string {
	return q.Title + " " + q.Body + " " + q.Answer + " " + q.Tags
}

// NormalizeQuery cleans up search queries
func NormalizeQuery(q string) string {
	q = strings.ToLower(q)
	q = strings.TrimSpace(q)
	return q
}
