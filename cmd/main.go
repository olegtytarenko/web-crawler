package main

import (
	"flag"
	"fmt"
	"github.com/olegtytarenko/web-crawler/internal"
	"log"
)

var urlByParse string

func init() {
	flag.Parse()
	urlByParse = flag.Arg(0)
}

func main() {
	packages := internal.NewPack()

	testResult, Url := packages.Validation.IsUrlAddress(urlByParse)

	if testResult == false {
		log.Fatal("Url argument is empty or wrong")
		return
	}
	ContentFromUrl, err := packages.NewContentFromUrl(*Url)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println(ContentFromUrl.ReadResponse())

}
