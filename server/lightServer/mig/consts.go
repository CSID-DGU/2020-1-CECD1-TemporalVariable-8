package mig

type (
	DBDriver string
)

const (
	SQLITE3 DBDriver = "sqlite3"
	//MySQL DBDriver= "mysql"
	//MariaDB DBDriver= "mysql"
	//Postgre DBDriver= "postgre"
)

const(
	CacheNSFile = "CNS_FILE:"
)