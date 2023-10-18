package config

var Server ServerConfig

type ServerConfig struct {
	Mysql *Mysql `json:"mysql"`
	Redis *Redis `json:"redis"`
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
