package ddl

// DDL is responsible for schema change.
type DDL interface {
	CreateDatabase()
	CreateTable()
	DropTable()
	DropDatabase()
}
