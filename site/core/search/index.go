package search

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve"
)

//
// Connect
//
func Connect(indexPath string) (bleve.Index, error) {
	index, err := bleve.Open(indexPath)
	if err != nil {
		index, err = bleve.New(indexPath, bleve.NewIndexMapping())
		if err != nil {
			return nil, err
		}
	}
	return index, nil
}

type IndexedPost struct {
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
}

//
// RunIndex
//
func RunIndex(blevepath string, posts []*IndexedPost) error {
	index, err := Connect(blevepath)
	defer index.Close()
	if err != nil {
		return fmt.Errorf("bleve.Open: err: %v", err)
	}
	for _, p := range posts {
		if err := index.Index(p.Title, p); err != nil {
			log.Println("index post: %v err: %v", p.Title, err)
			continue
		}
	}
	return nil
}
