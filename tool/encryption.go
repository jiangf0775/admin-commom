package tool

import (
	"crypto/md5"
	"fmt"
	"github.com/spaolacci/murmur3"
	"io"
)

/** 加密方式 **/

func Md5String(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		panic(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func Md5Bytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func Hash265(value string) string {
	// 初始化哈希器，指定种子（可选，默认为0）
	hash32 := murmur3.New32()

	// 将数据喂入哈希器
	hash32.Write([]byte(value))

	// 获取哈希值
	return fmt.Sprintf("%x", hash32.Sum32())
}
