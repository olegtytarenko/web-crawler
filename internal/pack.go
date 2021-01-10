package internal

import (
	"github.com/olegtytarenko/web-crawler/internal/pkg"
	"net/url"
)

type PKG struct {
	*pkg.Validation
	*pkg.ContentFromUrl
}

func (p *PKG) NewContentFromUrl(url *url.URL) error {
	contentFromUrl, err := pkg.NewContentFromUrl(url)
	if err != nil {
		return err
	}
	p.ContentFromUrl = contentFromUrl
	return nil
}

func NewPack() *PKG {
	return &PKG{
		&pkg.Validation{},
		&pkg.ContentFromUrl{},
	}
}
