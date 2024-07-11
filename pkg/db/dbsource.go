package db

type DB struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Name        string `yaml:"name"`
	IP2Location string `yaml:"ip2l"`
	Maxmind     string `yaml:"mm"`
}

type Dbsource struct {
	Database DB `yaml:"db"`
}
