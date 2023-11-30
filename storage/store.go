package storage

type StructDB struct {
	// implements a mutex lock for concurrent storage into object
	Storage map[string][]string
}

func (sdb *StructDB) Store(keyword string, urls []string) {
	// check / release  of StructDB
	_, ok := sdb.Storage[keyword]
	if !ok {
		sdb.Storage[keyword] = urls
	} else {
		sdb.Storage[keyword] = append(sdb.Storage[keyword], urls...)
	}
}
