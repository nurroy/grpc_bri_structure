package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"user_mgt/pkg/converter"
	upt "user_mgt/proto/v1/users"
)

const (
	SELECT_USER_DATA1 = `SELECT *
FROM mst_user
LIMIT 1`
	SELECT_USER_DATA_BYID = `SELECT * from mst_user WHERE uid = '%v'`
	CREATE_USER_DATA = `INSERT INTO mst_user (uid, uname, uaddress, age) VALUES ('%v','%v','%v','%v')`
)


type ConfigStore interface {
	GetKey(string) string
}



// Config is the struct to pass the Postgres configuration.
type PostgresConfig struct {
	Host string
	Port string
	User string
	Password string
	Dbname string
}

const POSTGRES string = "postgres"

var pgConn string //global variable postgres connection

type ResInsertInfo struct {
	LastInsertId 	int64
	RowsAffected 	int64
}

type User struct {
	UId string
	UName string
	UAddress string
	Age int
}

// New initializes the postgres writers
func New(config ConfigStore) error {
	var err error

	// set postgres config from env
	pgConfig := PostgresConfig{
		Host:     string(config.GetKey("postgres.host")),
		Port:     string(config.GetKey("postgres.port")),
		User:     string(config.GetKey("postgres.user")),
		Password: string(config.GetKey("postgres.password")),
		Dbname:   string(config.GetKey("postgres.dbname")),
	}

	// Set postgres config
	pgConn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.User,
		pgConfig.Password,
		pgConfig.Dbname)

	db, err := sql.Open(POSTGRES, pgConn)
	if err != nil {
		return err
	}

	sqlQuery := fmt.Sprintf(SELECT_USER_DATA1)
	rows, err := db.Query(sqlQuery)
	defer db.Close()
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

// func to select data from postgres
func SelectUsersData(idata upt.RetrieveRequest) ([]upt.Users, error) {
	db, err := sql.Open(POSTGRES, pgConn)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(fmt.Sprintf(SELECT_USER_DATA_BYID, converter.StringToFloat64(idata.GetUId())))
	defer db.Close()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []upt.Users

	//Fetch data to struct
	for rows.Next() {
		var rowsScan = upt.Users{}
		err := rows.Scan(&rowsScan.UId, &rowsScan.UName, &rowsScan.UAddress, &rowsScan.Age)

		if err != nil {
			return nil, err
		}

		// Append for every next row
		rowsScanArr = append(rowsScanArr, rowsScan)
	}

	return rowsScanArr, nil
}

// func to select data from postgres
func SelectUserData(idata upt.RetrieveRequest) (*User, error) {
	db, err := sql.Open(POSTGRES, pgConn)
	if err != nil {
		return nil, err
	}

	fmt.Println("ID", idata.UId)

	rows := db.QueryRow(fmt.Sprintf(SELECT_USER_DATA_BYID, converter.StringToFloat64(idata.GetUId())))
	defer db.Close()

	var rowsScan = User{}
	err = rows.Scan(&rowsScan.UId, &rowsScan.UName, &rowsScan.UAddress, &rowsScan.Age)

	if err != nil {
			return nil, err
		}

	return &rowsScan, nil
}

// func to post data from postgres
func InsertUserData(idata upt.Users) (error, error) {
	db, err := sql.Open(POSTGRES, pgConn)
	if err != nil {
		return err, nil
	}

	fmt.Println("data", idata)


	qry := "INSERT INTO mst_user (uid, uname, uaddress, age) VALUES ('%v','%v','%v','%v')"
	_, err = db.Exec(fmt.Sprintf(qry, idata.UId, idata.UName, idata.UAddress, idata.Age))
	defer db.Close()
	if err != nil {
		return err, nil
	}

	return nil, nil
}