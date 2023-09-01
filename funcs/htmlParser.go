package funcs

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)

/*
	func getReader(htmlPath *string) (*goquery.Document, error) {
		file, err := os.Open(*htmlPath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		return goquery.NewDocumentFromReader(file)
	}
*/
func getAnchorTags(selection *goquery.Selection, urls *[]string) {
	excludedValues := []string{"Name", "Last Modified", "Size", "Parent Directory"}
	linkText := selection.Text()
	if !contains(linkText, &excludedValues) {
		href, exists := selection.Attr("href")
		if exists {
			*urls = append(*urls, reValidateUrl(href))
		}
	}
}

func getImgVideoAudioTags(selection *goquery.Selection, urls *[]string) {
	src, exists := selection.Attr("src")
	if exists {
		*urls = append(*urls, reValidateUrl(src))
	}
}

func ParseHtmlTags(reader io.Reader) (*[]string, error) {
	/*	document, err := getReader(htmlFile)
	 */
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	var urls []string

	document.Find("a").Each(func(index int, selection *goquery.Selection) {
		getAnchorTags(selection, &urls)
	})
	document.Find("video source").Each(func(index int, selection *goquery.Selection) {
		getImgVideoAudioTags(selection, &urls)
	})
	document.Find("audio source").Each(func(index int, selection *goquery.Selection) {
		getImgVideoAudioTags(selection, &urls)
	})
	document.Find("img").Each(func(index int, selection *goquery.Selection) {
		getImgVideoAudioTags(selection, &urls)
	})

	return &urls, nil
}
