package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func imgServer(r *gin.Engine){
	r.StaticFS("/photo",http.Dir("./res/img"))
}


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

//数据库查重
func sqlDuplicateChecking(column string,table string, columnar string) string {
	err := initDB()
	if err != nil {
		return "sql err"
	}
	sqlStr := "select * from "+table+" where "+ column +"='"+ columnar +"'"
	fmt.Printf("%s\n",sqlStr)
	stmt, err := db.Prepare(sqlStr)
	if err != nil{
		fmt.Printf("prepare failed, err:%v\n",err)
		return "sql err"
	}
	defer stmt.Close()
	rows, err := stmt.Query()//如果用db.Query(sql,args) 就不行，必须用db.Query(sql)，否则也会报expected 0 arguments,got 1这个错误。
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return "sql err"
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	length := 0
	for rows.Next() {
		length++
		fmt.Printf("%d\n",length)
	}
	err = db.Close()
	if err != nil {
		return "sql err"
	}
	if length == 0{
		return "Yes"
	}else {
		return "No"
	}
}
