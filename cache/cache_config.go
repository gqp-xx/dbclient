package cache

type RedisCacheConf struct {
	RedisConf map[string]*RedisConfig `toml:"redis"`
}

// 基于gopkg.in/redis.v5 redis.Options
type RedisConfig struct {
	Addr         string `toml:"addr"` // host:port
	Password     string `toml:"password"`
	DB           int    `toml:"db_index"`
	ReadTimeout  int    `toml:"read_timeout"`  // 读超时，默认3s
	WriteTimeout int    `toml:"write_timeout"` // 写超时，默认3s
	PoolSize     int    `toml:"pool_size"`     // 连接池大小，默认10
	ConnTimeout  int    `toml:"conn_timeout"`  // 建立连接超时，默认为5s
	IdleTimeout  int    `toml:"idle_timeout"`  // 连接空闲时间段后关闭，默认不关闭空闲连接
}
