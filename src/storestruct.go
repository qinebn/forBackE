package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/*
	此处结构题中变量 首字母 全部用大写
	否则 json 发送数组时无法发送
*/

type usersdata struct {
	id string
	name string
	pwd string
	email string
}



type goodsvalue struct {
	Gid string
	Gname string
	Gnum string
	Gprice string
	Gintroduce string
	Gimglink string
}

type ordersvalue struct {
	Oid string
	Ostatus string
	ObelongtoUid string
	Ordertime string
	Oaddress string
}

type ordersmorevalue struct {
	Mid string
	Gid string
	Gnum string
	Gscore string
}

type UserBackValue struct {
	Bdate string
	Bdata string
	Status string
}

type MerrReplyValue struct {
	Uid string
	Rdate string
	Rdata string
}

type MerBaseValue struct {
	Id string
	Name string
	ServiceTime string
	Phone string
}