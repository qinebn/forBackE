package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func AddOrder(r *gin.Engine){
	r.POST("/Order/AddOrder", func(c *gin.Context) {
		userid := c.PostForm("userid")
		orderstatus := "0"
		orderaddress := c.PostForm("orderaddress")
		orderscore := "10.0"
		merchantid := c.PostForm("merchantid")
		goodsid := c.PostFormMap("goodsid")
		goodsnum := c.PostFormMap("goodsnum")

		timenowUnix := time.Now().Unix()
		Oid := strconv.Itoa(int(timenowUnix))
		timenow := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf(timenow)
		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "insert into ordersdata(Oid,Ostatus,ObelongtoUid,Ordertime,Oaddress,Oscore) values(\""+Oid+"\",\""+orderstatus+"\",\""+userid+"\",\""+timenow+"\",\""+orderaddress+"\",\""+orderscore+"\") "
		insertRowDemo(sqlStr)
		for iterOrder := range goodsid{
			fmt.Printf(Oid,merchantid,goodsid[iterOrder],goodsnum[iterOrder])
			sqlStr = "insert into orderinformation(Oid,Mid,Gid,Gnum) values(\""+Oid+"\",\""+merchantid+"\",\""+goodsid[iterOrder]+"\",\""+goodsnum[iterOrder]+"\") "
			insertRowDemo(sqlStr)
		}
		err = db.Close()
		if err != nil {
			return
		}
		c.JSON(200,gin.H{
			"code":0,
			"msg": "success",
		})


	})
}


func QueryOrderUser(r *gin.Engine){
	r.POST("/Order/QueryOrderUser", func(c *gin.Context) {
		userid := c.PostForm("userid")


		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "select * from ordersdata where ObelongtoUid='"+userid+"'"+"order by Ordertime"
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
		var  ordersid[] ordersvalue
		var g ordersvalue
		for rows.Next() {
			err := rows.Scan(&g.Oid, &g.Ostatus, &g.ObelongtoUid, &g.Ordertime, &g.Oaddress)
			ordersid = append(ordersid, g)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}
		}

		c.JSON(200,gin.H{
			"data": ordersid,
		})


	})
}


func QueryOrderUserMore(r *gin.Engine){
	r.POST("/Order/QueryOrderUserMore", func(c *gin.Context) {
		orderid := c.PostForm("orderid")

		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "select Mid,Gid,Gnum,Gscore from orderinformation where Oid='"+orderid+"'"
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
		var  ordersid[] ordersmorevalue
		var g ordersmorevalue
		for rows.Next() {
			err := rows.Scan(&g.Mid, &g.Gid, &g.Gnum, &g.Gscore)
			ordersid = append(ordersid, g)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}
		}

		c.JSON(200,gin.H{
			"data": ordersid,
		})
	})
}


func changeOrderStatus(r *gin.Engine){
	r.POST("/Order/changeOrderStatus/evaluate", func(c *gin.Context){
		orderid := c.PostForm("orderid")
		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "update ordersdata set Ostatus='1' where Oid = '"+orderid+"'"
		updateRowDemo(sqlStr)
		c.JSON(200,gin.H{
			"code": 0,
		})
	})
	r.POST("/Order/changeOrderStatus/afterSale", func(c *gin.Context){
		orderid := c.PostForm("orderid")
		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "update ordersdata set Ostatus='2' where Oid = '"+orderid+"'"
		updateRowDemo(sqlStr)
		c.JSON(200,gin.H{
			"code": 0,
		})
	})
	r.POST("/Order/changeOrderStatus/finished", func(c *gin.Context){
		orderid := c.PostForm("orderid")
		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "update ordersdata set Ostatus='3' where Oid = '"+orderid+"'"
		updateRowDemo(sqlStr)
		c.JSON(200,gin.H{
			"code": 0,
		})
	})
}


func QueryOrderMer(r *gin.Engine){
	r.POST("/Order/QueryOrderMer", func(c *gin.Context) {
		merchantid := c.PostForm("merchantid")


		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "select * from ordersdata where ObelongtoUid='"+merchantid+"'"
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
		var  ordersid[] ordersvalue
		var g ordersvalue
		for rows.Next() {
			err := rows.Scan(&g.Oid, &g.Ostatus, &g.ObelongtoUid, &g.Ordertime, &g.Oaddress)
			ordersid = append(ordersid, g)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}
		}

		c.JSON(200,gin.H{
			"hello": ordersid,
		})


	})
}















