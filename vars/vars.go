package vars

// default configuration names
var (
	AppYml  = "app.yml"
	AppYaml = "app.yaml"

	ApplicationYml  = "application.yml"
	ApplicationYaml = "application.yaml"

	ConfAppYml  = "conf/app.yml"
	ConfAppYaml = "conf/app.yaml"

	ConfApplicationYml  = "conf/application.yml"
	ConfApplicationYaml = "conf/application.yaml"

	DefaultTimeFormat = "2006-01-02 15:04:05"

	JSONContentType = "Content-Type:application/json"

	Crontab = "Crontab"
	Logrus  = "Logrus"

	Format  = "Format"
	TTY     = "TTY"
	GrayLog = "GrayLog"
	Enable  = "Enable"
	Extra   = "Extra"

	Host         = "Host"
	Addr         = "Addr"
	EndPoint     = "EndPoint"
	InstanceName = "InstanceName"
	Port         = "Port"

	Mysql      = "Mysql"
	Oracle     = "Oracle"
	Redis      = "Redis"
	RocketMq   = "RocketMq"
	TableStore = "TableStore"

	User       = "User"
	Datasource = "Datasource"
	Password   = "Password"
	AccessKey  = "AccessKey"
	SecretKey  = "SecretKey"

	AccessKeyId     = "AccessKeyId"
	AccessKeySecret = "AccessKeySecret"
	Retry           = "Retry"
	DB              = "DB"

	Secure = "Secure"
	Key    = "Key"
)

func SetCrontab(v string) {
	Crontab = v
}

func SetLogrus(v string) {
	Logrus = v
}

func SetFormat(v string) {
	Format = v
}

func SetTTY(v string) {
	TTY = v
}

func SetGrayLog(v string) {
	GrayLog = v
}

func SetEnable(v string) {
	Enable = v
}

func SetExtra(v string) {
	Extra = v
}

func SetHost(v string) {
	Host = v
}

func SetAddr(v string) {
	Addr = v
}

func SetEndPoint(v string) {
	EndPoint = v
}

func SetInstanceName(v string) {
	InstanceName = v
}

func SetPort(v string) {
	Port = v
}

func SetMysql(v string) {
	Mysql = v
}

func SetOracle(v string) {
	Oracle = v
}

func SetRedis(v string) {
	Redis = v
}

func SetRocketMq(v string) {
	RocketMq = v
}

func SetTableStore(v string) {
	TableStore = v
}

func SetUser(v string) {
	User = v
}

func SetDatasource(v string) {
	Datasource = v
}

func SetPassword(v string) {
	Password = v
}

func SetAccessKey(v string) {
	AccessKey = v
}

func SetSecretKey(val string) {
	SecretKey = val
}

func SetAccessKeyId(val string) {
	AccessKeyId = val
}

func SetAccessKeySecret(val string) {
	AccessKeySecret = val
}

func SetRetry(val string) {
	Retry = val
}

func SetDB(val string) {
	DB = val
}

func SetSecure(val string) {
	Secure = val
}

func SetKey(val string) {
	Key = val
}
