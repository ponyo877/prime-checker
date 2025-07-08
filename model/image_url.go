package model

type ImageURL struct {
	url string
}

func NewImage(url string) *ImageURL {
	return &ImageURL{
		url: url,
	}
}
