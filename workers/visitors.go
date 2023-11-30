package workers

import (
	"fmt"

	"github.com/t00mas/goimagecrawler/visitor"
)

func GettyImagesVisitorWorker(wi int, visitchan chan GettyImagesVisitJob, savechan chan StoreResultsJob) {
	visitor := visitor.GettyImagesVisitor{}
	for {
		visitorJob := <-visitchan
		visitor.SetKeyword(visitorJob.Keyword)
		visitor.SetPage(visitorJob.Page)
		srcs, err := visitor.Visit()
		if err != nil {
			fmt.Printf("error visiting keyword: %s error: %+v", visitorJob.Keyword, err)
			continue
		}
		saveJob := StoreResultsJob{
			Keyword: visitorJob.Keyword,
			Srcs:    srcs,
		}
		savechan <- saveJob
		visitorJob.Wg.Done()
	}
}
