package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ContentFromUrl struct {
	url.URL
	http.Response
}

func loadContent(url url.URL) (*http.Response, error) {
	hRes, err := http.Get(url.String())
	if err != nil {
		return &http.Response{StatusCode: http.StatusBadGateway}, err
	}

	if hRes.StatusCode != http.StatusOK {
		return &http.Response{StatusCode: http.StatusBadGateway}, errors.New(fmt.Sprintf("Status code is not correct %d", hRes.StatusCode))
	}

	return hRes, nil
}

func (cfu ContentFromUrl) ReadResponse() (string, error) {
	//goland:noinspection GoUnhandledErrorResult
	defer cfu.Response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(cfu.Response.Body)
	if err != nil {
		return "", err
	}
	rd := string(bodyBytes)
	return rd, nil
}

func (cfu ContentFromUrl) GetAnchorLinks() ([]ContentFromUrl, error) {
	listContentFromUrls := []ContentFromUrl{}

	return listContentFromUrls, nil
}

func NewContentFromUrl(url *url.URL) (*ContentFromUrl, error) {
	contentFromUrl := &ContentFromUrl{URL: *url}
	res, err := loadContent(contentFromUrl.URL)
	if err != nil {
		return nil, err
	}
	return &ContentFromUrl{
		URL:      *url,
		Response: *res,
	}, nil
}
