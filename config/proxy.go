package config

type ProxyConfig struct {
	Client Proxy `json:"proxy"`
}

type Proxy struct {
	Enabled bool   `json:"enabled"`
	Hosts   string `json:"hosts"`
	Ports   string `json:"ports"`
}
