package vars

// variables
var (
	Logrus     = "Logrus"
	Crontab    = "Crontab"
	Mysql      = "Mysql"
	Oracle     = "Oracle"
	Redis      = "Redis"
	Postgres   = "Postgres"
	RocketMq   = "RocketMq"
	TableStore = "TableStore"
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

func SetPostgres(v string) {
	Postgres = v
}

func SetRocketMq(v string) {
	RocketMq = v
}

func SetTableStore(v string) {
	TableStore = v
}
