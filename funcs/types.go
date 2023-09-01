package funcs

import (
	"net/url"
)

var URL *MyURL

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
