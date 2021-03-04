package main

import (
	"fmt"
	"plugin"
	"github.com/supermanc88/btcnet/sm2"
)

var (
	ProductObj *plugin.Plugin
)
//编写一个"构造函数"（单例模式，懒汉模式）
func new() *plugin.Plugin {
	if ProductObj == nil {
		ProductObj, _ = plugin.Open("/root/go/src/github.com/supermanc88/btcnet/sdrsbtc.so")
	}
	return ProductObj
}

func main()	{
	ppp := []byte{82, 90, 94, 93, 91, 80, 88, 75}
	p := new()
	sm3Sum256, _ := p.Lookup("Sm3Sum256")
	sum := sm3Sum256.(func(string, string, []byte, []byte) []byte)("10.20.88.223", "20002", ppp, []byte("teststring"))
	fmt.Println(sum)

	generateSm2Key, _ := p.Lookup("GenerateSm2Key")
	var priv *sm2.PrivateKey
	priv, _ = generateSm2Key.(func(ip, port string, passwd []byte) (*sm2.PrivateKey, error))("10.20.88.223", "20002", ppp)
	fmt.Println(priv)

	sm2Encrypt, _ := p.Lookup("Sm2Encrypt")
	encData, _ := sm2Encrypt.(func(ip, port string, passwd []byte, pub *sm2.PublicKey, message []byte)([]byte, error))("10.20.88.223", "20002", ppp, &priv.PublicKey, []byte("teststring"))
	fmt.Println("encdata:")
	fmt.Println(encData)

}
