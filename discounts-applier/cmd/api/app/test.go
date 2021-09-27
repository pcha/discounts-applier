package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fatih/color"
	null "gopkg.in/guregu/null.v4"
)

func IntegrationTest() {
	go Serve()
	for !ping() {
		time.Sleep(time.Second)
	}
	product1 := PresentableProduct{
		SKU:      "000001",
		Name:     "BV Lean leather ankle boots",
		Category: "boots",
		Price: PresentablePrice{
			Original:           89000,
			Final:              62300,
			DiscountPercentage: null.StringFrom("30%"),
			Currency:           "EUR",
		},
	}
	product2 := PresentableProduct{
		SKU:      "000002",
		Name:     "BV Lean leather ankle boots",
		Category: "boots",
		Price: PresentablePrice{
			Original:           99000,
			Final:              69300,
			DiscountPercentage: null.StringFrom("30%"),
			Currency:           "EUR",
		},
	}
	product3 := PresentableProduct{
		SKU:      "000003",
		Name:     "Ashlington leather ankle boots",
		Category: "boots",
		Price: PresentablePrice{
			Original:           71000,
			Final:              49700,
			DiscountPercentage: null.StringFrom("30%"),
			Currency:           "EUR",
		},
	}
	product4 := PresentableProduct{
		SKU:      "000004",
		Name:     "Naima embellished suede sandals",
		Category: "sandals",
		Price: PresentablePrice{
			Original:           79500,
			Final:              79500,
			DiscountPercentage: null.StringFromPtr(nil),
			Currency:           "EUR",
		},
	}
	product5 := PresentableProduct{
		SKU:      "000005",
		Name:     "Nathane leather sneakers",
		Category: "sneakers",
		Price: PresentablePrice{
			Original:           59000,
			Final:              59000,
			DiscountPercentage: null.StringFromPtr(nil),
			Currency:           "EUR",
		},
	}
	testCase("get products without filter", "", []PresentableProduct{
		product1,
		product2,
		product3,
		product4,
		product5,
	})
	testCase("get products filtering by boots category", "?category=boots", []PresentableProduct{
		product1,
		product2,
		product3,
	})
	testCase("get products filtering by price less than 79500", "?priceLessThan=79500", []PresentableProduct{
		product3,
		product4,
		product5,
	})
	testCase("get products filtering by category boots and by price less than 89000", "?category=boots&priceLessThan=89000", []PresentableProduct{
		product1,
		product3,
	})
	testCase("get products filtering by category sandals", "?category=sandals", []PresentableProduct{
		product4,
	})
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

func testCase(testName, query string, expected []PresentableProduct) bool {
	fmt.Println("== TEST", testName, "==")
	code, body, err := get("http://localhost:8080/products" + query)
	if err != nil {
		failTest(err.Error())
		return false
	}
	strBody := string(body)
	if code != http.StatusOK {
		return failTest("status code != 200. Got %v. Body %q", code, strBody)
	}
	expectedJson, err := json.Marshal(expected)
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
	color.Red("FAIL:", fmt.Sprintf(message, args...))
	return false
}

func passTest(message string, args ...interface{}) bool {
	color.Green("PASS:", fmt.Sprintf(message, args...))
	return true
}
