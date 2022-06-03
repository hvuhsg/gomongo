package filtering_test

import (
	"testing"

	"github.com/hvuhsg/gomongo/filtering"
)

type testCase struct {
	filter   map[string]interface{}
	document map[string]interface{}
}

func TestFilter(t *testing.T) {
	trueResults := []testCase{
		{
			filter:   map[string]interface{}{"$and": map[string]interface{}{"name": map[string]interface{}{"$exists": true}}},
			document: map[string]interface{}{"name": "mosh"},
		},
		{
			filter:   map[string]interface{}{"$and": map[string]interface{}{"name": map[string]interface{}{"$exists": true}, "age": map[string]interface{}{"$eq": 8}}},
			document: map[string]interface{}{"name": "mosh", "age": 8},
		},
		{
			filter:   map[string]interface{}{"$or": map[string]interface{}{"name": map[string]interface{}{"$exists": true}, "age": map[string]interface{}{"$eq": 8}}},
			document: map[string]interface{}{"name": "mosh", "age": 3},
		},
		{
			filter:   map[string]interface{}{"$nor": map[string]interface{}{"name": map[string]interface{}{"$exists": true}, "age": map[string]interface{}{"$eq": 8}}},
			document: map[string]interface{}{"age": 3},
		},
		{
			filter:   map[string]interface{}{"$not": map[string]interface{}{"name": map[string]interface{}{"$exists": true}}},
			document: map[string]interface{}{"age": 3},
		},
		{
			filter:   map[string]interface{}{"age": map[string]interface{}{"$gt": 18}},
			document: map[string]interface{}{"age": 29},
		},
		{
			filter:   map[string]interface{}{"age": map[string]interface{}{"$gte": 18}},
			document: map[string]interface{}{"age": 18},
		},
		{
			filter:   map[string]interface{}{"age": map[string]interface{}{"$lt": 18.5}},
			document: map[string]interface{}{"age": 15},
		},
		{
			filter:   map[string]interface{}{"age": map[string]interface{}{"$lte": 17}},
			document: map[string]interface{}{"age": 17.0},
		},
	}

	falseResults := []testCase{
		{
			filter:   map[string]interface{}{"$and": map[string]interface{}{"name": map[string]interface{}{"$exists": true}}},
			document: map[string]interface{}{},
		},
		{
			filter:   map[string]interface{}{"$and": map[string]interface{}{"name": map[string]interface{}{"$exists": true}, "age": map[string]interface{}{"$eq": 8}}},
			document: map[string]interface{}{"age": 8},
		},
		{
			filter:   map[string]interface{}{"$or": map[string]interface{}{"name": map[string]interface{}{"$exists": true}, "age": map[string]interface{}{"$eq": 8}}},
			document: map[string]interface{}{"age": 3},
		},
		{
			filter:   map[string]interface{}{"$nor": map[string]interface{}{"name": map[string]interface{}{"$exists": true}, "age": map[string]interface{}{"$eq": 8}}},
			document: map[string]interface{}{"age": 8},
		},
		{
			filter:   map[string]interface{}{"$not": map[string]interface{}{"name": map[string]interface{}{"$exists": true}}},
			document: map[string]interface{}{"name": "josh"},
		},
		{
			filter:   map[string]interface{}{"age": map[string]interface{}{"$lt": 18}},
			document: map[string]interface{}{"age": 29},
		},
		{
			filter:   map[string]interface{}{"age": map[string]interface{}{"$gte": 18}},
			document: map[string]interface{}{"age": 17},
		},
		{
			filter:   map[string]interface{}{"age": map[string]interface{}{"$gt": 18}},
			document: map[string]interface{}{"age": 14},
		},
		{
			filter:   map[string]interface{}{"age": map[string]interface{}{"$lte": 17}},
			document: map[string]interface{}{"age": 18},
		},
	}

	for _, testCase := range trueResults {
		t.Run("true result expected", func(t *testing.T) {
			result, err := filtering.Filter(testCase.filter, testCase.document)
			if err != nil {
				t.Fatal(err)
			}

			if !result {
				t.Errorf("filter %v for document %v expects to return true", testCase.filter, testCase.document)
			}
		})
	}

	for _, testCase := range falseResults {
		t.Run("false result expected", func(t *testing.T) {
			result, err := filtering.Filter(testCase.filter, testCase.document)
			if err != nil {
				t.Fatal(err)
			}

			if result {
				t.Errorf("filter %v for document %v expects to return false", testCase.filter, testCase.document)
			}
		})
	}
}
