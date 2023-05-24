package destination

func NewConfig(t, address, port string) *Config {
	switch t {
	case TypeCore:
		return newCoreConfig(address, port)
	default:
		return &Config{}
	}
}

func newCoreConfig(address, port string) *Config {
	return &Config{Address: address, Port: port}
}
