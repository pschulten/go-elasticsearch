// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/indices/create-index.asciidoc#L123>
//
// --------------------------------------------------------------------------------
// PUT /test
// {
//     "settings" : {
//         "number_of_shards" : 1
//     },
//     "mappings" : {
//         "properties" : {
//             "field1" : { "type" : "text" }
//         }
//     }
// }
// --------------------------------------------------------------------------------

func Test_indices_create_index_dfef545b1e2c247bafd1347e8e807ac1(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:dfef545b1e2c247bafd1347e8e807ac1[]
	res, err := es.Indices.Create(
		"test",
		es.Indices.Create.WithBody(strings.NewReader(`{
		  "settings": {
		    "number_of_shards": 1
		  },
		  "mappings": {
		    "properties": {
		      "field1": {
		        "type": "text"
		      }
		    }
		  }
		}`)),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:dfef545b1e2c247bafd1347e8e807ac1[]
}
