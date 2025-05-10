# Pala

A performant and reliable election management system.

## Frontend
To start a local dev server,
You need `npm` or `bun` installed
```sh
bun run dev
```
## Backend
To compile the backend, you need `golang 1.23.0` installed

```sh
GIN_MODE=release CGO_ENABLED=0 go build -o server -ldflags='-s -w' cmd/server/main.go
./server
```

Or you can run the docker image directly,

```sh
cd backend
docker build -t pala:0.0.3 .
sudo docker compose up
```

Alternatively, if `make` is installed,
```sh
make frontend # for the frontend
make backend # for the backend
```

# Report
To compile the report, you need [Typst](https://typst.app/) installed

``` sh
typst compile submissions/report.typ
```

`
