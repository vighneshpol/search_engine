// package utils

// import (
// 	//"os"

// 	"github.com/xitongsys/parquet-go/reader"
// 	"github.com/xitongsys/parquet-go-source/local"
// 	"search-engine/backend/model"
// )

package utils

import (
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
	"log"
	
)

func ReadParquetFile(filePath string) ([]map[string]interface{}, error) {
	fr, err := local.NewLocalFileReader(filePath)
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return nil, err
	}
	defer fr.Close()

	pr, err := reader.NewParquetReader(fr, nil, int64(4))
	if err != nil {
		log.Printf("Failed to create parquet reader: %v", err)
		return nil, err
	}
	defer pr.ReadStop()

	num := int(pr.GetNumRows())
	result := make([]map[string]interface{}, 0, num)

	for i := 0; i < num; i++ {
		row := make(map[string]interface{})
		if err := pr.Read(&row); err != nil {
			log.Printf("Failed to read row: %v", err)
			break
		}
		result = append(result, row)
	}

	return result, nil
}

