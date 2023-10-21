package configs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type App struct {
	Server struct {
		Host string
		Port string
	}

	Mysql struct {
		Host string
		Port string
		Name string
		User string
		Pass string
	}

	Redis struct {
		Host string
		Port string
	}

	JWTSecret string
	Midtrans  struct {
		Sandbox    string
		Production string
	}

	Jaeger struct {
		Host string
		Port string
	}
}

var app *App

var Core coreapi.Client

func SetupMidtransSandbox() {
	conf := GetConfig()
	midtrans.ServerKey = conf.Midtrans.Sandbox
	midtrans.Environment = midtrans.Sandbox
}

func GetMysqlConn() *gorm.DB {
	loggger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             5 * time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)

	conf := GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&charset=utf8mb4&loc=Local", conf.Mysql.Name, conf.Mysql.Pass, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggger,
	})
	if err != nil {
		log.Printf("Cant connect to database because %v", err)
	}

	sqldb, errs := db.DB()
	if errs != nil {
		log.Printf("Something error : %v", errs)
	}

	sqldb.SetConnMaxIdleTime(10)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxLifetime(time.Hour)

	return db
}

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}

	return app
}

func initConfig() *App {
	conf := App{}
	if err := godotenv.Load(); err != nil {
		// config mysql env
		conf.Mysql.Host = "localhost"
		conf.Mysql.Port = "3306"
		conf.Mysql.Name = ""
		conf.Mysql.User = "root"
		conf.Mysql.Pass = ""

		// config redis env
		conf.Redis.Host = "localhost"
		conf.Redis.Port = ""

		// config server env
		conf.Server.Host = "localhost"
		conf.Server.Port = "8080"

		// config jwtsecret env
		conf.JWTSecret = ""

		// config secret midtrans env
		conf.Midtrans.Sandbox = ""
		conf.Midtrans.Production = ""

		conf.Jaeger.Host = ""
		conf.Jaeger.Port = ""

		return &conf
	}

	conf.Mysql.Host = os.Getenv("MYSQL_HOST")
	conf.Mysql.Pass = os.Getenv("MYSQL_PASS")
	conf.Mysql.User = os.Getenv("MYSQL_USER")
	conf.Mysql.Port = os.Getenv("MYSQL_PORT")
	conf.Mysql.Name = os.Getenv("MYSQL_DATABASE")

	conf.Redis.Host = os.Getenv("REDIS_HOST")
	conf.Redis.Port = os.Getenv("REDIS_PORT")

	conf.Server.Host = os.Getenv("APP_HOST")
	conf.Server.Port = os.Getenv("APP_PORT")

	conf.JWTSecret = os.Getenv("JWT_SECRET")
	conf.Midtrans.Sandbox = os.Getenv("MIDTRANS_SANDBOX")
	conf.Midtrans.Production = os.Getenv("MIDTRANS_PRODUCTION")

	conf.Jaeger.Host = os.Getenv("JAEGER_HOST")
	conf.Jaeger.Port = os.Getenv("JAEGER_PORT")
	return &conf
}
