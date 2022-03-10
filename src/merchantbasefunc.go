package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RegisterMerchant(r *gin.Engine){
	r.POST("/User/RegisterMerchant", func(c *gin.Context) {
		merchantUser := c.PostForm("merchantUser")
		pwdmerchant := c.PostForm("pwdmerchant")
		merchantPhone := c.PostForm("phonemerchant")
		merchantEmail := c.PostForm("emailmerchant")
		nameCheck := sqlDuplicateChecking("Mname","merchantsdata",merchantUser)
		emailCheck := sqlDuplicateChecking("Memail","merchantsdata",merchantEmail)
		phoneCheck := sqlDuplicateChecking("Mphone","merchantsdata",merchantPhone)
		if nameCheck == "Yes" {
			if emailCheck == "Yes" {
				if phoneCheck =="Yes"{
					err := initDB()
					if err != nil {
						return
					}
					sqlStr := "select max(Mid) from merchantsdata"
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
						err := rows.Scan(&u.id)
						if err != nil {
							fmt.Printf("scan failed, err:%v\n", err)
							return
						}
					}
					maxid, _ := strconv.Atoi(u.id)
					maxid = maxid+1
					newid := strconv.Itoa(maxid)
					sqlStr = "insert into merchantsdata(Mid,Mname,Mpwd,Mphone,Memail) values(\""+newid+"\",\""+merchantUser+"\",\""+pwdmerchant+"\",\""+merchantPhone+"\",\""+merchantEmail+"\") "
					insertRowDemo(sqlStr)
					err = db.Close()
					if err != nil {
						return
					}

					c.JSON(200,gin.H{
						"merchantid": newid,
						"account": merchantUser,
						"pwd": pwdmerchant,
						"phone": merchantPhone,
						"email": merchantEmail,

					})
				}else {
					c.JSON(200,gin.H{
						"msg": "phone has been used",
					})
				}
			}else {
				c.JSON(200,gin.H{
					"msg": "email has been used",
				})
			}
		}else {
			c.JSON(200,gin.H{
				"msg": "name has been used",
			})
		}



	})
}


func QueryMerchant(r *gin.Engine){
	r.POST("/Merchant/QueryMerchant", func(c *gin.Context) {
		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "select Mid,Mname,Mservicetime,Mphone from merchantsdata"
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
		var m MerBaseValue
		var merchants []MerBaseValue
		for rows.Next() {
			err := rows.Scan(&m.Id, &m.Name, &m.ServiceTime, &m.Phone)
			merchants = append(merchants, m)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}

		}
		err = db.Close()
		if err != nil {
			return
		}

		c.JSON(200,gin.H{
			"data": merchants,
		})

	})
}