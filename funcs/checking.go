package funcs

import (
	"path"
	"strings"
)

func ExtOf(ext string) string {
	switch {
	case contains(ext, &imageExtensions):
		return Image
	case contains(ext, &videoExtensions):
		return Audio
	case contains(ext, &audioExtensions):
		return Video
	default:
		return Unknown
	}
}

func contains(target string, array *[]string) bool {
	for _, element := range *array {
		if element == target {
			return true
		}
	}
	return false
}

func reValidateUrl(url string) string {
	switch {
	case strings.HasPrefix(url, "http://"), strings.HasPrefix(url, "https://"):
		return url
	case strings.HasPrefix(url, "/"):
		return URL.Scheme + "://" + URL.Hostname() + url
	default:
		_url := URL
		_url.Path = path.Join(_url.Path, url)
		return _url.String()
	}
}
