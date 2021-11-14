フォーマットツールは`gofmt`を使用する

# 前準備
https://console.cloud.google.com/
に自身のアカウントでログインし、Text To Speechの認証を設定しておく。

## GCPのクレデンシャルファイル名をenvに設定する
ホスト側のマシンの`$HOME/.gcp`にクレデンシャルファイルを配置する

.envファイルにGOOGLE_APPLICATION_CREDENTIALSを記載する

例)
`GOOGLE_APPLICATION_CREDENTIALS=/usr/gcp/XXXX.json`

※ ホスト側の`$HOME/.gcp`がコンテナの`/usr/gcp`にマウントされているため、
上記の設定でコンテナ側の環境変数として設定される

## アップロード処理を行うときはバケット名は各自で作成したものを使用する
GOOGLE_STORAGE_BUCKET=バケット名
GOOGLE_STORAGE_FOLDER=フォルダー名
を.envファイルに設定する
例)
GOOGLE_STORAGE_BUCKET=test_file_output
GOOGLE_STORAGE_FOLDER=test-folder


##　バックエンドの立ち上げ
### ビルド
`docker-compose build`
### 実行
`docker-compose up`


## ライブラリ追加
`go get [ライブラリ名]`
go.modに追加される

docker-composeコマンドで実行する場合は
`docker-compose run cloud-go-google-api-app go get [ライブラリ]`で実行する


## migration
* into container
`docker exec -it コンテナID bash`
`cd /go`

### sql-migrateによるマイグレーション
#### 参照リポジトリ
`https://github.com/rubenv/sql-migrate`

#### マイグレーションのコマンド
* sql-migrate
`sql-migrate status`
* add migration file
`sql-migrate new [ファイル名]`
* create table
`sql-migrate up`
* migrate down
`sql-migrate down`

``



# DynamoDBの初期構築(ローカル環境)
`endpoint_url="http://localhost:9603"`

# テーブル作成
```
aws dynamodb create-table \
  --cli-input-json file://db/dynamoDB/test_table.json \
  --billing-mode PAY_PER_REQUEST \
  --endpoint-url ${endpoint_url} >> result.txt
```
# テストデータ投入
```
aws dynamodb put-item --table-name test_table \
  --item '{"id": {"S": "1"}}' \
  --endpoint-url ${endpoint_url}
```

# ツールインストール (DB Browser for SQLite)
`brew install --cask db-browser-for-sqlite`
