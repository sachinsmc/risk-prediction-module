package main

import (
	"log"
	"time"
)


func main()  {
	start := time.Now()
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}
