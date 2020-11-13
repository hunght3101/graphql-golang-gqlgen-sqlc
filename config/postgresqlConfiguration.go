package config

// PostgresqlConfiguration : Config of postgresql database
type PostgresqlConfiguration struct {
	Host     string
	Port     int32
	User     string
	Password string
	DbName   string
	SSLMode  string
}
