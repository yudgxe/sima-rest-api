package database

var (
	ModeDisable string = "disable"
)

type PostgresConnInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}


