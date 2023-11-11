package config

var (
	SR ServerConfig
	CR ClientConfig
	PR ProxyConfig
)

type ServerConfig struct {
	Server *Server `json:"service"`
	Mysql  *Mysql  `json:"mysql"`
	Redis  *Redis  `json:"redis"`
	Email  *Email  `json:"email"`
}

type Server struct {
	Hosts      string `json:"hosts"`
	Ports      string `json:"ports"`
	EnabledSSL bool   `json:"enabled_ssl"`
	CertFile   string `json:"certFile"`
	KeyFile    string `json:"keyFile"`
}

type Mysql struct {
	DbType       string `json:"dbtype"`
	Hosts        string `json:"hosts"`
	Ports        string `json:"ports"`
	Users        string `json:"users"`
	Passwd       string `json:"passwd"`
	DbName       string `json:"dbName"`
	MaxIdleConns int    `json:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns"`
}

type Redis struct {
	Enabled bool   `json:"enabled"`
	Hosts   string `json:"hosts"`
	Ports   string `json:"ports"`
	Db      int    `json:"db"`
	Passwd  string `json:"passwd"`
}

type Email struct {
	Hosts      string `json:"hosts"`
	Ports      string `json:"ports"`
	EnabledSSL bool   `json:"enabled_ssl"`
	From       string `json:"from"`
	Secret     string `json:"secret"`
}
