package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"

	"github.com/t00mas/goimagecrawler/storage"
	"github.com/t00mas/goimagecrawler/workers"
)

var (
	keywords   = flag.String("keywords", "", "keywords to scrape images for")
	numPages   = flag.Int("numPages", 1, "number of pages to scrape")
	numWorkers = flag.Int("numWorkers", 1, "number of parallel workers to use")
)

func main() {
	flag.Parse()
	kws := strings.Split(*keywords, ",")

	visitChan := make(chan workers.GettyImagesVisitJob)
	saveChan := make(chan workers.StoreResultsJob)
	for wi := 0; wi < *numWorkers; wi++ {
		go workers.GettyImagesVisitorWorker(wi, visitChan, saveChan)
	}

	store := &storage.StructDB{
		Storage: make(map[string][]string),
	}
	go workers.SrcSaverWorker(store, saveChan)

	wg := new(sync.WaitGroup)
	for _, kw := range kws {
		for p := 0; p < *numPages; p++ {
			wg.Add(1)
			job := workers.GettyImagesVisitJob{
				Keyword: kw,
				Page:    p,
				Wg:      wg,
			}
			visitChan <- job
		}
	}

	wg.Wait()
	for k, v := range store.Storage {
		fmt.Printf("\nkeyword: %s srcs: %d", k, len(v))
	}
}
