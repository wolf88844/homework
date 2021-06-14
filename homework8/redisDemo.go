package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func RedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.DialURL("redis://192.168.124.17:6379")
			if err != nil {
				return nil, err
			}
			return conn, err
		},
	}
}

func main() {
	redisPool := RedisPool()
	defer redisPool.Close()

	c1 := redisPool.Get()
	//10byte
	//var byteTen [10]byte
	//for i := 0; i < 10; i++ {
	//	v:=i
	//	byteTen[i]= byte(v)
	//}
	//
	//rec0,err:=c1.Do("set","10byte",byteTen)
	//fmt.Println(rec0)
	//
	//str,err:=redis.String(c1.Do("get","10byte"))
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(str)

	//20byte
	//var byteTwenty [20]byte
	//for i := 0; i < 20; i++ {
	//	v:=i
	//	byteTwenty[i]= byte(v)
	//}
	//
	//rec0,err:=c1.Do("set","20byte",byteTwenty)
	//fmt.Println(rec0)
	//
	//str,err:=redis.String(c1.Do("get","20byte"))
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(str)

	//50byte
	//var byteFifty [50]byte
	//for i := 0; i < 50; i++ {
	//	v:=i
	//	byteFifty[i]= byte(v)
	//}
	//
	//rec0,err:=c1.Do("set","50byte",byteFifty)
	//fmt.Println(rec0)
	//
	//str,err:=redis.String(c1.Do("get","50byte"))
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(str)

	//100byte
	var bytess [5000]byte
	for i := 0; i < 5000; i++ {
		v := i
		bytess[i] = byte(v)
	}
	for i := 0; i < 10000; i++ {
		rec0, _ := c1.Do("set", fmt.Sprintf("bytes_no_%s", string(i)), bytess)
		fmt.Println(rec0)
	}

}
