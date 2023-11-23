package handlers

import (
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
	"log"
	"math"
	"parquet_1/models"
	"parquet_1/repositories"
)

func NewTitanicExportHandler(filename string) *TitanicExportHandler {
	return &TitanicExportHandler{filename: filename}
}

type TitanicExportHandler struct {
	filename string
}

func (handler *TitanicExportHandler) Run() {
	TitanicRepo := repositories.NewTitanicRepository()
	dbRowsCount := TitanicRepo.GetRowsCount()
	log.Println("Rows in DB: ", dbRowsCount)

	fw, err := local.NewLocalFileWriter(handler.filename)
	if err != nil {
		log.Println("Can't open file ", handler.filename)
		return
	}

	pw, err := writer.NewParquetWriter(fw, new(models.TitanicPassenger), 4)
	if err != nil {
		log.Println("Can't create parquet writer", err)
		return
	}

	//pw.RowGroupSize = 128 * 1024 * 1024
	pw.CompressionType = parquet.CompressionCodec_SNAPPY
	batchSize := 30
	maxIteration := int(math.Ceil(float64(dbRowsCount) / float64(batchSize)))
	for i := 0; i < maxIteration; i++ {

		passengers := TitanicRepo.GetPassengersBatch(i*batchSize, batchSize)

		for _, passenger := range passengers {

			log.Println("Writing ", passenger)
			if err = pw.Write(passenger); err != nil {
				log.Println("Write error", err)
			}
		}
	}
	if err = pw.WriteStop(); err != nil {
		log.Println("WriteStop error", err)
		return
	}
	log.Println("Write Finished")
	if err = fw.Close(); err != nil {
		log.Println("Close error", err)
	}
}
