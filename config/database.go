package config

type DBConfiguration struct {
	Dblog        bool
	MaxIdleConns int
	MaxOpenConns int
	Databasehost string
	Schemaname       string
	Username     string
	Password     string
	Drivername   string
}
