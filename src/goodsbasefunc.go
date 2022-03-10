package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"time"
)

func AddGoods(r *gin.Engine){
	r.POST("/Merchant/AddGoods", func(c *gin.Context) {
		merchantid := c.PostForm("merchantid")
		goodsname := c.PostForm("goodsname")
		goodsnum := c.PostForm("goodsnum")
		goodsprice := c.PostForm("goodsprice")
		goodsintroduce := c.PostForm("goodsintroduce")
		goodsimg, err := c.FormFile("goodsimg")
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		suffix := path.Ext(goodsimg.Filename)
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt,10)
		fileName := fileNameStr + suffix
		filePath := "D:/GoLand/goCode/try01/res/img/goods/"
		filePath = filePath + fileName
		err = c.SaveUploadedFile(goodsimg, filePath)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		goodslink := "/photo/goods/" + fileName

		err = initDB()
		if err != nil {
			return
		}
		sqlStr := "select max(Gid) from goods where GbelongtoMid='"+ merchantid+"'"
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
		qStatus := 0
		firstgoodsid := merchantid+"000"
		for rows.Next() {
			err := rows.Scan(&u.id)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				qStatus = 1
			}
		}
		newid := ""
		if qStatus == 1{
			maxid, _ := strconv.Atoi(firstgoodsid)
			maxid = maxid+1
			newid = strconv.Itoa(maxid)
		}else{
			maxid, _ := strconv.Atoi(u.id)
			maxid = maxid+1
			newid = strconv.Itoa(maxid)
		}

		sqlStr = "insert into goods(Gid,GbelongtoMid,Gname,Gnum,Gprice,Gintroduce,Gimglink) values(\""+newid+"\",\""+merchantid+"\",\""+goodsname+"\",\""+goodsnum+"\",\""+goodsprice+"\",\""+goodsintroduce+"\",\""+goodslink+"\") "
		insertRowDemo(sqlStr)
		err = db.Close()
		if err != nil {
			return
		}

		c.JSON(200,gin.H{
			"merchantid": merchantid,
			"goodsname": goodsname,
			"goodsnum": goodsnum,
			"goodsprice": goodsprice,
			"goodsintroduce": goodsintroduce,
			//"goodsimg": goodsimg,

		})


	})
}


func QueryGoods(r *gin.Engine){
	r.POST("/Merchant/QueryGoods", func(c *gin.Context) {
		merchantid := c.PostForm("merchantid")
		fmt.Printf(merchantid)
		err := initDB()
		if err != nil {
			return
		}
		sqlStr := "select Gid, Gname, Gnum, Gprice, Gintroduce, Gimglink from goods where GbelongtoMid="+"'"+merchantid+"'"
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
		var  goodsid []goodsvalue
		var g goodsvalue
		for rows.Next() {
			err := rows.Scan(&g.Gid, &g.Gname, &g.Gnum, &g.Gprice, &g.Gintroduce, &g.Gimglink)
			goodsid = append(goodsid, g)
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
			"data": goodsid,
		})
	})
}
