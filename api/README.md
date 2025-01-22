#  Indexer API

Go API built using [Chi](https://go-chi.io/#/) to retrieve indexed emails in ZincSearch.

## Getting Started

You can use the provided Dockerfile to run the API in a docker container

```bash
docker build -t indexer-api .
```

Then, run the container (on port 8080 by default)

```bash
docker run -p 8080:8080 --name indexer-api indexer-api
```

