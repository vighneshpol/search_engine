package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	// "search-engine/backend/data"
	"search-engine/backend/utils"
)

// 🛠 Declare global variable here
var Records []map[string]interface{}

func loadParquetFiles() {
	// Fetch all files in the parquet_files directory
	files, err := filepath.Glob("./parquet_files/*")
	if err != nil {
		log.Fatalf("Failed to list parquet files: %v", err)
	}

	// Debug print: List all files that are found
	fmt.Println("Found files:", files)

	// Loop through each file found
	for _, file := range files {
		fmt.Println("Loading file:", file)

		// Read the data from the current Parquet file
		records, err := utils.ReadParquetFile(file)
		if err != nil {
			log.Printf("Failed to read file %s: %v", file, err)
			continue
		}

		// Debug print: Number of records loaded from this file
		fmt.Printf("Loaded %d records from file %s\n", len(records), file)

		// Append the records to the global Records slice
		Records = append(Records, records...)
	}

	// Final debug print: Total records loaded
	fmt.Printf("Loaded %d total records into memory.\n", len(Records))
}


func main() {
	// Check if parquet_files folder exists
	if _, err := os.Stat("./parquet_files"); os.IsNotExist(err) {
		log.Fatalf("parquet_files folder not found. Please create it and add files.")
	}

	loadParquetFiles()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is running!")
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
	
		query := r.URL.Query().Get("q")
		var results []map[string]interface{}
	
		// Debug print to see query received
		fmt.Printf("Searching for query: %s\n", query)
	
		for _, record := range Records {
			// Print each record being checked
			// fmt.Printf("Checking record: %+v\n", record)
	
			// Search in important fields
			if (fieldContains(record, "Message", query) ||
				fieldContains(record, "MessageRaw", query) ||
				fieldContains(record, "Tag", query) ||
				fieldContains(record, "Sender", query) ||
				fieldContains(record, "Groupings", query) ||
				fieldContains(record, "SeverityString", query)) {
	
				results = append(results, record)
			}
		}
	
		duration := time.Since(start)
	
		response := map[string]interface{}{
			"query":   query,
			"results": len(results),
			"time_ms": duration.Milliseconds(),
			"data":    results, // if you want to see actual matches
		}
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	
	
	
	
	// http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
	// 	query := r.URL.Query().Get("q")
	// 	if query == "" {
	// 		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
	// 		return
	// 	}
	
	// 	start := time.Now()
	
	// 	var results []map[string]interface{}
	// 	for _, record := range Records {
	// 		for _, value := range record {
	// 			if str, ok := value.(string); ok && containsIgnoreCase(str, query) {
	// 				results = append(results, record)
	// 				break
	// 			}
	// 		}
	// 	}
	
	// 	duration := time.Since(start)
	
	// 	w.Header().Set("Content-Type", "application/json")
	// 	fmt.Fprintf(w, `{"query":"%s","results":%d,"time_ms":%d}`, query, len(results), duration.Milliseconds())
	// })
	

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func containsIgnoreCase(str, substr string) bool {
// 	return len(str) >= len(substr) && (stringContainsCaseInsensitive(str, substr))
// }
func fieldContains(record map[string]interface{}, field string, query string) bool {
	if val, ok := record[field]; ok {
		if str, ok := val.(string); ok {
			return containsIgnoreCase(str, query)
		}
	}
	return false
}
func containsIgnoreCase(str, substr string) bool {
    return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

// func stringContainsCaseInsensitive(a, b string) bool {
// 	return stringIndexCaseInsensitive(a, b) != -1
// }

// func stringIndexCaseInsensitive(a, b string) int {
// 	return indexFunc(a, b)
// }

// func indexFunc(a, b string) int {
// 	return len(a) - len(b)
// }
