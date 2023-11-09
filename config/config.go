package config

var (
	SR ServerConfig
	CR ClientConfig
	PR ProxyConfig
)

type ServerConfig struct {
	Mysql *Mysql `json:"mysql"`
	Redis *Redis `json:"redis"`
}

type Server struct {
	Enabled  bool   `json:"enabled"`
	Hosts    string `json:"hosts"`
	Ports    string `json:"ports"`
	Ssl      bool   `json:"ssl"`
	CertFile string `json:"certFile"`
	KeyFile  string `json:"keyFile"`
}

type Mysql struct {
	Enabled      bool   `json:"enabled"`
	DbType       string `json:"dbtype"`
	Hosts        string `json:"hosts"`
	Ports        string
	Users        string `json:"users"`
	Passwd       string `json:"passwd"`
	DbName       string `json:"dbName"`
	MaxIdleConns int    `json:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns"`
}

type Redis struct {
	Enabled bool   `json:"enabled"`
	Hosts   string `json:"hosts"`
	Db      int    `json:"db"`
	Passwd  string `json:"passwd"`
}
