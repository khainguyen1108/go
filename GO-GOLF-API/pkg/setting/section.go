package setting

type Config struct {
	MySql  MySQLSetting  `mapstructure:"mysql"`
	Server ServerSetting `mapstructure:"server"`
	Log    LogSetting    `mapstructure:"log"`
	Redis  RedisSetting  `mapstructure:"redis"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	UserName        string `mapstructure:"username"`
	PassWord        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type ServerSetting struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type LogSetting struct {
	LogLevel    string `mapstructure:"logLevel"`
	FileLogName string `mapstructure:"fileLogName"`
	MaxSize     int    `mapstructure:"maxSize"`
	MaxBackups  int    `mapstructure:"maxBackUps"`
	MaxAge      int    `mapstructure:"maxAge"`
	Compress    bool   `mapstructure:"compress"`
}
