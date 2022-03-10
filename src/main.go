package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	mRun()

}
func mRun(){
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.Use(Cors())
	RegisterUser(r)
	Login(r)
	RegisterMerchant(r)
	AddGoods(r)
	QueryGoods(r)
	AddOrder(r)
	QueryOrderUser(r)
	QueryOrderUserMore(r)
	AddUserBack(r)
	queryUserBack(r)
	AddMerReply(r)
	queryMerReply(r)
	QueryMerchant(r)
	changeOrderStatus(r)
	imgServer(r)
	err := r.Run(":6003")
	if err != nil {
		return
	}

}