package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	//"github.com/elastic/go-elasticsearch/v7/internal/version"
)

var es *elasticsearch.Client

func main() {
	// Initialize the Elasticsearch client

	// Command to genrate the password '/usr/share/elasticsearch/bin/elasticsearch-reset-password -u elastic'
	//Elasticsearch password = fM3dfFJ6Vn6R8GGJo=Ee
	cfg := elasticsearch.Config{
		Addresses: []string{"https://192.168.0.133:9200"},
		Username:  "elastic",
		Password:  "fM3dfFJ6Vn6R8GGJo=Ee",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	var err error
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)

	var your_index_name = "with_shrad_in_windows_restaurant_index"
	fmt.Println("IndexName : " + your_index_name)

	// Create the index
	createIndex(your_index_name)

	// Add a document
	routingValue := "florida_routing" //--> used for shrad
	addDocument(your_index_name, "1", routingValue, map[string]interface{}{
		"state": "Florida",
		"restaurant": map[string]interface{}{
			"name":     "Babbu Restaurant",
			"address":  "789 Main St",
			"location": map[string]float64{"lat": 34.0522, "lon": -118.2437},
		},
		"menu": []map[string]interface{}{
			{"name": "Pizza"},
			{"name": "French Fries"},
		},
		"menu_item": map[string]interface{}{
			"name":  "Large Pizza",
			"price": 59.02,
		},
	})

	//var routingValueArray []string
	//routingValueArray = append(routingValueArray, routingValue)

	// Search documents
	searchResults := searchDocuments(your_index_name, "Florida", append([]string{}, routingValue))
	fmt.Println("Search Results:")
	fmt.Println(searchResults)

	// Delete a document
	deleteDocument(your_index_name, "1")
}

func createIndex(indexName string) {

	indexBody := `
	{
	    "settings": {
			"number_of_shards": 3,
			"number_of_replicas": 2
		},
		"mappings": {
		"properties": {
		    "state": {
				"type": "keyword"
		    },
			"routing_field": {
				"type": "keyword"
		    },
		    "restaurant": {
			"properties": {
			    "name": { "type": "text" },
			    "address": { "type": "text" },
			    "location": { "type": "geo_point" }
			}
		    },
		    "menu": {
			"properties": {
			    "name": { "type": "text" }
			}
		    },
		    "menu_item": {
			"properties": {
			    "name": { "type": "text" },
			    "price": { "type": "float" }
			}
		    }
		}
	    }
	}
	`

	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(indexBody),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error creating the index: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error in response: %s", res.String())
	}

	log.Println("Index created successfully")
}

func addDocument(indexName, docID string, routingValue string, doc map[string]interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		log.Fatalf("Error encoding document: %s", err)
	}

	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: docID,
		Body:       &buf,
		Refresh:    "true",
		Routing:    routingValue,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	log.Println("Document indexed successfully")
}

func searchDocuments(indexName, query string, routingValue []string) string {
	var buf bytes.Buffer
	queryBody := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"state": query,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(queryBody); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	req := esapi.SearchRequest{
		Index:   []string{indexName},
		Body:    &buf,
		Routing: routingValue, //--> this is optional use it with shrad
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error searching documents: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	var searchResult map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	response, err := json.MarshalIndent(searchResult, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling search response: %s", err)
	}

	return string(response)
}

func deleteDocument(indexName, docID string) {
	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: docID,
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error deleting document: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	log.Println("Document deleted successfully")
}
