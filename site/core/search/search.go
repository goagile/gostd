package search

import (
	"fmt"

	"github.com/blevesearch/bleve"
)

//
// FindPostShortInfo
//
func FindPostShortInfo(blevepath, searchstring string) ([]map[string]interface{}, error) {
	result := []map[string]interface{}{}

	index, err := Connect(blevepath)
	defer index.Close()
	if err != nil {
		return result, fmt.Errorf("bleve.Open: err: %v", err)
	}

	ch := bleve.NewTermQuery(searchstring)
	mq := bleve.NewMatchPhraseQuery(searchstring)
	rq := bleve.NewRegexpQuery(searchstring)
	qsq := bleve.NewQueryStringQuery(searchstring)
	q := bleve.NewDisjunctionQuery(ch, mq, rq, qsq)

	search := bleve.NewSearchRequest(q)
	search.Fields = []string{"title"}
	searchResults, err := index.Search(search)
	if err != nil {
		return result, fmt.Errorf("bleve.Search: err: %v", err)
	}
	// fmt.Println(searchResults)
	// fmt.Println(searchResults.Hits[0].ID)
	// fmt.Println(searchResults.Hits[0].Fields)

	for _, hit := range searchResults.Hits {
		result = append(result, hit.Fields)
	}

	return result, nil
}
