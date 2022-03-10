package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func AddUserBack(r *gin.Engine){
	r.POST("/Message/AddUserBack", func(c *gin.Context) {
		userid := c.PostForm("userid")
		merchantid := c.PostForm("merchantid")
		backtime := time.Now().Format("2006-01-02 15:04:05")
		backdata := c.PostForm("backdata")


		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "insert into usersback(Uid,Mid,Bdate,Bdata,status) values(\""+userid+"\",\""+merchantid+"\",\""+backtime+"\",\""+backdata+"\",\""+"0"+"\") "
		insertRowDemo(sqlStr)
		err = db.Close()
		if err != nil {
			return
		}
		c.JSON(200,gin.H{
			"userid": userid,
			"merchantid": merchantid,
			"backtime": backtime,
			"backdata": backdata,
		})


	})
}


func queryUserBack(r *gin.Engine){
	r.POST("/Message/queryUserBack", func(c *gin.Context) {
		userid := c.PostForm("userid")
		merchantid := c.PostForm("merchantid")

		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "select Bdate,Bdata,status from usersback where Uid='"+userid+"'"+" and Mid='"+merchantid+"'"+"order by Bdate"
		stmt, err := db.Prepare(sqlStr)
		if err != nil{
			fmt.Printf("prepare failed, err:%v\n",err)
			return
		}
		defer stmt.Close()
		rows, err := stmt.Query()//如果用db.Query(sql,args) 就不行，必须用db.Query(sql)，否则也会报expected 0 arguments,got 1这个错误。
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return
		}
		// 非常重要：关闭rows释放持有的数据库链接
		defer rows.Close()
		var  usersthing[] UserBackValue
		var g UserBackValue
		for rows.Next() {
			err := rows.Scan(&g.Bdate, &g.Bdata, &g.Status)
			usersthing = append(usersthing, g)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}
		}

		c.JSON(200,gin.H{
			"data": usersthing,
		})


	})
}


func AddMerReply(r *gin.Engine){
	r.POST("/Message/AddMerReply", func(c *gin.Context) {
		userid := c.PostForm("userid")
		merchantid := c.PostForm("merchantid")
		replytime := time.Now().Format("2006-01-02 15:04:05")
		backdata := c.PostForm("backdata")


		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "insert into usersback(Mid,Uid,Bdate,Bdata,status) values(\""+merchantid+"\",\""+userid+"\",\""+replytime+"\",\""+backdata+"\",\""+"1"+"\") "
		insertRowDemo(sqlStr)
		err = db.Close()
		if err != nil {
			return
		}
		c.JSON(200,gin.H{
			"merchantid": merchantid,
			"userid": userid,
			"backtime": replytime,
			"backdata": backdata,
		})


	})
}


func queryMerReply(r *gin.Engine){
	r.POST("/Message/queryMerReply", func(c *gin.Context) {
		merchantid := c.PostForm("merchantid")


		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "select Uid,Rdate,Rdata from ordersdata where Uid='"+merchantid+"'"+"order by Rdate"
		stmt, err := db.Prepare(sqlStr)
		if err != nil{
			fmt.Printf("prepare failed, err:%v\n",err)
			return
		}
		defer stmt.Close()
		rows, err := stmt.Query()//如果用db.Query(sql,args) 就不行，必须用db.Query(sql)，否则也会报expected 0 arguments,got 1这个错误。
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return
		}
		// 非常重要：关闭rows释放持有的数据库链接
		defer rows.Close()
		var  Merthing[] MerrReplyValue
		var g MerrReplyValue
		for rows.Next() {
			err := rows.Scan(&g.Uid, &g.Rdate, &g.Rdata)
			Merthing = append(Merthing, g)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}
		}

		c.JSON(200,gin.H{
			"hello": Merthing,
		})



	})
}