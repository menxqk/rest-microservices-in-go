package users_db

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/menxqk/rest-microservices-in-go/common/logger"
)

const (
	MYSQL_USERS_USERNAME = "MYSQL_USERS_USERNAME"
	MYSQL_USERS_PASSWORD = "MYSQL_USERS_PASSWORD"
	MYSQL_USERS_HOST     = "MYSQL_USERS_HOST"
	MYSQL_USERS_SCHEMA   = "MYSQL_USERS_SCHEMA"
)

var (
	Client *sql.DB
)

func init() {
	loadEnv()

	username := os.Getenv(MYSQL_USERS_USERNAME)
	password := os.Getenv(MYSQL_USERS_PASSWORD)
	host := os.Getenv(MYSQL_USERS_HOST)
	schema := os.Getenv(MYSQL_USERS_SCHEMA)

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		fmt.Println("could not ping database:", err)
	}

	mysql.SetLogger(logger.GetLogger())

	fmt.Println("database successfully configured")
}

func loadEnv() {
	f, err := os.Open(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		elems := strings.Split(line, "=")
		if len(elems) == 2 {
			err := os.Setenv(elems[0], elems[1])
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
