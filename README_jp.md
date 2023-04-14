# 最初の Go プロジェクト

このリポジトリは、私の最初の[Go](https://go.dev)になります。

-   [🇯🇵 日本語](README_jp.md)
-   [🇬🇧 English](README.md)

## 概要

このプロジェクトは、Go と MongoDB を利用した、ユーザーとその日記を管理するバックエンドアプリです。

-   特徴
    -   🔒 JWT を利用した認証・認可
    -   📖 日記の読み書き
-   実装したい部分
    -   🗑️ 日記の訂正、削除
    -   📒 日記削除ンプレート
    -   🖼️ フロントエンド

## Set up

セットアップするには、以下のツールが必要になります。

-   Git
-   Go
-   Docker
-   Openssl

1. このリポをクローンする

```bash
git clone https://github.com/LinoRino/first_go.git
```

2. Docker Compose を使って Docker(MongoDB)を立ち上げる

```bash
docker-compose up -d
```

3. **"certificate"**フォルダに ssl 証明書を用意する

```bash
mkdir certificate
openssl genrsa -out certificate/key.pem 2048
openssl req -new -key certificate/key.pem -out certificate/csr.pem
openssl x509 -req -days 365 -in certificate/csr.pem -signkey certificate/key.pem -out certificate/cert.pem
```

4. **"server.go"**をビルドして、出たファイルを実行し、HTTPS サーバーを開始する

```bash
go build server.go
./server
```

### パッケージ

-   Go
    -   [Echo](https://github.com/labstack/echo/) : High performance, minimalist Go web framework
