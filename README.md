# Cinemark API

A very simple and limited API for the Cinemark site using web scraping.

# How to run:

1. Install the dependencies:

```
go mod download
```

2. Run:

```
go run main.go
```

# Important routes:

- `/v1/health`
- `/v1/movies/in-theaters?city=name` returns the movies in the theaters of a city.
- `/v1/movies/in-theaters-iframe?city=name` returns the movies in the theaters of a city in an HTML document that can be used as an iFrame (designed to be used with [Homarr](https://github.com/ajnart/homarr)).
