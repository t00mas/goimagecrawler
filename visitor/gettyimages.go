package visitor

import (
	"fmt"
	"io"
	"net/http"

	"github.com/t00mas/goimagecrawler/htmlparsers"
)

const GETTYIMAGES_BASE_URL = "https://www.gettyimages.es/fotos/%s?assettype=image&excludenudity=false&alloweduse=availableforalluses&family=creative&phrase=hotel&sort=mostpopular&page=%d"

type GettyImagesVisitor struct {
	keyword string
	page    int
}

func (giv *GettyImagesVisitor) SetKeyword(keyword string) {
	giv.keyword = keyword
}

func (giv *GettyImagesVisitor) SetPage(page int) {
	giv.page = page
}

func (giv *GettyImagesVisitor) Target() string {
	return fmt.Sprintf(GETTYIMAGES_BASE_URL, giv.keyword, giv.page)
}

func (giv *GettyImagesVisitor) Visit() ([]string, error) {
	res, err := http.Get(giv.Target())
	if err != nil {
		return []string{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, err
	}
	defer res.Body.Close()
	bodyStr := string(body)
	srcs := htmlparsers.FindAllImgSrcVals(bodyStr)

	return srcs, nil
}
