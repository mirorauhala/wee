package main

import "time"

type ShortenedUrl struct {
	ID        string
	URL       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Shortener struct{}

func (s *Shortener) Shorten(url string) (ShortenedUrl, error) {
	id, err := sid.Generate()

	if err != nil {
		return ShortenedUrl{}, err
	}

	return ShortenedUrl{
		ID:        id,
		URL:       url,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
