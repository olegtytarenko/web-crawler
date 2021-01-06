package pkg

import "net/url"

type Validation struct{}

func (v *Validation) IsUrlAddress(urlAddress string) (bool, *url.URL) {
	urlInfo, err := url.Parse(urlAddress)
	if err != nil {
		return false, nil
	}

	if urlInfo.Scheme == "" {
		return false, nil
	}

	// support only http(s) protocols
	if urlInfo.Scheme != "http" && urlInfo.Scheme != "https" {
		return false, nil
	}

	return true, urlInfo
}
