package workers

import "sync"

type GettyImagesVisitJob struct {
	Keyword string
	Page    int
	Wg      *sync.WaitGroup
}

type StoreResultsJob struct {
	Keyword string
	Srcs    []string
}
