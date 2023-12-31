package main

import (
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
	"log"
	"os"
)

type Student struct {
	Name    string  `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Age     int32   `parquet:"name=age, type=INT32, encoding=PLAIN"`
	Id      int64   `parquet:"name=id, type=INT64"`
	Weight  float32 `parquet:"name=weight, type=FLOAT"`
	Sex     bool    `parquet:"name=sex, type=BOOLEAN"`
	Day     int32   `parquet:"name=day, type=INT32, convertedtype=DATE"`
	Ignored int32   //without parquet tag and won't write
}

type TitanicPassengerA struct {
	Id int64 `parquet:"name=passenger_id, type=INT64, convertedtype=INT_64, repetitiontype=OPTIONAL"`
}

func main() {
	var err error
	w, err := os.Create("export/flat.parquet")
	if err != nil {
		log.Println("Can't create local file", err)
		return
	}

	//write
	pw, err := writer.NewParquetWriterFromWriter(w, new(TitanicPassengerA), 4)
	if err != nil {
		log.Println("Can't create parquet writer", err)
		return
	}

	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY
	num := 100
	for i := 0; i < num; i++ {
		//stu := Student{
		//	Name:   "StudentName",
		//	Age:    int32(20 + i%5),
		//	Id:     int64(i),
		//	Weight: float32(50.0 + float32(i)*0.1),
		//	Sex:    bool(i%2 == 0),
		//	Day:    int32(time.Now().Unix() / 3600 / 24),
		//}

		stu := TitanicPassengerA{
			Id: 1,
		}

		if err = pw.Write(stu); err != nil {
			log.Println("Write error", err)
		}
	}
	if err = pw.WriteStop(); err != nil {
		log.Println("WriteStop error", err)
		return
	}
	log.Println("Write Finished")
	w.Close()

}
