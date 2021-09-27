package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"discounts-applier/internal/discounts/products"
)

type FileStruct struct {
	Products []struct {
		SKU      string `json:"sku"`
		Name     string `json:"name"`
		Category string `json:"category"`
		Price    int    `json:"price"`
	} `json:"products"`
}

func main() {
	file, err := ioutil.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	fileData := new(FileStruct)
	err = json.Unmarshal(file, fileData)
	if err != nil {
		log.Fatal(err)
	}

	mongoProducts := make([]products.Product, len(fileData.Products))
	for i, p := range fileData.Products {
		mongoProducts[i] = products.Product{
			SKU:      p.SKU,
			Name:     p.Name,
			Category: p.Category,
			Price:    p.Price,
		}
	}

	mongoURI := os.Getenv("MONGO_URI")
	rep, err := products.NewRepository(mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	err = rep.Clean()
	if err != nil {
		log.Fatal(err)
	}
	err = rep.Write(mongoProducts)
	if err != nil {
		log.Fatal(err)
	}
}
