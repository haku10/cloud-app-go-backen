FROM golang:latest

# イメージのアップグレード
RUN apt update -y && apt upgrade -y

# ツールのインストール
RUN apt install -y less vim curl unzip sudo
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
RUN unzip awscliv2.zip
RUN sudo ./aws/install

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
