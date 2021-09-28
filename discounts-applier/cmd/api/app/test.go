package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"discounts-applier/cmd/api/app/router"

	null "gopkg.in/guregu/null.v4"
)

type testCase struct {
	name     string
	query    string
	expected []router.PresentableProduct
}

func IntegrationTest() bool {
	startServer()
	product1 := router.PresentableProduct{
		SKU:      "000001",
		Name:     "BV Lean leather ankle boots",
		Category: "boots",
		Price: router.PresentablePrice{
			Original:           89000,
			Final:              62300,
			DiscountPercentage: null.StringFrom("30%"),
			Currency:           "EUR",
		},
	}
	product2 := router.PresentableProduct{
		SKU:      "000002",
		Name:     "BV Lean leather ankle boots",
		Category: "boots",
		Price: router.PresentablePrice{
			Original:           99000,
			Final:              69300,
			DiscountPercentage: null.StringFrom("30%"),
			Currency:           "EUR",
		},
	}
	product3 := router.PresentableProduct{
		SKU:      "000003",
		Name:     "Ashlington leather ankle boots",
		Category: "boots",
		Price: router.PresentablePrice{
			Original:           71000,
			Final:              49700,
			DiscountPercentage: null.StringFrom("30%"),
			Currency:           "EUR",
		},
	}
	product4 := router.PresentableProduct{
		SKU:      "000004",
		Name:     "Naima embellished suede sandals",
		Category: "sandals",
		Price: router.PresentablePrice{
			Original:           79500,
			Final:              79500,
			DiscountPercentage: null.StringFromPtr(nil),
			Currency:           "EUR",
		},
	}
	product5 := router.PresentableProduct{
		SKU:      "000005",
		Name:     "Nathane leather sneakers",
		Category: "sneakers",
		Price: router.PresentablePrice{
			Original:           59000,
			Final:              59000,
			DiscountPercentage: null.StringFromPtr(nil),
			Currency:           "EUR",
		},
	}

	cases := []testCase{
		{
			"get products without filter",
			"",
			[]router.PresentableProduct{
				product1,
				product2,
				product3,
				product4,
				product5,
			},
		},
		{
			"get products filtering by boots category",
			"?category=boots",
			[]router.PresentableProduct{
				product1,
				product2,
				product3,
			},
		},
		{
			"get products filtering by price less than 79500",
			"?priceLessThan=79500",
			[]router.PresentableProduct{
				product3,
				product4,
				product5,
			},
		},
		{
			"get products filtering by category boots and by price less than 89000",
			"?category=boots&priceLessThan=89000",
			[]router.PresentableProduct{
				product1,
				product3,
			},
		},
		{
			"get products filtering by category sandals",
			"?category=sandals",
			[]router.PresentableProduct{
				product4,
			},
		},
	}
	totalTests := len(cases)
	okTests := 0
	for _, tc := range cases {
		if runCase(tc) {
			okTests++
		}
	}
	fmt.Printf("Succesfull tests: %v/%v", okTests, totalTests)
	return okTests == totalTests
}

func startServer() {
	go func() {
		err := Serve()
		log.Fatal(err)
	}()
	for !ping() {
		time.Sleep(time.Second)
	}
}

func ping() bool {
	code, _, err := get("http://localhost:8080/ping")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return code == http.StatusOK
}

func get(url string) (int, []byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, body, nil
}

func runCase(tc testCase) bool {
	fmt.Println("== TEST", tc.name, "==")
	code, body, err := get("http://localhost:8080/products" + tc.query)
	if err != nil {
		failTest(err.Error())
		return false
	}
	strBody := string(body)
	if code != http.StatusOK {
		return failTest("status code != 200. Got %v. Body %q", code, strBody)
	}
	expectedJson, err := json.Marshal(tc.expected)
	if err != nil {
		return failTest(err.Error())
	}

	strExpected := expectedJson
	if strBody != string(strExpected) {
		return failTest("%q != %q", strBody, string(strExpected))
	}
	return passTest(strBody)
}

func failTest(message string, args ...interface{}) bool {
	fmt.Println("FAIL:", fmt.Sprintf(message, args...))
	return false
}

func passTest(message string, args ...interface{}) bool {
	fmt.Println("PASS:", fmt.Sprintf(message, args...))
	return true
}
