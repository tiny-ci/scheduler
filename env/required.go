package env

const (
	RedisHostname = "REDIS_SERVER_HOST"
	RedisUsername = "REDIS_SERVER_USER"
	RedisPassword = "REDIS_SERVER_PASS"
)

var RequiredEnv = map[string]Var{
	"RedisHostname": makeEV(RedisHostname, DTString, false),
	"RedisUsername": makeEV(RedisUsername, DTString, false),
	"RedisPassword": makeEV(RedisPassword, DTString, false),
}
