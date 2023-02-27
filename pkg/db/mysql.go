package db

import (
	"fmt"
	"gateway/pkg/constant"

	"github.com/richkeyu/gocommons/config"
	mysql "github.com/richkeyu/gocommons/db"

	"gorm.io/gorm"
)

var PaySvrMysql *gorm.DB

func InitMysql() {
	var err error
	var mysqlConf map[string]mysql.MysqlConf

	if err = config.Load("mysql", &mysqlConf); err != nil {
		panic(fmt.Sprintf("load mysql config failed, err: %v", err))
	}

	for dbName, dbConf := range mysqlConf {
		switch dbName {
		case constant.MysqlPaySvr:
			PaySvrMysql, err = mysql.InitMysqlClient(dbConf)
		}

		if err != nil {
			panic("mysql connect error: %v" + err.Error())
		}
	}
}
