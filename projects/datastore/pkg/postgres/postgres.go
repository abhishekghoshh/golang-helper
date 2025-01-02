package postgres

type PostgresDB interface {
	Create(person *Person) (*Person, error)
	Close()
}
type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	Db       string
	SSLMode  string
}
