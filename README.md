# My first Go project

Litterally, this is my first project use [Go](https://go.dev) mainly.

-   [🇯🇵 日本語](README_jp.md)
-   [🇬🇧 English](README.md)

## Feature

This is a backend app for write diary and save it to MongoDB.

-   Today's feature
    -   🔒 Simple JWT auth
    -   📖 Write or read your diary
-   Comming soon
    -   🗑️ Rewrite or delete your diary
    -   📒 Basic template
    -   🖼️ Frontend

## Set up

Before the setting up, here are tools that you'll nead at least.

-   Git
-   Go
-   Docker
-   Openssl

1. Clone this repo.

```bash
git clone https://github.com/LinoRino/first_go.git
```

2. Start Docker Container using Docker compose

```bash
docker-compose up -d
```

3. Create **"certificate"** folder & generate ssl in that.

```bash
mkdir certificate
openssl genrsa -out certificate/key.pem 2048
openssl req -new -key certificate/key.pem -out certificate/csr.pem
openssl x509 -req -days 365 -in certificate/csr.pem -signkey certificate/key.pem -out certificate/cert.pem
```

4. Build **"server.go"** & Run it to start https server

```bash
go build server.go
./server
```

## packages

-   Go
    -   [Echo](https://github.com/labstack/echo/) : High performance, minimalist Go web framework
