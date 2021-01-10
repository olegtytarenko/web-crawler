package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ContentFromUrl struct {
	*url.URL
	*http.Response
}

func loadContent(url url.URL) (*http.Response, error) {
	hRes, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	if hRes.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("E1: Status code is not correct %d", hRes.StatusCode))
	}

	return hRes, nil
}

func (cfu *ContentFromUrl) ReadResponse() (string, error) {
	if cfu.Response == nil {
		return "", errors.New("HTTP Response is empty")
	}
	defer cfu.Response.Body.Close()
	if cfu.Response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("E2: Status code is not correct %d", cfu.Response.StatusCode))
	}

	bytesReadResponses, err := ioutil.ReadAll(cfu.Response.Body)

	if err != nil {
		return "", err
	}
	rd := string(bytesReadResponses)
	return rd, nil
}

func (cfu *ContentFromUrl) GetAnchorLinks() ([]ContentFromUrl, error) {
	listContentFromUrls := []ContentFromUrl{}

	return listContentFromUrls, nil
}

func NewContentFromUrl(url *url.URL) (*ContentFromUrl, error) {
	res, err := loadContent(*url)
	if err != nil {
		return nil, err
	}
	return &ContentFromUrl{
		URL:      url,
		Response: res,
	}, nil
}
