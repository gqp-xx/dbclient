package database

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"reflect"
	"strconv"
	"testing"
)

func TestDbConfig(t *testing.T) {
	tomlStr := `
[db]
[db.master]
addr="127.0.0.1"
port=3306
user_name="test"
password="test"
db_name="testdb"
time_loc="Local"
[db.slave]
addr="127.0.0.1"
port=3306
user_name="slave"
password="slave"
db_name="slavedb"
time_loc="Local"
`
	configs := new(DbConfig)
	_, err := toml.Decode(tomlStr, configs)
	if err != nil {
		fmt.Println("DecodeFile error")
	}
	fmt.Println(configs)

	for _, config := range configs.AllDbConf {
		configType := reflect.TypeOf(config)
		configVal := reflect.ValueOf(config)
		if configType.Kind() != reflect.Ptr {
			panic("InitDBClient error, config type invalied")
		}
		configType = configType.Elem()
		configVal = configVal.Elem()
		for i := 0; i < configType.NumField(); i++ {
			fieldType := configType.Field(i)
			fieldVal := configVal.Field(i)
			defaultTagVal := fieldType.Tag.Get("default")
			if defaultTagVal == "" {
				continue
			}
			switch fieldVal.Kind() {
			case reflect.String:
				if fieldVal.String() == "" {
					fieldVal.SetString(defaultTagVal)
				}
			case reflect.Int:
				if fieldVal.Int() == 0 {
					newVal, _ := strconv.ParseInt(defaultTagVal, 10, 64)
					fieldVal.SetInt(newVal)
				}
			case reflect.Ptr:
				if fieldVal.IsNil() {
					newVal, _ := strconv.ParseBool(defaultTagVal)
					fieldVal.Set(reflect.ValueOf(&newVal))
				}
			default:
				fmt.Errorf("类型不支持")
			}
		}
	}
}

func TestMysqlClient(t *testing.T) {
	tomlStr := `
[db]
[db.master]
addr="192.169.0.1"
port=3306
user_name="test"
password="test"
db_name="testdb"
time_loc="Local"
[db.slave]
addr="192.169.0.2"
port=3306
user_name="slave"
password="slave"
db_name="slavedb"
time_loc="Local"
`
	configs := new(DbConfig)
	_, err := toml.Decode(tomlStr, configs)
	if err != nil {
		fmt.Println("DecodeFile error")
	}
	fmt.Println(configs)
	InitDBClient(configs)
	fmt.Println(GetDbClient("master"))
}
