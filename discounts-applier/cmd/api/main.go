package main

import (
	"fmt"
	"log"
	"os"

	"discounts-applier/cmd/api/app"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("the program expects fo one and only one arguments. Try with help")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "serve":
		err := app.Serve()
		log.Fatal(err)
	case "test":
		res := app.IntegrationTest()
		if !res {
			os.Exit(1)
		}
	case "help":
		fmt.Println(getHelp())
	default:
		fmt.Println("unkown option\n" + getHelp())
		os.Exit(1)
	}
}

func getHelp() string {
	return "The available arguments are: \n" +
		"help \t show this help \n" +
		"serve \t serve the api \n" +
		"test \t serve the api on background and run the integrations tens against it \n"
}
