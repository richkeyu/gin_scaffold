package db

import (
	"context"
	"fmt"
	"gateway/pkg/constant"
	"testing"

	"github.com/richkeyu/gocommons/config"
)

func TestMysqlConn(t *testing.T) {

	var err error
	if _, err = config.NewConfig("../../config/app_dev.yaml"); err != nil {
		panic(fmt.Sprintf("init config failed, err: %v", err))
	}

	InitMysql()

	type aaa struct {
		MerchantId int64  `json:"merchant_id"`
		ApiKey     string `json:"api_key"`
	}

	var list []aaa

	if err = PaySvrMysql.WithContext(nil).Table("api_key").Where("merchant_id = ?", 2).Find(&list).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Println(list)

}

func TestRedisConn(t *testing.T) {
	var err error
	if _, err = config.NewConfig("../../config/app_dev.yaml"); err != nil {
		panic(fmt.Sprintf("init config failed, err: %v", err))
	}

	InitRedis()

	c := GetKvConn(constant.RedisPaySvr)

	fmt.Println(c.Get(context.Background(), "yf_test").Val())
	//fmt.Println(c.Set(context.Background(), "yf_test", 1, 0).Val())
}
