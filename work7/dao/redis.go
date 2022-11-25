package dao

import (
	"github.com/garyburd/redigo/redis"
)

// 定义一个全局的 pool
var Pool *redis.Pool

// 当启动程序时，就初始化连接池
func InitRedis() {
	Pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个ip 的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func IsLike(username string) bool { //当前用户是否点赞
	conn := Pool.Get()
	defer conn.Close()
	_, err := conn.Do("GET", username)
	if err != nil {
		return false
	}
	return true
}

// ---点赞&&取消点赞
func LikeSet(username string) bool {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", username, "like")
	if err != nil {
		return false
	}
	return true
}
func LikeDel(username string) bool {
	conn := Pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", username)
	if err != nil {
		return false
	}
	return true
}
