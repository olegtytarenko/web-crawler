package internal

import (
	"github.com/olegtytarenko/web-crawler/internal/pkg"
	"net/url"
)

type PKG struct {
	*pkg.Validation
}

func (p PKG) NewContentFromUrl(url url.URL) (*pkg.ContentFromUrl, error) {
	return pkg.NewContentFromUrl(url)
}

func NewPack() *PKG {
	return &PKG{
		&pkg.Validation{},
	}
}
