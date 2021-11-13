FROM golang:latest

# イメージのアップグレード
RUN apt update -y
RUN apt upgrade -y

# 初期化
COPY . /go/src/
COPY go.mod /go/src/go.mod
COPY go.sum /go/src/go.sum
WORKDIR /go/src/

# Go.modに記載のライブラリをインストール
ENV GO111MODULE=on
RUN go mod download

# sql-migrateが何故かインストールされないので個別でインストール
RUN go get -v github.com/rubenv/sql-migrate/...

CMD ["go", "run", "main.go"]
