package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// StorageOutput -> GCSへのファイル出力
func StorageOutput() {
	credentialFilePath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	fmt.Print(credentialFilePath)

	// クライアントを作成する
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}

	// GCSオブジェクトを書き込むファイルの作成
	f, err := os.Create("testfiles/sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	// オブジェクトのReaderを作成
	bucketName := os.Getenv("GOOGLE_STORAGE_BUCKET")
	folderName := os.Getenv("GOOGLE_STORAGE_FOLDER")
	objectPath := folderName + "/sample.txt"
	writer := client.Bucket(bucketName).Object(objectPath).NewWriter(ctx)
	_, err = io.Copy(writer, f)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	log.Println("done")
}
