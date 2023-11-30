package workers

import "github.com/t00mas/goimagecrawler/storage"

func SrcSaverWorker(st *storage.StructDB, savechan chan StoreResultsJob) {
	for {
		job := <-savechan
		st.Store(job.Keyword, job.Srcs)
	}
}
