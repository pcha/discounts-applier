package main

import (
	"fmt"
	"log"
	"os"

	"discounts-applier/cmd/api/app"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("the program expects fo one and only one arguments. Try with help")
	}

	switch os.Args[1] {
	case "serve":
		app.Serve()
	case "test":
		app.IntegrationTest()
	case "help":
		fmt.Println(getHelp())
	default:
		log.Fatal("unkown option\n" + getHelp())
	}
}

func getHelp() string {
	return "The available arguments are: \n" +
		"help \t show this help \n" +
		"serve \t serve the api \n" +
		"test \t serve the api on background and run the integrations tens against it \n"
}
