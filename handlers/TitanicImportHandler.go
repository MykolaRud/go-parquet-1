package handlers

import (
	"fmt"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
	"log"
	"math"
	"parquet_1/models"
	"parquet_1/repositories"
)

func NewTitanicImportHandler(filename string) *TitanicImportHandler {
	return &TitanicImportHandler{filename: filename}
}

type TitanicImportHandler struct {
	filename string
}

func (handler *TitanicImportHandler) Run() {
	// ./parquet_tools -cmd schema -file data/Titanic.parquet -tag

	//clean mysql table
	TitanicRepo := repositories.NewTitanicRepository()
	TitanicRepo.TruncateRecords()

	//fill data
	fr, err := local.NewLocalFileReader(handler.filename)
	if err != nil {
		log.Println("Can't open file ", handler.filename)
		return
	}

	pr, err := reader.NewParquetReader(fr, new(models.TitanicPassenger), 4)
	if err != nil {
		log.Println("Can't create parquet reader", err)
		return
	}
	num := int(pr.GetNumRows())

	fmt.Println("Num rows: ", num)
	dataBatchSize := 10
	maxIteration := int(math.Ceil(float64(num) / float64(dataBatchSize)))
	for i := 0; i < maxIteration; i++ {
		rows := make([]models.TitanicPassenger, dataBatchSize)
		if err = pr.Read(&rows); err != nil {
			log.Println("Read error", err)
		}

		TitanicRepo.AddPassengersBatch(rows)
	}

	log.Println("Import finished")
	pr.ReadStop()
	fr.Close()
}
