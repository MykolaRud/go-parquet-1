package models

type TitanicPassengerV struct {
	Id       int64   `parquet:"name=PassengerId, type=INT64, convertedtype=INT_64, repetitiontype=OPTIONAL"`
	Survived int64   `parquet:"name=Survived, type=INT64, convertedtype=INT_64, repetitiontype=OPTIONAL"`
	PClass   int64   `parquet:"name=Pclass, type=INT64, convertedtype=INT_64, repetitiontype=OPTIONAL"`
	Name     string  `parquet:"name=Name, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Sex      string  `parquet:"name=Sex, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Age      float64 `parquet:"name=Age, type=DOUBLE, repetitiontype=OPTIONAL"`
	SibSp    int64   `parquet:"name=SibSp, type=INT64, convertedtype=INT_64, repetitiontype=OPTIONAL"`
	Parch    int64   `parquet:"name=Parch, type=INT64, convertedtype=INT_64, repetitiontype=OPTIONAL"`
	Ticket   string  `parquet:"name=Ticket, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Fare     float64 `parquet:"name=Fare, type=DOUBLE, repetitiontype=OPTIONAL"`
	Cabin    string  `parquet:"name=Cabin, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
	Embarked string  `parquet:"name=Embarked, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"`
}

type TitanicPassenger struct {
	Id       int64   `parquet:"name=PassengerId, type=INT64, convertedtype=INT_64"`
	Survived int64   `parquet:"name=Survived, type=INT64, convertedtype=INT_64"`
	PClass   int64   `parquet:"name=Pclass, type=INT64, convertedtype=INT_64"`
	Name     string  `parquet:"name=Name, type=BYTE_ARRAY, convertedtype=UTF8"`
	Sex      string  `parquet:"name=Sex, type=BYTE_ARRAY, convertedtype=UTF8"`
	Age      float64 `parquet:"name=Age, type=DOUBLE"`
	SibSp    int64   `parquet:"name=SibSp, type=INT64, convertedtype=INT_64"`
	Parch    int64   `parquet:"name=Parch, type=INT64, convertedtype=INT_64"`
	Ticket   string  `parquet:"name=Ticket, type=BYTE_ARRAY, convertedtype=UTF8"`
	Fare     float64 `parquet:"name=Fare, type=DOUBLE"`
	Cabin    string  `parquet:"name=Cabin, type=BYTE_ARRAY, convertedtype=UTF8"`
	Embarked string  `parquet:"name=Embarked, type=BYTE_ARRAY, convertedtype=UTF8"`
}

var TitanicJsonSchema string = `
{
  "Tag": "name=Duckdb_schema, repetitiontype=REQUIRED",
  "Fields": [
    {
      "Tag": "name=PassengerId, type=INT64, convertedtype=INT_64, repetitiontype=OPTIONAL"
    },
    {
      "Tag": "name=Survived, type=INT64, convertedtype=INT_64, repetitiontype=OPTIONAL"
    },
    {
      "Tag": "name=Pclass, type=INT64, convertedtype=INT_64, repetitiontype=OPTIONAL"
    },
    {
      "Tag": "name=Name, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"
    },
    {
      "Tag": "name=Sex, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"
    },
    {
      "Tag": "name=Age, type=DOUBLE, repetitiontype=OPTIONAL"
    },
    {
      "Tag": "name=SibSp, type=INT64, convertedtype=INT_64, repetitiontype=OPTIONAL"
    },
    {
      "Tag": "name=Parch, type=INT64, convertedtype=INT_64, repetitiontype=OPTIONAL"
    },
    {
      "Tag": "name=Ticket, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"
    },
    {
      "Tag": "name=Fare, type=DOUBLE, repetitiontype=OPTIONAL"
    },
    {
      "Tag": "name=Cabin, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"
    },
    {
      "Tag": "name=Embarked, type=BYTE_ARRAY, convertedtype=UTF8, repetitiontype=OPTIONAL"
    }
  ]
}
`
