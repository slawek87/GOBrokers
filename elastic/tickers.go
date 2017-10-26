package elastic

import (
	"time"
	"fmt"
)

// function ticker indexing. Each hour we are indexing new records.
func IndexDocumentsTicker() {
	ticker := time.NewTicker(time.Minute * 60)
	go func() {
		for t := range ticker.C {
			fmt.Println("Indexing Documents at:", t)
			IndexDocumentsTicker()
		}
	}()
}

