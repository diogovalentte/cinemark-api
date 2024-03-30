# Cinemark API

A very simple and limited API for the Cinemark site using web scraping.

# How to run:

## Run using Docker:

1. Build the Docker image:

```sh
docker build -t cinemark-api .
```

2. Run:

```sh
docker run --name cinemark-api -p 8080:8080 cinemark-api
```

## Run manually:

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
  - `city` (not optional): city to get the movies in all of the city's theaters. First, check if the Cinemark site has a page for this city, if it doesn't, it'll return the page of São Paulo by default. Go to https://cinemark.com.br/rio-de-janeiro/filmes/em-cartaz and select your city.
  - `limit` (optional): limit the number of movies returned.
- `/v1/movies/in-theaters-iframe` returns the movies in the theaters of a city in an HTML document that can be used as an iFrame (designed to be used with [Homarr](https://github.com/ajnart/homarr)). Allow the following query arguments:
  - `city` (not optional): city to get the movies in all of the city's theaters. First, check if the Cinemark site has a page for this city, if it doesn't, it'll return the page of São Paulo by default. Go to https://cinemark.com.br/rio-de-janeiro/filmes/em-cartaz and select your city.
  - `limit` (optional): limit the number of movies returned.
  - `theme` (optional): "light" or "dark". It's used to match the HTML returned with the Homarr theme. Defaults to "light".
