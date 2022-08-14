package main

import (
	"fmt"
	"log"

	"github.com/moltenpl85/storage/v2/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it is restored", restoredFile)
}
