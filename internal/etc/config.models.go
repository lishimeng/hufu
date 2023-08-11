package etc

type Configuration struct {
	Db      db
	Web     web
	Redis   redis
	Session session
}

type web struct {
	Listen string
}

type redis struct {
	Enable   bool
	Addr     string
	Password string
	Db       int
}

type db struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Ssl      string
}

type session struct {
	Key    string
	Alg    string
	Issuer string
}
