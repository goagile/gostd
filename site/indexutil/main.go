package main

import (
	"io/ioutil"
	"log"
	"path"
	"trygo/site/core"
	"trygo/site/core/search"
)

var (
	// runPostsIndex bool
	pathToPostsDir = "../posts"
	blevePath      = "../example.bleve"
)

func main() {
	// Load Posts
	files, err := ioutil.ReadDir(pathToPostsDir)
	if err != nil {
		log.Fatalln("read posts directory err:", err)
	}
	posts := []core.Post{}
	for _, f := range files {
		log.Println(f.Name())
		post, _, err := core.LoadPost(path.Join(pathToPostsDir, f.Name()))
		if err != nil {
			log.Println("fail to load post:", f.Name(), err)
			continue
		}
		posts = append(posts, post)
		log.Printf(post.Title)
	}

	//
	poststoindex := []*search.IndexedPost{}
	for _, p := range posts {
		poststoindex = append(poststoindex, &search.IndexedPost{
			Title:    p.Title,
			Markdown: p.Markdown,
		})
	}
	if err := search.RunIndex(blevePath, poststoindex); err != nil {
		log.Fatalln("run index err:", err)
	}
}
