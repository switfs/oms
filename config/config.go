package config

var (
	SR ServerConfig
	CR ClientConfig
)

type ServerConfig struct {
	Mysql *Mysql `json:"mysql"`
	Redis *Redis `json:"redis"`
}

type Server struct {
	Hosts    string `json:"hosts"`
	Ports    string `json:"ports"`
	Ssl      bool   `json:"ssl"`
	CertFile string `json:"certFile"`
	KeyFile  string `json:"keyFile"`
}

type Mysql struct {
	DbType       string `json:"dbtype"`
	Hosts        string `json:"hosts"`
	Ports        string
	Users        string `json:"users"`
	Passwd       string `json:"passwd"`
	DbName       string `json:"dbName"`
	MaxIdleConns int
	MaxOpenConns int
	Engine       string
	Local        string
	GetSqlDebug  int 
}

type Redis struct {
	Hosts  string `json:"hosts"`
	Db     int    `json:"db"`
	Passwd string `json:"passwd"`
}

type ClientConfig struct {
	Client Client `json:"client"`
}

type Client struct {
	Hosts string `json:"hosts"`
	Ports string `json:"ports"`
}
