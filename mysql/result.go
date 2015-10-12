package mysql

type Result struct {
	Status uint16

	InsertId     uint64
	AffectedRows uint64

	*Resultset
}

type RowData []byte

type Resultset struct {
	Fields     []*Field
	FieldNames map[string]int
	Values     [][]interface{}

	RowDatas []RowData
}
