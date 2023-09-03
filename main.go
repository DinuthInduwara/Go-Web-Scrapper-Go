package main

import (
	"Go_Web_Scrapper/db"
	"Go_Web_Scrapper/funcs"
	"Go_Web_Scrapper/logfuncs"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

// var semaphore = funcs.NewSemaphore(100)

func init() {
	log.Println("[S] Web Crawling Bot Started [S]")

	logfuncs.LogError.SetOutput(funcs.LogFile)

	logfuncs.Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	logfuncs.Logger.SetOutput(os.Stdout)

}

func main() {

	parsedURL, _ := url.Parse("")
	funcs.URL = &funcs.MyURL{URL: parsedURL}
	go DatabaseHandler()

	wg.Add(1)
	workerLoop(parsedURL.String())
	wg.Wait()
}

func workerLoop(Url string) {
	parser, err := url.Parse(Url)
	typeOfFile := funcs.ExtOf(filepath.Ext(Url))
	if parser.Hostname() != funcs.URL.Hostname() || err != nil {
		logfuncs.LogError.Errorln("[£] We Skipped ", Url, " Due to Hostname Miss Matched Or Error = ", err)
		return
	}
	if typeOfFile == funcs.Unknown {
		headers, err := funcs.RertiveDetails(Url)
		if err != nil {
			logfuncs.LogError.Errorln("[£] We Skipped ", Url, " Because Error = ", err)
			return
		}
		if strings.Contains(headers.ContentType, "text/html") {

			resp, err := http.Get(Url)
			if err != nil {
				logfuncs.LogError.Errorln("[£] We Skipped ", Url, "Error = ", err)
				return
			}
			urls, err := funcs.ParseHtmlTags(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				logfuncs.LogError.Errorln("[£] We Skipped ", Url, "Error = ", err)
				return
			}
			for _, item := range *urls {
				go workerLoop(item)
				/*go func() {
					wg.Add(1)
					semaphore.Acquire()
					workerLoop(item)
					wg.Done()
					semaphore.Release()
				}()*/
			}
			return
		}
	}

	funcs.PendingItems <- &db.FileDB{
		Url:  Url,
		Type: typeOfFile,
	}

}

func DatabaseHandler() {
	logfuncs.Logger.Infoln("[S] Starting Database Handler")
	for {
		select {
		case item := <-funcs.PendingItems:
			db.AddFileToDb(item)

		case <-time.After(time.Second):
		}
	}
}
