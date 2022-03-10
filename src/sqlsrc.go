package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)


// 定义一个初始化数据库的函数
func initDB() (err error) {
	dsn := "root:101579@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 查询多条数据示例
//func queryMultiRowDemo(sqlStr string, rowsdata *sql.Rows){
//	stmt, err := db.Prepare(sqlStr)
//	if err != nil{
//		fmt.Printf("prepare failed, err:%v\n",err)
//		return
//	}
//	defer func(stmt *sql.Stmt) {
//		err := stmt.Close()
//		if err != nil {
//		}
//	}(stmt)
//	rows, err := stmt.Query()//如果用db.Query(sql,args) 就不行，必须用db.Query(sql)，否则也会报expected 0 arguments,got 1这个错误。
//	if err != nil {
//		fmt.Printf("query failed, err:%v\n", err)
//		return
//	}
//	// 非常重要：关闭rows释放持有的数据库链接
//	defer func(rows *sql.Rows) {
//		err := rows.Close()
//		if err != nil {
//		}
//	}(rows)
//	rowsdata = rows
//	// 循环读取结果集中的数据
//	//for rows.Next() {
//	//	var u usersdata
//	//	err := rows.Scan(&u.Uid, &u.Uname, &u.Upwd, &u.Uemail)
//	//	if err != nil {
//	//		fmt.Printf("scan failed, err:%v\n", err)
//	//		return
//	//	}
//	//	fmt.Printf("%s %s %s %s \n", u.Uid, u.Uname, u.Upwd, u.Uemail)
//	//}
//}

func insertRowDemo(sqlStr string) {
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	ret, err := stmt.Exec()
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	var theID int64
	theID, err = ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}//插入

func updateRowDemo(sqlStr string) {
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	ret, err := stmt.Exec()
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	var n int64
	n, err = ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}//更新

func deleteRowDemo() {
	sqlStr := "delete from commodity where comid = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	ret, err := stmt.Exec(6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	var n int64
	n, err = ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}// 删除数据
