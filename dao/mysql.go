package dao

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"storageSystem/pbfiles"
)

var db *sqlx.DB

func Init()(err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/storage?charset=utf8mb4&parseTime=True"

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return nil
}

func Query(key string)(value string,err error) {
	sqlStr := "select value from `system` where `ikey`=?"
	err = db.Get(&value ,sqlStr, key)
	if err != nil {
		return "",err
	}
	return value,nil
}

// 插入数据
func Insert(u pbfiles.ProdRegister) error{
	res, err :=Query(u.Key)
	if res != "" {
		Update(u)
		return nil
	} else{
		sqlStr := "insert into `system`(ikey, value) values (?,?)"
		_, err = db.Exec(sqlStr, u.Key, u.Value)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
	return nil
}

// 更新数据
func Update(u pbfiles.ProdRegister)error {
	sqlStr := "update `system` set value=? where `ikey` = ?"
	ret, err := db.Exec(sqlStr, u.Value, u.Key)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil{
		return err
	}
	if n!= 1 {
		return errors.New("更细失败，没有匹配的key")
	}
	return nil
}

// 删除数据
func Delete(key string) error{
	sqlStr := "delete from `system` where `ikey` = ?"
	ret, err := db.Exec(sqlStr, key)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		return err
	}
	if n== 1 {
		return nil
	} else{
		return errors.New("删除失败，没有匹配的key")
	}
}