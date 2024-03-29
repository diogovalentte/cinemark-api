package scraper

import "testing"

func TestGetInTheatersMovies(t *testing.T) {
	t.Run("should return a list of movies in theaters for a specific city", func(t *testing.T) {
		scraper := Scraper{}
		movies, err := scraper.GetInTheatersMovies("sao-paulo")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if len(movies) == 0 {
			t.Errorf("expected movies, got none")
		}

		for _, movie := range movies {
			if movie.Title == "" {
				t.Errorf("expected movie title, got none")
			}

			if movie.CoverImgURL == "" {
				t.Errorf("expected movie cover, got none")
			}

		}
	})
}
