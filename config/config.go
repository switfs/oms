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
	Port     string `json:"port"`
	Ssl      bool   `json:"ssl"`
	CertFile string `json:"certFile"`
	KeyFile  string `json:"keyFile"`
}

type Mysql struct {
	Hosts  string `json:"hosts"`
	Users  string `json:"users"`
	Passwd string `json:"passwd"`
	DbName string `json:"dbName"`
}

type Redis struct {
	Hosts  string `json:"hosts"`
	Db     int    `json:"db"`
	Passwd string `json:"passwd"`
}

type ClientConfig struct {
}
