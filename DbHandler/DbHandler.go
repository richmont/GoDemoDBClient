package dbhandler

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbHandler interface {
	ShowTables() []string
	checkConnection() bool
	connect() *gorm.DB
	buildDsn() string
}
type MariadbHandler struct {
	dsn        string
	user       string
	password   string
	hostname   string
	port       int64
	dbname     string
	connection *gorm.DB
}

func NewMariadbHandler(user string, pass string, hostname string, port int64, dbname string) MariadbHandler {
	log.Println("Starting NewMariadbHandler")
	return MariadbHandler{
		user:     user,
		password: pass,
		hostname: hostname,
		port:     port,
		dbname:   dbname,
	}
}
func (m *MariadbHandler) Connect() error {
	log.Println("Connect started")
	dsn, errDSN := m.BuildDsn()
	if errDSN != nil {
		log.Println("Error in BuildDsn")
		return errDSN
	} else {
		log.Println("No error in BuildDsn")
		m.dsn = dsn
		db, errOpen := gorm.Open(mysql.New(mysql.Config{
			DSN: dsn,
		}), &gorm.Config{})
		if errOpen != nil {
			log.Println("Error in gorm.Open")
			return errOpen
		} else {
			log.Println("No error in gorm.Open")
			m.connection = db
			return nil

		}

	}

}
func (m *MariadbHandler) BuildDsn() (string, error) {
	log.Println("BuildDsn started")
	if m.user == "" {
		log.Println("BuildDsn Invalid user")
		return "", errors.New("invalid user")
	} else if m.password == "" {
		log.Println("BuildDsn Invalid password")
		return "", errors.New("invalid password")
	} else if m.hostname == "" {
		log.Println("BuildDsn Invalid hostname")
		return "", errors.New("invalid hostname")
	} else if m.port == 0 {
		log.Println("BuildDsn Invalid port")
		return "", errors.New("invalid port")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		m.user, m.password, m.hostname, m.port, m.dbname)
	log.Println("BuildDsn DSN generated: " + dsn)
	return dsn, nil

}
func (m *MariadbHandler) ShowTables() []string {
	var results []string
	m.connection.Raw("show tables").Scan(&results)
	for _, row := range results {
		fmt.Println(row)
	}
	return results
}
func (m *MariadbHandler) getConnection() *gorm.DB {
	return m.connection
}
