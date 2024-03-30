# Cinemark API

A very simple and limited API for the [Cinemark Brazil site](https://cinemark.com.br) using web scraping.

# How to run:

## Using Docker:

1. Build the Docker image:

```sh
docker build -t cinemark-api .
```

2. Run:

```sh
docker run --name cinemark-api -p 8080:8080 cinemark-api
```

## Using Docker Compose:

1. Create a file named `docker-compose.yml` with the contents of the `docker-compose.yml` file in this repository.
2. Run:

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
- `/v1/movies/in-theaters-iframe` returns the movies in the theaters of a city in an HTML document that can be used as an iFrame (designed to be used with [Homarr](https://github.com/ajnart/homarr)). Allow the following query arguments:
  - `city`: same as above.
  - `limit`: same as above.
  - `theme` (optional): "light" or "dark". It's used to match the HTML returned with the Homarr theme. Defaults to "light".

# Adding to Homarr

1. Click on "Enter edit mode" -> "Add a tile" -> "Widgets" -> "iFrame".
2. Click to edit the iFrame widget.
3. Add the API URL, like `http://192.168.1.15:8080/v1/movies/in-theaters-iframe?city=rio-de-janeiro&limit=3&theme=dark`. Change the query arguments for your needs.

# Obs:
- The Cinemark site stores your city in a geolocation cookie, if you change the city anytime on the site, you change this cookie’s value. The default value of the cookie is the city of São Paulo, it's set the first time you access the site. If you change the city anytime, it’ll change the value of this cookie, then the next time you access a link to the site, it’ll use this cookie’s value and show your city.
