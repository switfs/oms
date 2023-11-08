package config

type ProxyConfig struct {
	Client Proxy `json:"proxy"`
}

type Proxy struct {
	Hosts string `json:"hosts"`
	Ports string `json:"ports"`
}
