package funcs

import (
	"Go_Web_Scrapper/db"
	"net/url"
	"os"
)

var URL *MyURL
var PendingItems = make(chan *db.FileDB, 200)
var LogFile, _ = os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

type MyURL struct {
	*url.URL
}

/*func (u *MyURL) GetLocalPath() string {
	return path.Join(u.Hostname(), u.Path)
}
*/

type Headers struct {
	ContentType string
	// ContentLength int64
}

// Image file extensions
var imageExtensions = []string{
	".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".svg",
	".webp", ".ico", ".heic", ".heif", ".jfif", ".avif",
}

// Audio file extensions
var audioExtensions = []string{
	".mp3", ".wav", ".ogg", ".flac", ".aac", ".m4a", ".wma",
	".webm", ".amr", ".mid", ".midi", ".opus", ".ac3", ".ec3",
}

// Video file extensions
var videoExtensions = []string{
	".mp4", ".mkv", ".avi", ".mov", ".flv", ".wmv", ".webm",
	".mpeg", ".mpg", ".m4v", ".3gp", ".ogv", ".rm", ".rmvb",
}

const (
	Video   = "video"
	Audio   = "audio"
	Image   = "image"
	Unknown = "unknown"
)
