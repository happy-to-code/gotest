package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/tjfoc/gmsm/sm4"
	"os"
)

// RSA_Encrypt RSA加密
// plainText 要加密的数据
// path 公钥匙文件地址
func RSA_Encrypt(plainText []byte, path string) []byte {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	// pem解码
	block, _ := pem.Decode(buf)
	// x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	// 对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	// 返回密文
	return cipherText
}

// RSA_Decrypt RSA解密
// cipherText 需要解密的byte数据
// path 私钥文件路径
func RSA_Decrypt(cipherText []byte, path string) []byte {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	// pem解码
	block, _ := pem.Decode(buf)
	// X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	// 返回明文
	return plainText
}

func main() {
	// 加密
	data := []byte("hello yida")
	// 随机生成一个字符串作为sm4 key
	sm4key := []byte("1q2w3a4r5t6y7u8i")
	// 使用sm4加密
	sm4mw := sm4Encrypt(data, sm4key)
	fmt.Println("sm4加密后的密文：", string(sm4mw))

	// 用rsa的公钥对sm4的公钥进行加密
	sm4Keyencrypt := RSA_Encrypt(sm4key, "D:\\200707Go\\gotest\\src\\MyTest\\public.pem")
	fmt.Println("RSA对sm4Key公钥进行加密后文件：", string(sm4Keyencrypt))
	decrypt := RSA_Decrypt(sm4Keyencrypt, "D:\\200707Go\\gotest\\src\\MyTest\\private.pem")
	fmt.Println("RSA对sm4Key公钥进行解密结果：", string(decrypt))
	sm4jm := sm4Dectypt(sm4mw, decrypt)
	fmt.Println("sm4jm结果：", string(sm4jm))

	// encrypt := RSA_Encrypt(data, "D:\\200707Go\\gotest\\src\\MyTest\\public.pem")

	// base64Str := base64.StdEncoding.EncodeToString(encrypt)
	// fmt.Println("base64Str:",base64Str)
	// fmt.Println(string(encrypt))
	//
	// decodeBytes, err := base64.StdEncoding.DecodeString(base64Str)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("-->",string(decodeBytes))
	// 解密
	// decrypt := RSA_Decrypt(encrypt, "D:\\200707Go\\gotest\\src\\MyTest\\private.pem")
	// fmt.Println("解密结果：", string(decrypt))
}

// 明文数据填充
func paddingLastGroup(plainText []byte, blockSize int) []byte {
	// 1.计算最后一个分组中明文后需要填充的字节数
	padNum := blockSize - len(plainText)%blockSize
	// 2.将字节数转换为byte类型
	char := []byte{byte(padNum)}
	// 3.创建切片并初始化
	newPlain := bytes.Repeat(char, padNum)
	// 4.将填充数据追加到原始数据后
	newText := append(plainText, newPlain...)

	return newText
}

// 去掉明文后面的填充数据
func unpaddingLastGroup(plainText []byte) []byte {
	// 1.拿到切片中的最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1]
	// 2.将最后一个数据转换为整数
	number := int(lastChar)
	return plainText[:length-number]
}

func sm4Encrypt(plainText, key []byte) []byte {
	block, err := sm4.NewCipher(key)
	if err != nil {
		panic(err)
	}
	paddData := paddingLastGroup(plainText, block.BlockSize())
	iv := []byte("12345678qwertyui")
	blokMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(paddData))
	blokMode.CryptBlocks(cipherText, paddData)
	return cipherText
}

func sm4Dectypt(cipherText, key []byte) []byte {
	block, err := sm4.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678qwertyui")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(cipherText, cipherText)
	plainText := unpaddingLastGroup(cipherText)
	return plainText
}
