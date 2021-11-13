フォーマットツールは`gofmt`を使用する

# 前準備
https://console.cloud.google.com/
に自身のアカウントでログインし、Text To Speechの認証を設定しておく。

## GCPのクレデンシャルファイル名をenvに設定する
ホスト側のマシンの`$HOME/.gcp`にクレデンシャルファイルを配置する

.envファイルにGOOGLE_APPLICATION_CREDENTIALSを記載する


## ライブラリ追加
`go get [ライブラリ名]`
go.modに追加される


## migration
* into container
`docker exec -it コンテナID bash`
`cd /go`

* sql-migrate
`sql-migrate status`
* add migration file
`sql-migrate new [ファイル名]`
* create table
`sql-migrate up`
* migrate down
`sql-migrate down`

