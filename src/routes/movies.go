package routes

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/diogovalentte/cinemark-api/src/scraper"
)

// MoviesRoutes registers the movies routes
func MoviesRoutes(group *gin.RouterGroup) {
	moviesGroup := group.Group("/movies")
	moviesGroup.GET("/in-theaters", GetInTheatersMovies)
	moviesGroup.GET("/in-theaters-iframe", GetMoviesiFrame)
}

// GetInTheatersMovies returns in theater movies for a specific city
func GetInTheatersMovies(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
		return
	}

	queryLimit := c.Query("limit")
	var limit int
	var err error
	if queryLimit == "" {
		limit = -1
	} else {
		limit, err = strconv.Atoi(queryLimit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be a number"})
		}
	}

	scraper := scraper.Scraper{}
	movies, err := scraper.GetInTheatersMovies(city, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// GetMoviesiFrame returns an iframe with the in theater movies for a specific city
func GetMoviesiFrame(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
		return
	}

	queryLimit := c.Query("limit")
	var limit int
	var err error
	if queryLimit == "" {
		limit = -1
	} else {
		limit, err = strconv.Atoi(queryLimit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be a number"})
		}
	}

	scraper := scraper.Scraper{}
	movies, err := scraper.GetInTheatersMovies(city, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	iframe, err := getMoviesiFrame(movies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "text/html", []byte(iframe))
}

func getMoviesiFrame(movies []scraper.Movie) ([]byte, error) {
	containerWidth := "1.6"
	if len(movies) > 3 {
		containerWidth = "18"
	}

	tmpl := template.Must(template.New("movies").Parse(strings.Replace(`
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Movie Display Template</title>
    <style>
      body {
        background-color: #f0f0f0;
        margin: 0;
        padding: 0;
      }

      .movie-container {
        width: calc(100% - MOVIE-CONTAINER-WIDTHpx);
        height: 84px;

        position: relative;
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin-bottom: 14px;

        border-radius: 10px;
        border: 1px solid rgba(56, 58, 64, 1);
      }

      .background-image {
        background-image: url('https://static.vecteezy.com/system/resources/previews/025/470/292/large_2x/background-image-date-at-the-cinema-popcorn-ai-generated-photo.jpeg');
        background-position: center;
        background-size: cover;
        position: absolute;
        filter: brightness(0.3);
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        z-index: -1;
        border-radius: 10px;
      }

      .movie-cover {
        border-radius: 2px;
        margin-left: 20px;
        object-fit: cover;
        width: 30px;
        height: 50px;
      }

      .movie-details {
        flex: 1;
        padding: 0 20px;
      }

      .movie-name {
        font-size: 15px;
        font-weight: bold;
        color: white;

        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji";

        text-decoration: none;
        overflow: hidden;
        text-overflow: ellipsis;
      }

      .new-movie-tag {
        display: inline-block;
        padding: 5px 10px;
        margin: 20px;
        background-color: rgb(150, 109, 109, 0.5);
        color: rgb(230, 101, 101);

        text-decoration: none; /* Remove underline */
        border-radius: 5px;
        font-size: 20px;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
          Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji";
      }

      .movie-age-rating {
        display: inline-block;
        padding: 7px 10px;
        margin-right: 20px;
        width: 42.08px;
        height: 42.08px;
        border-radius: 5px;
        box-sizing: border-box;

        font-size: 20px;
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
          Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji";
        color: white;
        font-weight: 800;
        text-align: center;
      }
    </style>
  </head>
  <body style="background-color: #25262b">
    {{range .}}
        <div class="movie-container">
          <div class="background-image"></div>
          <img
            class="movie-cover"
            src="{{.CoverImgURL}}"
            alt="Movie Cover"
          />
          <div class="movie-details">
            <a href="{{.URL}}" target="_blank" class="movie-name">{{.Title}}</a>
          </div>
          {{ if .New }}
            <div class="new-movie-tag">New</div>
          {{end}}

        <div style="background-color: {{.AgeRatingColor}}" class="movie-age-rating">{{.AgeRating}}</div>
        </div>
    {{end}}
  </body>
</html>
	`, "MOVIE-CONTAINER-WIDTH", containerWidth, -1)))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, movies)
	if err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}
