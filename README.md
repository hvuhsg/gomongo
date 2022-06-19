# gomongo
Mongo like database writen in GO  

----
[![Go](https://github.com/hvuhsg/gomongo/actions/workflows/go.yml/badge.svg)](https://github.com/hvuhsg/gomongo/actions/workflows/go.yml)  


### NOTE
This project is a work in progress


### Install
```shell
go get github.com/hvuhsg/gomongo
```

### Simple usage examples
```go
package main

import (
	"fmt"

	"github.com/hvuhsg/gomongo/engine"
	"github.com/hvuhsg/gomongo/indexing"
	"github.com/hvuhsg/gomongo/storage/ram"
	"github.com/hvuhsg/gomongo/validation"
)

func main() {
	validator := validation.New()
	indexor := indexing.New()
	ramStorage := ram.New()
	dbEngine := engine.New(validator, indexor, ramStorage)

	dbEngine.CreateDatabase("db")
	dbEngine.CreateCollection("db", "col")

	documents := []map[string]any{{"a": 5}, {"a": 2.5}}
	dbEngine.Insert("db", "col", documents)

	filter := map[string]any{"a": map[string]any{"$gt": 2.5}}
	result, _ := dbEngine.Find("db", "col", filter)

	fmt.Println(result) // -> [map[a:5]]
}
```
