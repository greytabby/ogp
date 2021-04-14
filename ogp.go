package ogp

type OpenGraph struct {
	// Basic metadata
	Title  string
	Type   string
	Images []Image
	URL    string

	// Optional metadata
	Audios      []Audio
	Videos      []Video
	Description string
	Determiner  string
	Locale      Locale
	SiteName    string
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

type Locale struct {
	Alternate []string
}
