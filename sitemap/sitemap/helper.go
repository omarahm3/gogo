package sitemap

import (
	"errors"
	"net/url"
	"strings"
)

func ParseLink(link string) *url.URL {
	base, err := url.ParseRequestURI(link)

	if err != nil {
		panic(err)
	}

	base.Path = strings.TrimRight(base.Path, "/")

	return base
}

func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)

	if err != nil {
		return false
	}

	u := ParseLink(toTest)

	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func IsValidHostUrl(baseHostUrl, comparedLink string) bool {
	if baseHostUrl == "" || comparedLink == "" {
		return false
	}

	source := ParseLink(baseHostUrl).Host
	destination := ParseLink(comparedLink).Host

	if source == destination {
		return true
	}

	return false
}

func GenerateValidUrl(baseUrl, path string) (*string, error) {
	if !IsValidUrl(baseUrl) {
		return nil, errors.New("Base URL is not valid URL")
	}

	if IsValidUrl(path) {
		if !IsValidHostUrl(baseUrl, path) {
			return nil, errors.New("Link is from another host")
		}

    path = ParseLink(path).String()
		return &path, nil
	}

  
  link := ParseLink(baseUrl)

	if strings.Contains(path, "mailto:") {
		return nil, errors.New("email links are not supported")
	}

	l := link.JoinPath(strings.Trim(path, "/")).String()
  l = ParseLink(l).String()

	if IsValidUrl(l) {
		return &l, nil
	}

	return nil, errors.New("unable to generate valid URL")
}
