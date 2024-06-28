package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strconv"
	"sync"
)

var (
	dbs  = make(map[string]*sql.DB)
	ones sync.Once
)

/**
 * db clients 初始化，需要在调用前初始化，只需初始化一次
 */
func InitDBClient(configs *DbConfig) {
	ones.Do(func() {
		if configs == nil || len(configs.AllDbConf) == 0 {
			panic("InitDBClient error, config is nil")
		}

		for dbname, config := range configs.AllDbConf {
			setDefaultVal(config)

			dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%dms&readTimeout=%dms&charset=%s&parseTime=%v&loc=%s",
				config.UserName, config.Pwd, config.Addr, config.Port, config.DbName, config.SoTimeOut, config.ReadTimeOut,
				config.Charset, *config.ParseTime, config.TimeLoc)
			dbClient, err := sql.Open("mysql", dbConfig)
			if err != nil {
				fmt.Errorf("open mysql error, err is:%+v", err)
				panic("open mysql error")
			}
			dbClient.SetMaxOpenConns(config.MaxOpenConns)
			dbClient.SetMaxIdleConns(config.MaxIdleConns)
			dbClient.Ping()
			dbs[dbname] = dbClient
		}
	})
}

func setDefaultVal(config *SingleDb) {
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
		defaultval := fieldType.Tag.Get("default")
		if defaultval == "" {
			continue
		}
		switch fieldVal.Kind() {
		case reflect.String:
			if fieldVal.String() == "" {
				fieldVal.SetString(defaultval)
			}
		case reflect.Int:
			if fieldVal.Int() == 0 {
				newVal, _ := strconv.ParseInt(defaultval, 10, 64)
				fieldVal.SetInt(newVal)
			}
		case reflect.Ptr:
			if fieldVal.IsNil() {
				newVal, _ := strconv.ParseBool(defaultval)
				fieldVal.Set(reflect.ValueOf(&newVal))
			}
		default:
			fmt.Errorf("类型不支持")
		}
	}
}

/**
 * 通过名称获取dbclient,需要先调用初始化
 */
func GetDbClient(name string) (*sql.DB, error) {
	client, ok := dbs[name]
	if !ok {
		fmt.Printf("dbclient %s not exists", name)
		return nil, errors.New("client not exists")
	}
	return client, nil
}
