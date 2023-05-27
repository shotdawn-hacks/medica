package destination

type Config struct {
	ID      string `json:"_id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"hostname" bson:"hostname"`
	Port    string `json:"port" bson:"port"`
}

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
