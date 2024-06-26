// Package scraper implements the scraper for the cinemark site
package scraper

import (
	"github.com/gocolly/colly/v2"
)

var siteURL = "https://www.cinemark.com.br"

// Scraper is the struct for the cinemark site scraper
type Scraper struct {
	c *colly.Collector
}

var userAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:30.0) Gecko/20100101 Firefox/30.0"

func newCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(userAgent),
	)

	return c
}

func (s *Scraper) resetCollector() {
	if s.c != nil {
		s.c.Wait()
	}

	s.c = newCollector()
}

func getMovieAgeRatingColor(rating string) string {
	switch rating {
	case "L":
		return "#00bb22"
	case "12":
		return "#edcb0c"
	case "14":
		return "#f6962d"
	case "16":
		return "#dd021c"
	case "18":
		return "#000"
	default:
		return "gray"
	}
}
