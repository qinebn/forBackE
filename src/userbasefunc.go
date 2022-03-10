package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)


func RegisterUser(r *gin.Engine){
	r.POST("/User/RegisterUser", func(c *gin.Context) {
		accountUser := c.PostForm("accountUser")
		pwdUser := c.PostForm("pwdUser")
		accountEmail := c.PostForm("emailUser")
		nameCheck :=sqlDuplicateChecking("Uname","usersdata",accountUser)
		emailCheck := sqlDuplicateChecking("Uemail","usersdata",accountEmail)

		if nameCheck == "Yes"{
			if emailCheck == "Yes"{
				err := initDB()
				if err != nil {
					return
				}
				sqlStr := "select Uid, Uname, Upwd, Uemail from usersdata"
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
				mostid := 0
				for rows.Next() {
					var u usersdata
					err := rows.Scan(&u.id, &u.name, &u.pwd, &u.email)
					if err != nil {
						fmt.Printf("scan failed, err:%v\n", err)
						return
					}
					tempid, err := strconv.Atoi(u.id)
					if tempid > mostid {
						mostid = tempid
					}
				}
				mostid += 1
				newid := strconv.Itoa(mostid)


				sqlStr = "insert into usersdata(Uid,Uname,Upwd,Uemail) values(\""+newid+"\",\""+accountUser+"\",\""+pwdUser+"\",\""+accountEmail+"\") "
				insertRowDemo(sqlStr)
				err = db.Close()
				if err != nil {
					return
				}
				c.JSON(200,gin.H{
					"account": accountUser,
					"pwd": pwdUser,
					"email": accountEmail,
				})
			}else {
				c.JSON(200,gin.H{
					"code":1,
					"msg": "email has been used",
				})
			}
		}else {
			c.JSON(200,gin.H{
				"code":1,
				"msg": "name has been used",
			})
		}

	})
}


func Login(r *gin.Engine){
	r.POST("/User/Login", func(c *gin.Context) {
		accountID := c.PostForm("idUser")
		accountName := c.PostForm("nameUser")
		accountEmail := c.PostForm("emailUser")
		pwdUser := c.PostForm("pwdUser")
		typeUser := c.PostForm("typeUser")
		err := initDB()
		if err != nil {
			return
		}
		sqlStr := ""
		if typeUser == "cust"{
			if accountName==""{
				if accountID == ""{
					sqlStr = "select * from usersdata where Uemail="+"'"+accountEmail+"'"
				}else if accountEmail =="" {
					sqlStr = "select * from usersdata where Uid=" + "'" + accountID + "'"
				}
			}else {
				sqlStr = "select * from usersdata where Uname="+"'"+accountName+"'"
			}


		} else if typeUser == "busi" {
			if accountName ==""{
				if accountID == ""{
					sqlStr = "select Mid,Mname,Mpwd,Memail from merchantsdata where Memail="+"'"+accountEmail+"'"
				}else if accountEmail ==""{
					sqlStr = "select Mid,Mname,Mpwd,Memail from merchantsdata where Mid="+"'"+accountID+"'"
				}
			}else {
				sqlStr = "select Mid,Mname,Mpwd,Memail from merchantsdata where Mname="+"'"+accountName+"'"
			}

		}else{
			return
		}
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
		var u usersdata
		for rows.Next() {
			err := rows.Scan(&u.id, &u.name, &u.pwd, &u.email)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}
		}
		err = db.Close()
		if err != nil {
			return
		}

		if u.pwd == pwdUser{
			c.JSON(200, gin.H{
				"code": "0",
				"id":u.id,
				"name":u.name,
				"email":u.email,
			})
		} else {
			c.JSON(200,gin.H{
				"code": "1",
			})
		}
	})
}




