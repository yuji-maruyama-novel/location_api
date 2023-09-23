# goバージョン
FROM golang:1.21.0-alpine as builder
# アップデートとgitのインストール
RUN apk update && apk add git
RUN apk add bash

# zonefileの追加
RUN apk add tzdata

# go関連
ENV TZ=Asia/Tokyo

# ローカル環境
FROM builder as local

RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/cosmtrek/air@latest
# golang-migrateをインストール
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . /go/src

# ワーキングディレクトリの設定
# WORKDIR /go/src/gqlgen
WORKDIR /go/src

# airの起動
CMD ["air", "-c", ".air.toml"]

EXPOSE 3004