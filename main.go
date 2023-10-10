package main

import (
	"final_satu/internal/entity"
	"final_satu/internal/repository"
	"fmt"
	"log"
)

func main() {
	var err error
	fmt.Println("Hello, world!")
	db := repository.New()
	err = db.Update(entity.Todos{
		1, "Andi", "s", "s", "s",
	}, 1)
	err, todos := db.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range todos {
		fmt.Println(i)
	}
}
