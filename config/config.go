package config

import (
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func LoadConfig(path string) Iconfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load env failed: %v", err)
	}

	return &config{
		app: &app{
			host: envMap["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["APP_PORT"])
				if err != nil {
					log.Fatalf("Load App Port Failed: %v", err)
				}
				return p
			}(),
			name:    envMap["APP_NAME"],
			version: envMap["APP_VERSION"],
			readTimeOut: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_READ_TIMEOUT"])
				if err != nil {
					log.Fatalf("Load Read Time Out Failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			writeTimeOut: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_WRTIE_TIMEOUT"])
				if err != nil {
					log.Fatalf("Load Write Time Out Failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			bodyLimit: func() int {
				limit, err := strconv.Atoi(envMap["APP_BODY_LIMIT"])
				if err != nil {
					log.Fatalf("Load Body Limit Failed: %v", err)
				}

				return limit
			}(),
			fileLimit: func() int {
				limit, err := strconv.Atoi(envMap["APP_FILE_LIMIT"])
				if err != nil {
					log.Fatalf("Load File Limit Failed: %v", err)
				}

				return limit
			}(),
			gcpBucket: envMap["APP_GCP_BUCKET"],
		},
		db:  &db{
			host: envMap["DB_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["DB_PORT"])
				if err != nil {
					log.Fatalf("Load Port DB Failed: %v", err)
				}
				return p
			}(),
			protocol: envMap["DB_PROTOCOL"],
			username: envMap["DB_USERNAME"],
			password: envMap["DB_PASSWORD"],
			database: envMap["DB_DATABASE"],
			sslMode: envMap["DB_SSL_MODE"],
			maxConnection: func() int {
				con, err := strconv.Atoi(envMap["DB_MAX_CONNECTIONS"])
				if err != nil {
					log.Fatalf("Load Max Connection Failed: %v", err)
				}
				return con
			}(),
		},
		jwt: &jwt{
			adminKey: envMap["APP_ADMIN_KEY"],
			secretKey: envMap["JWT_SECRET_KEY"],
			apiKey: envMap["APP_API_KEY"],
			accessExpiresAt: func() int {
				ex, err := strconv.Atoi(envMap["JWT_ACCESS_EXPIRES"])
				if err != nil {
					log.Fatalf("Load Access Expires Failed: %v", err)
				}
				return ex
			}(),
			refreshExpiresAt: func() int {
				ref, err := strconv.Atoi(envMap["JWT_REFRESH_EXPIRES"])
				if err != nil {
					log.Fatalf("Load Refresh Expires Failed: %v", err)
				}
				return ref
			}(),
		},
	}
}
// Struct
type config struct {
	app *app
	db  *db
	jwt *jwt
}
// Interface
type Iconfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
}
type IAppConfig interface {
}
type IDbConfig interface{
}
type IJwtConfig interface{
}


func (a *config) App() IAppConfig {
	return &app{}
}

type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeOut  time.Duration
	writeTimeOut time.Duration
	bodyLimit    int //bytes
	fileLimit    int //bytes
	gcpBucket    string
}



func (a *config) Db() IDbConfig {
	return &db{}
}

type db struct {
	host          string
	port          int
	protocol      string
	username      string
	password      string
	database      string
	sslMode       string
	maxConnection int
}



func (a *config) Jwt() IJwtConfig {
	return &jwt{}
}

type jwt struct {
	adminKey         string
	secretKey        string
	apiKey           string
	accessExpiresAt  int //seconds
	refreshExpiresAt int //seconds
}
