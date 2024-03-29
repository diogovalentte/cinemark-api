package scraper

// Movie is the struct for a movie
type Movie struct {
	// Title is the movie title
	Title string
	// CoverImgURL is the URL for the movie cover image
	CoverImgURL string
	// URL is the URL for the movie details page
	URL string
	// AgeRating is the movie rating, can be: L, 12, 14, 16, 18
	AgeRating string
	// AgeRatingColor is the color for the movie rating
	AgeRatingColor string
	// New is a boolean that indicates if the movie is new
	New bool
}
