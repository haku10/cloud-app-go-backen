package service

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

// StorageOutput -> GCSへのファイル出力
func StorageOutput() {

	// クライアントを作成する
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// GCSオブジェクトを書き込むファイルの作成
	file, err := os.Create("testfiles/sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	// オブジェクトのReaderを作成
	bucketName := os.Getenv("GOOGLE_STORAGE_BUCKET")
	folderName := os.Getenv("GOOGLE_STORAGE_FOLDER")
	objectPath := folderName + "/sample.txt"
	writer := client.Bucket(bucketName).Object(objectPath).NewWriter(ctx)
	_, err = io.Copy(writer, file)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	log.Println("done")
}
