package pkg

import (
	"net/http"
	"net/url"
	"testing"
)

func TestNewContentFromUrl(t *testing.T) {
	// Url (string): Bool (is error url)
	DataProviderTestUrls := map[string]bool{
		"https://google.com": false,
		"https://unset.url":  true,
	}

	for UrlAddress, isNotError := range DataProviderTestUrls {
		urlSchema, err := url.Parse(UrlAddress)
		if err != nil {
			if isNotError == false {
				t.Logf("Check valid url %s: %s", UrlAddress, err)
				t.Fail()
				continue
			}
		}
		ContentFromUrl, err := NewContentFromUrl(urlSchema)
		if err != nil {
			if isNotError == false {
				t.Logf("Error NewContentFromUrl: %s", err)
				t.Fail()
				continue
			} else {
				continue
			}
		}

		if ContentFromUrl == nil {
			if isNotError == false {
				continue
			}
			t.Logf("Error Object test ContentFromIUrl is empty")
			t.Fail()
			continue
		}

		if ContentFromUrl.Response.StatusCode != http.StatusOK && isNotError == false {
			t.Logf("Read response Status code is %d", ContentFromUrl.Response.StatusCode)
			t.Fail()
		}
		_, err = ContentFromUrl.ReadResponse()

		if err != nil && isNotError == false {
			t.Logf("Read Response has error: %s", err)
			t.Fail()
		}
	}
}
