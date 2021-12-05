package vars

// constants
const (
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
)

// variables
var (
	Crontab = "Crontab"

	Logrus     = "Logrus"
	Format     = "Format"
	TTYDiscard = "TTYDiscard"
	GrayLog    = "GrayLog"
	Enable     = "Enable"
	Extra      = "Extra"

	Host         = "Host"
	Addr         = "Addr"
	EndPoint     = "EndPoint"
	InstanceName = "InstanceName"
	Port         = "Port"

	Mysql      = "Mysql"
	Oracle     = "Oracle"
	Redis      = "Redis"
	Postgres   = "Postgres"
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
	SSLMode         = "SSLMode"

	Secure = "Secure"
	Key    = "Key"
)

func SetCrontab(v string) {
	Crontab = v
}

func SetLogrus(v string) {
	Logrus = v
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
