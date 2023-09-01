package funcs

import (
	"Go_Web_Scrapper/logfuncs"
	"net/http"
)

func RertiveDetails(url string) (*Headers, error) {
	logfuncs.Logger.Infoln("[:] We Checking ", url)
	resp, err := http.Head(url)
	if err != nil {
		logfuncs.Logger.Errorln("[£] We Cant Checking ", url, " Because Error = ", err)
		return nil, err
	}
	defer resp.Body.Close()
	return &Headers{
		ContentType: resp.Header.Get("Content-Type"),
		// ContentLength: resp.ContentLength,
	}, nil
}
