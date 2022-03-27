package common

import (
	"Algo/cache"
	"Algo/global"
	"Algo/model"
	"fmt"
	"strconv"
)

// 在程序开始时迁移数据
func GetUserInfoToCache(){
	var userState []model.User
	err := global.DBEngine.Find(&userState).Error
	if err !=nil{
		fmt.Printf("%v",err)
		return
	}
	rConn := cache.RedisPool().Get()//连接池中获取
	defer rConn.Close()
	for _,v := range userState{
		rConn.Do("SET",v.Address,strconv.Itoa(v.Status))
	}
}
