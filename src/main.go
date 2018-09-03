package main

import (
	"unicode/utf8"
	"wxApi"
	"dbServer"
	"fmt"
	"time"
)

func main() {

	//配置文件初始化
	appId := "wx293dbb0f011bcac3"
	mysqlApi := dbServer.CreateMysqlApi()
	appSec, _ := mysqlApi.GetWxApp(appId)
	//fmt.Print(appSec)
	//y := "afsaafafasfafa"
	//x := "sfa"
	//r := Count(y, x)
	//fmt.Print(r)
	fans := wxApi.Fans{}
	fmt.Print(appSec)
	fans.Refresh(appId, appSec, "")
	openIds := fans.Data.Openid
	for i := 100; i < 100; i++ {
		go mysqlApi.SaveOpenIds(appId, openIds[i])
	}
	time.Sleep(time.Minute)
}

// primeRK 是用于 Rabin-Karp 算法中的素数，也就是上面说的 M
const primeRK = 16777619

// 返回 Rabin-Karp 算法中“搜索词” sep 的“哈希值”及相应的“乘数因子（权值）”
func hashstr(sep string) (uint32, uint32) {
	// 计算 sep 的 hash 值
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	// 计算 sep 最高位 + 1 位的权值 pow（乘数因子）
	// 也就是上面说的 M 的 n 次方
	// 这里通过遍历 len(sep) 的二进制位来计算，减少计算次数
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 { // 如果二进制最低位不是 0
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

// Count 计算字符串 sep 在 s 中的非重叠个数
// 如果 sep 为空字符串，则返回 s 中的字符(非字节)个数 + 1
// 使用 Rabin-Karp 算法实现
func Count(s, sep string) int {
	n := 0
	// 特殊情况判断
	switch {
	case len(sep) == 0: // 空字符，返回字符个数 + 1
		return utf8.RuneCountInString(s) + 1
	case len(sep) == 1: // 单个字符，可以用快速方法
		c := sep[0]
		for i := 0; i < len(s); i++ {
			if s[i] == c {
				n++
			}
		}
		return n
	case len(sep) > len(s):
		return 0
	case len(sep) == len(s):
		if sep == s {
			return 1
		}
		return 0
	}
	// 计算 sep 的 hash 值和乘数因子
	hashsep, pow := hashstr(sep)
	// 计算 s 中要进行比较的字符串的 hash 值
	h := uint32(0)
	for i := 0; i < len(sep); i++ {
		h = h*primeRK + uint32(s[i])
	}
	lastmatch := 0 // 下一次查找的起始位置，用于确保找到的字符串不重叠
	// 找到一个匹配项（进行一次朴素比较）
	if h == hashsep && s[:len(sep)] == sep {
		n++
		lastmatch = len(sep)
	}
	// 滚动 s 的 hash 值并与 sep 的 hash 值进行比较
	for i := len(sep); i < len(s); {
		// 加上下一个字符的 hash 值
		h *= primeRK
		h += uint32(s[i])
		// 去掉第一个字符的 hash 值
		h -= pow * uint32(s[i-len(sep)])
		i++
		// 开始比较
		// lastmatch <= i-len(sep) 确保不重叠
		if h == hashsep && lastmatch <= i-len(sep) && s[i-len(sep):i] == sep {
			n++
			lastmatch = i
		}
	}
	return n
}
