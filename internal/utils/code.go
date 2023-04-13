package utils

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

type Code struct {
	Base    string // 进制的包含字符, string类型
	Decimal uint64 // 进制长度
	Pad     string // 补位字符,若生成的code小于最小长度,则补位+随机字符, 补位字符不能在进制字符中
	Len     int    // code最小长度
}

var precision uint64 = 1e8

func NewCode() *Code {
	return &Code{
		Base:    "HVE8S2DZX9C7P5IK3MJUAR4WYLTN6BGQ",
		Decimal: 32,
		Pad:     "F",
		Len:     6,
	}
}

// IdToCode id转code
func (c *Code) IdToCode(id uint64) string {
	id += precision
	mod := uint64(0)
	res := ""
	for id != 0 {
		mod = id % c.Decimal
		id = id / c.Decimal
		res += string(c.Base[mod])
	}
	resLen := len(res)
	if resLen < c.Len {
		res += c.Pad
		for i := 0; i < c.Len-resLen-1; i++ {
			rand.Seed(time.Now().UnixNano())
			res += string(c.Base[rand.Intn(int(c.Decimal))])
		}
	}
	return res
}

// CodeToId code转id
func (c *Code) CodeToId(code string) uint64 {
	res := uint64(0)
	lenCode := len(code)

	//var baseArr [] byte = []byte(c.base)
	baseArr := []byte(c.Base)     // 字符串进制转换为byte数组
	baseRev := make(map[byte]int) // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = k
	}

	// 查找补位字符的位置
	isPad := strings.Index(code, c.Pad)
	if isPad != -1 {
		lenCode = isPad
	}

	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) == c.Pad {
			continue
		}
		index := baseRev[code[i]]
		b := uint64(1)
		for j := 0; j < r; j++ {
			b *= c.Decimal
		}
		// pow 类型为 float64 , 类型转换太麻烦, 所以自己循环实现pow的功能
		//res += float64(index) * math.Pow(float64(32), float64(2))
		res += uint64(index) * b
		r++
	}

	res -= precision
	return res
}

// InitCheck 初始化检查
func (c *Code) InitCheck() (bool, error) {
	lenBase := len(c.Base)
	// 检查进制字符
	if c.Base == "" {
		return false, errors.New("base string is nil or empty")
	}
	// 检查长度是否符合
	if uint64(lenBase) != c.Decimal {
		return false, errors.New("base length and len not match")
	}
	return true, errors.New("")
}
