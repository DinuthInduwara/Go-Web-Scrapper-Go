package main

import (
	"Go_Web_Scrapper/db"
	"Go_Web_Scrapper/funcs"
	"Go_Web_Scrapper/logfuncs"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var semaphore = funcs.NewSemaphore(10)

func init() {
	log.Println("[S] Web Crawling Bot Started [S]")
	logfuncs.Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	logfuncs.Logger.SetOutput(os.Stdout)

}

func main() {

	parsedURL, _ := url.Parse("")
	funcs.URL = &funcs.MyURL{URL: parsedURL}

	wg.Add(1)
	workerLoop(parsedURL.String())
	wg.Wait()
}

func workerLoop(Url string) {
	parser, err := url.Parse(Url)
	typeOfFile := funcs.ExtOf(filepath.Ext(Url))
	if parser.Hostname() != funcs.URL.Hostname() || err != nil {
		logfuncs.Logger.Errorln("[£] We Skipped ", Url, " Due to Hostname Miss Matched Or Error = ", err)
		return
	}
	if typeOfFile == funcs.Unknown {
		headers, err := funcs.RertiveDetails(Url)
		if err != nil {
			logfuncs.Logger.Errorln("[£] We Skipped ", Url, " Because Error = ", err)
			return
		}
		if strings.Contains(headers.ContentType, "text/html") {

			resp, err := http.Get(Url)
			if err != nil {
				logfuncs.Logger.Errorln("[£] We Skipped ", Url, "Error = ", err)
				return
			}
			defer resp.Body.Close()
			urls, err := funcs.ParseHtmlTags(resp.Body)
			if err != nil {
				logfuncs.Logger.Errorln("[£] We Skipped ", Url, "Error = ", err)
				return
			}
			for _, item := range *urls {
				go func() {
					wg.Add(1)
					semaphore.Acquire()
					workerLoop(item)
					wg.Done()
					semaphore.Release()
				}()
			}
			return
		}
	}

	db.AddFileToDb(&db.FileDB{
		Url:  Url,
		Type: typeOfFile,
	})

}
