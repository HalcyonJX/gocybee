package main

import (
	"gocybee/work7/api"
	"gocybee/work7/dao"
)

func main() {
	dao.InitRedis()
	dao.InitDB()
	api.InitRouter()
}
