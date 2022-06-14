package pre

// prefabricate fields
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

func ChangeCrontab(v string) {
	Crontab = v
}

func ChangeLogrus(v string) {
	Logrus = v
}

func ChangeMysql(v string) {
	Mysql = v
}

func ChangeOracle(v string) {
	Oracle = v
}

func ChangeRedis(v string) {
	Redis = v
}

func ChangePostgres(v string) {
	Postgres = v
}

func ChangeRocketMq(v string) {
	RocketMq = v
}

func ChangeTableStore(v string) {
	TableStore = v
}
