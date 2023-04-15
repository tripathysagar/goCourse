package main

import "fmt"

type StringStorer interface {
	GetAll() ([]string, error)
	Put(string) error
}

type MongoDBStringStore struct {
	// mongoDB connection stuff
}

func (m MongoDBStringStore) GetAll() ([]string, error) {
	return []string{"a", "b", "c", "d"}, nil
}

func (m MongoDBStringStore) Put(String string) error {
	fmt.Println("store the sstring into the mongoDb storage")
	return nil
}

type MySQLStringStore struct {
	// mysqlconnection stuff
}

func (m MySQLStringStore) GetAll() ([]string, error) {
	return []string{"a", "b", "c", "d", "e"}, nil
}

func (m MySQLStringStore) Put(String string) error {
	fmt.Println("store the sstring into the mysql storage")
	return nil
}

type ApiServer struct {
	stringStore StringStorer
}

func main() {
	//	apiServer := ApiServer{
	//		stringStore: MongoDBStringStore{},
	//	}
	//	strings, err := apiServer.stringStore.GetAll()
	//	if err != nil {
	//		panic("got error during db reading !!!!")
	//	}
	//	fmt.Println("data after insertion", strings)
	//	if err := apiServer.stringStore.Put("e"); err != nil {
	//		panic("got error during db writing !!!!")
	//	}
	//

	apiServer := ApiServer{
		stringStore: MySQLStringStore{},
	}

	strings, err := apiServer.stringStore.GetAll()
	if err != nil {
		panic("got error during db reading !!!!")
	}
	fmt.Println("data after insertion", strings)

	if err := apiServer.stringStore.Put("e"); err != nil {
		panic("got error during db writing !!!!")
	}

}
