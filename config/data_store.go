package config

import model "challenges_7/module/model/book"

type DataStore struct {
	DataBook map[uint64]model.Book
}

func ConnectDataStore() DataStore {
	// init map
	bookData := make(map[uint64]model.Book)

	return DataStore{
		DataBook: bookData,
	}
}
