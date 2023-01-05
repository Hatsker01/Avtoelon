package config

import "os"
import 	"github.com/spf13/cast"


type Config struct{
	Environment string
	PostgresHost string
	PostgresPort int
	PostgresDatabase string
	PostgresUser string
	PostgresPassword string
	LogLevel string
	HTTPPort string
	SignKey string
}

func Load() Config{
	c:=Config{}

	c.Environment= cast.ToString(getOrReturnDefault("ENVIRONMENT","release"))

	c.HTTPPort=cast.ToString(getOrReturnDefault("HTTP_PORT",":8080"))

	c.PostgresHost=cast.ToString(getOrReturnDefault("POSTGRES_HOST","localhost"))
	c.PostgresPort=cast.ToInt(getOrReturnDefault("POSTGRES_PORT",5432))
	c.PostgresDatabase=cast.ToString(getOrReturnDefault("POSTGRES_DATABASE","Avtoelon"))
	c.PostgresUser=cast.ToString(getOrReturnDefault("POSTGRES_USER","postgres"))
	c.PostgresPassword=cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD","1"))

	c.LogLevel=cast.ToString(getOrReturnDefault("LOG_LEVEL","debug"))
	c.SignKey=cast.ToString(getOrReturnDefault("SIGN_KEY","sijxoxyffnfxemfhuoehmniihgs"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{})interface{}{
	_,exists:=os.LookupEnv(key)
	if exists{
		return os.Getenv(key)
	}
	return defaultValue
}