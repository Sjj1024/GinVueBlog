package utils

import (
	"log"
	"math/rand"
	"time"
)

//GenerateRand 生成随机字符串
func GenerateRand(r int) string {
	// 使用随机种子
	rand.Seed(time.Now().UnixNano())
	var Source = "sdgalskowenskdfahsdhfashkahsdkh"
	var result []byte
	for i := 0; i < r; i++ {
		r = rand.Intn(len(Source))
		result = append(result, Source[r])
	}
	var res = string(result)
	log.Println("生成的随机数是：", res)
	return res
}
