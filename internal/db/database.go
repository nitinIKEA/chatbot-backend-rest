package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/godror/godror"
	"github.com/nitinIKEA/chatbot-backend-rest/internal/config"
)

type DBConns map[string]*sql.DB

func GetConnections(conf *config.Conf) DBConns {
	conns := DBConns{}
	//Create connection for Dev env
	var connectionParams godror.ConnectionParams
	// Connection for Dev Env
	connectionParams.ConnectString = fmt.Sprintf("%v:%v/%v", conf.DBConfigDev.DBHost, conf.DBConfigDev.DBPort, conf.DBConfigDev.DBServiceName)
	connectionParams.SessionTimeout = time.Duration(60) * time.Second
	connectionParams.Username = conf.DBConfigDev.DBUserName
	connectionParams.Password = godror.NewPassword(conf.DBConfigDev.DBPassword)
	// log.Print("Connect to ", connectionParams.ConnectString)
	conns["Dev"] = sql.OpenDB(godror.NewConnector(connectionParams))
	if err := conns["Dev"].Ping(); err != nil {
		log.Fatalf("error while ping to Dev database: %v", err)
	}

	// Connection for Test Env
	connectionParams.ConnectString = fmt.Sprintf("%v:%v/%v", conf.DBConfigTest.DBHost, conf.DBConfigTest.DBPort, conf.DBConfigTest.DBServiceName)
	connectionParams.SessionTimeout = time.Duration(60) * time.Second
	connectionParams.Username = conf.DBConfigTest.DBUserName
	connectionParams.Password = godror.NewPassword(conf.DBConfigTest.DBPassword)
	// log.Print("Connect to ", connectionParams.ConnectString)
	conns["Test"] = sql.OpenDB(godror.NewConnector(connectionParams))
	if err := conns["Test"].Ping(); err != nil {
		log.Fatalf("error while ping to Test database: %v", err)
	}
	return conns
}
