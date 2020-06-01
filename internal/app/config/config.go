package config

// Config is the application configuration
// parse using github.com/vrischmann/envconfig
type Config struct {
	Port           string `envconfig:"WALLET_APP_PORT"`
	MysqlConnR     string `envconfig:"WALLET_MYSQL_CONN_R"`
	MysqlConnW     string `envconfig:"WALLET_MYSQL_CONN_W"`
	RedisAddr      string `envconfig:"WALLET_REDIS_ADDRESS"`
	RedisKey       string `envconfig:"WALLET_REDIS_KEY"`
	UserCtxKey     string `envconfig:"WALLET_USER_CONTEXT_KEY"`
	OAuthEndpoint  string `envconfig:"WALLET_OAUTH2_ENDPOINT"`
	OAuthTokeninfo string `envconfig:"WALLET_OAUTH2_TOKENINFO"`
	LogLevel       uint   `envconfig:"WALLET_LOG_LEVEL"`
}
