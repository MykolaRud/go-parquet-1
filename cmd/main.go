package main

import (
	"context"
	"parquet_1/handlers"
)

func main() {
	exportFileName := "export/TitanicExprot.parquet"
	//titanicImportHandler := handlers.NewTitanicImportHandler("data/Titanic.parquet")
	//titanicImportHandler.Run()

	titanicExportHandler := handlers.NewTitanicExportHandler(exportFileName)
	titanicExportHandler.Run()

	s3Uploader := handlers.NewS3UploadHandler()
	s3Uploader.Save(context.TODO(), exportFileName, "test")
}
