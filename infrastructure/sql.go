package infrastructure

import (
	"os"

	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"

	// pq lib
	_ "github.com/lib/pq"
)

const (
	//DBMS Database management system
	DBMS = "postgres"
)

// SQL struct.
type SQL struct {
	// Master connections master database.
	Master *gorm.DB
}

type dbInfo struct {
	host    string
	port    string
	user    string
	pass    string
	name    string
	logmode bool
}

// NewSQL returns new SQL.
func NewSQL() (*SQL, error) {
	dbInfo := dbInfo{
		host:    os.Getenv("DB_READ_HOST"),
		user:    os.Getenv("DB_READ_USER"),
		pass:    os.Getenv("DB_READ_PASSWORD"),
		name:    os.Getenv("DB_NAME"),
		port:    os.Getenv("DB_PORT"),
		logmode: GetenvBool("DB_READ_LOG_MODE"),
	}
	var db *gorm.DB

	connect := "host=" + dbInfo.host + " port=" + dbInfo.port + " user=" + dbInfo.user + " dbname=" + dbInfo.name + " sslmode=disable password=" + dbInfo.pass
	db, err := gorm.Open(DBMS, connect)
	if err != nil {
		return nil, errors.Wrap(err, "can't open database")
	}
	db.LogMode(dbInfo.logmode)
	// Disable table name's pluralization globally
	// if set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected
	db.SingularTable(true)
	return &SQL{db}, nil
}
