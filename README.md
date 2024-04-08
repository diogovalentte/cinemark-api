# 2024-04-08: Moved to the project [Homarr iFrames](https://github.com/diogovalentte/homarr-iframes)

# Cinemark API

A very simple and limited API for the [Cinemark Brazil site](https://cinemark.com.br) using web scraping.

This is how the iFrame is shown on the dashboard (on the right in the image below). It is made based on the Homarr app to show movies/show requests on Jellyseer/Overseer (on the left).
![image](https://github.com/diogovalentte/cinemark-api/assets/49578155/fe6fd2bf-f2b1-45b7-8b40-4676819042f1)

# How to run:

## Using Docker:

1. Run the latest version:

```sh
docker run --name cinemark-api -p 8080:8080 ghcr.io/diogovalentte/cinemark-api:latest
```

## Using Docker Compose:

1. There is a `docker-compose.yml` file in this repository. Clone this repository to use this file or create one.
2. Start the container by running:

```sh
docker compose up
```

## Manually:

1. Install the dependencies:

```sh
go mod download
```

2. Run:

```sh
go run main.go
```

# Simple docs:

- `/v1/health`
- `/v1/movies/in-theaters?city=sao-paulo` returns the movies in the theaters of a city. Allow the following query arguments:
  - `city` (not optional): city to get the movies in all of the city's theaters. First, check if the Cinemark site has a page for this city, if it doesn't, it'll return the page of São Paulo by default. Go to https://cinemark.com.br/rio-de-janeiro/filmes/em-cartaz and select your city. Then grab the city name on the URL.
  - `limit` (optional): limit the number of movies returned.
  - `theaters` (optional): filter the theaters to get the movies. You can find the filter keywords by going to your city page, like https://cinemark.com.br/sao-paulo/filmes/em-cartaz, clicking to filter by theater, and then grabbing the filters in the URL. The filter is the theaters' IDs separated by **%2C**. For example, in the URL https://cinemark.com.br/sao-paulo/filmes/em-cartaz?cinema=716%2C690%2C699 we have the IDs 716, 690, and 699. You have to pass the text `716%2C690%2C699` to the API!
- `/v1/movies/in-theaters-iframe` returns the movies in the theaters of a city in an HTML document that can be used as an iFrame (designed to be used with [Homarr](https://github.com/ajnart/homarr)). Allow the following query arguments:
  - `city`: same as above.
  - `limit`: same as above.
  - `theaters`: same as above.
  - `theme` (optional): "light" or "dark". It's used to match the HTML returned with the Homarr theme. Defaults to "light".

# Adding to Homarr

1. Click on "Enter edit mode" -> "Add a tile" -> "Widgets" -> "iFrame".
2. Click to edit the iFrame widget.
3. Add the API URL, like `http://192.168.1.15:8080/v1/movies/in-theaters-iframe?city=rio-de-janeiro&limit=3&theme=dark`. Change the query arguments for your needs.

# Obs:

- The Cinemark site stores your city in a geolocation cookie, if you change the city anytime on the site, you change this cookie’s value. The default value of the cookie is the city of São Paulo, it's set the first time you access the site. If you change the city anytime, it’ll change the value of this cookie, then the next time you access a link to the site, it’ll use this cookie’s value and show your city.
