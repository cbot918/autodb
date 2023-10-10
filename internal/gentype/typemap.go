package gentype

type mapp map[string]string

func NewTypeMap() mapp {

	m := mapp{
		"int":       "int",
		"varchar":   "string",
		"text":      "string",
		"integer":   "int32",
		"bigint":    "int64",
		"datetime":  "time.Time",
		"float":     "float32",
		"tinyint":   "int8",
		"timestamp": "time.Time",

		// "tinyint unsigned":  "uint8",
		// "smallint unsigned": "uint16",
		// "integer unsigned":  "uint32",
		// "bigint unsigned":   "uint64",

		// "double":            "float64",
		// "date":              "time.Time",

	}
	return m
}

/*

	m := mapp{
		"tinyint":           "int8",
		"smallint":          "int16",
		"int":               "int32",
		"integer":           "int32",
		"bigint":            "int64",
		"tinyint unsigned":  "uint8",
		"smallint unsigned": "uint16",
		"integer unsigned":  "uint32",
		"bigint unsigned":   "uint64",
		"float":             "float32",
		"double":            "float64",
		"varchar":           "string",
		"text":              "string",
		"date":              "time.Time",
		"datetime":          "time.Time",
		"timestamp":         "time.Time",
	}

*/
/*
Go	SQL(MySQL)
int8	tinyint
int16	smallint
int32	integer
int64	bigint
uint8	tinyint unsigned
uint16	smallint unsigned
uint32	integer unsigned
uint64	bigint unsigned
float32	float
float64	double
string	varchar
string	text
time.Time	date
time.Time	datetime
time.Time	timestamp
*/
