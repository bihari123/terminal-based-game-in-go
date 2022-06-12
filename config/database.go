package config

type DBConfiguration struct {
	Dblog        bool
	MaxIdleConns int
	MaxOpenConns int
	Databasehost string
	DbName       string
	Username     string
	Password     string
	Drivername   string
}
