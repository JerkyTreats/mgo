package main

import (
	"github.com/globalsign/mgo"
)

func main() {
	session, err := mgo.Dial("mongodb")
	if err != nil {
		panic(err)
	}
	defer session.Close()
}
