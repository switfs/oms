package config

type ProxyConfig struct {
	Proxy Proxy `json:"proxy"`
	Mysql Mysql `json:"mysql"`
}

type Proxy struct {
	Enabled bool   `json:"enabled"`
	Hosts   string `json:"hosts"`
	Ports   string `json:"ports"`
}
