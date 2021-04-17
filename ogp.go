package ogp

import (
	"bytes"
	"strings"

	"github.com/greytabby/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"golang.org/x/xerrors"
)

type OpenGraph struct {
	// Basic metadata
	Title  string
	Type   string
	Images []*Image
	URL    string

	// Optional metadata
	Audios         []*Audio
	Videos         []*Video
	Description    string
	Determiner     string
	Locale         string
	LocaleAltanate []string
	SiteName       string
}

type Image struct {
	URL       string
	SecureURL string
	Type      string
	Width     string
	Height    string
	Alt       string
}

type Video struct {
	URL       string
	SecureURL string
	Type      string
	Width     string
	Height    string
}

type Audio struct {
	URL       string
	SecureURL string
	Type      string
}

func Parse(document []byte) (*OpenGraph, error) {
	return parse(document)
}

func parse(document []byte) (*OpenGraph, error) {
	doc, err := html.Parse(bytes.NewBuffer(document))
	if err != nil {
		return nil, err
	}
	ogpNodes := scrape.FindAll(doc, func(node *html.Node) bool {
		if node.DataAtom != atom.Meta {
			return false
		}
		return strings.HasPrefix(scrape.Attr(node, "property"), "og:")
	})

	if len(ogpNodes) == 0 {
		return nil, xerrors.New("This document is not open graph object")
	}

	og := &OpenGraph{}
	for i := 0; i < len(ogpNodes); i++ {
		node := ogpNodes[i]
		prop := scrape.Attr(node, "property")
		content := scrape.Attr(node, "content")
		switch {
		case strings.EqualFold(prop, "og:title"):
			og.Title = content
		case strings.EqualFold(prop, "og:type"):
			og.Type = content
		case strings.EqualFold(prop, "og:url"):
			og.URL = content
		case strings.EqualFold(prop, "og:description"):
			og.Description = content
		case strings.EqualFold(prop, "og:detarminer"):
			og.Determiner = content
		case strings.EqualFold(prop, "og:locale"):
			og.Locale = content
		case strings.EqualFold(prop, "og:locale:alternate"):
			og.LocaleAltanate = append(og.LocaleAltanate, content)
		case strings.EqualFold(prop, "og:site_name"):
			og.SiteName = content
		case strings.EqualFold(prop, "og:image"):
			img := &Image{}
			img.URL = content
			og.Images = append(og.Images, img)
		case strings.HasPrefix(prop, "og:image:"):
			curImg := og.Images[len(og.Images)-1]
			for strings.HasPrefix(prop, "og:image:") {
				switch {
				case strings.EqualFold(prop, "og:image:secure_url"):
					curImg.SecureURL = content
				case strings.EqualFold(prop, "og:image:type"):
					curImg.Type = content
				case strings.EqualFold(prop, "og:image:width"):
					curImg.Width = content
				case strings.EqualFold(prop, "og:image:height"):
					curImg.Height = content
				case strings.EqualFold(prop, "og:image:alt"):
					curImg.Alt = content
				}
				i++
				if i == len(ogpNodes) {
					break
				}
				node = ogpNodes[i]
				prop = scrape.Attr(node, "property")
				content = scrape.Attr(node, "content")
			}
			i--
		case strings.EqualFold(prop, "og:video"):
			video := &Video{}
			video.URL = content
			og.Videos = append(og.Videos, video)
		case strings.HasPrefix(prop, "og:video:"):
			curAudio := og.Images[len(og.Audios)-1]
			for strings.HasPrefix(prop, "og:video:") {
				switch {
				case strings.EqualFold(prop, "og:video:secure_url"):
					curAudio.SecureURL = content
				case strings.EqualFold(prop, "og:video:type"):
					curAudio.Type = content
				case strings.EqualFold(prop, "og:video:width"):
					curAudio.Width = content
				case strings.EqualFold(prop, "og:video:height"):
					curAudio.Height = content
				}
				i++
				if i == len(ogpNodes) {
					break
				}
				node = ogpNodes[i]
				prop = scrape.Attr(node, "property")
				content = scrape.Attr(node, "content")
			}
			i--
		case strings.EqualFold(prop, "og:audio"):
			audio := &Audio{}
			audio.URL = content
			og.Audios = append(og.Audios, audio)
		case strings.HasPrefix(prop, "og:audio:"):
			curAudio := og.Images[len(og.Audios)-1]
			for strings.HasPrefix(prop, "og:audio:") {
				switch {
				case strings.EqualFold(prop, "og:audio:secure_url"):
					curAudio.SecureURL = content
				case strings.EqualFold(prop, "og:audio:type"):
					curAudio.Type = content
				}
				i++
				if i == len(ogpNodes) {
					break
				}
				node = ogpNodes[i]
				prop = scrape.Attr(node, "property")
				content = scrape.Attr(node, "content")
			}
			i--
		}
	}
	return og, nil
}
