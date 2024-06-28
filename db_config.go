package database

type DbConfig struct {
	AllDbConf map[string]*SingleDb `toml:"db"`
}

type SingleDb struct {
	Addr         string `toml:"addr"`
	Port         int    `toml:"port"`
	UserName     string `toml:"user_name"`
	Pwd          string `toml:"password"`
	DbName       string `toml:"db_name"`
	SoTimeOut    int    `toml:"so_timeout" default:"50"`
	ReadTimeOut  int    `toml:"read_timeout"   default:"300"`
	ParseTime    *bool  `toml:"parse_time" default:"true"`
	TimeLoc      string `toml:"time_loc" default:"local"`
	Charset      string `toml:"cahrset" default:"utf8"`
	MaxOpenConns int    `toml:"max_conn" default:"5"`
	MaxIdleConns int    `toml:"max_idle" default:"2"`
}
