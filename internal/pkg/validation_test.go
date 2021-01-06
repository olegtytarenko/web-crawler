package pkg

import (
	"fmt"
	"testing"
)

func TestIsUrlAddress(t *testing.T) {
	testDataUrl := map[string]bool{
		"https://golang.org/pkg/testing/#example_B_RunParallel": true,
		"golang.org":                   false,
		"localhost":                    false,
		"ftp://golang.org/pkg/testing": false,
		"https://golang.org":           true,
		"https://golang.org/":          true,
		"http://golang.org/":           true,
		"//golang.org/":                false,
		"test.com":                     false,
	}

	Validation := &Validation{}

	for urlAddressByTest, assertResult := range testDataUrl {
		testResult, _ := Validation.IsUrlAddress(urlAddressByTest)
		if assertResult != testResult {
			fmt.Println(fmt.Sprintf("Test case is faile with address %s", urlAddressByTest))
			t.Failed()
			continue
		}
	}
}
