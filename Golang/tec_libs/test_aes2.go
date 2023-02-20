package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func Decrypt(key []byte, encrypted string) ([]byte, error) {
	ciphertext, err := base64.RawURLEncoding.DecodeString(encrypted)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

func Encrypt(key, data []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return base64.RawURLEncoding.EncodeToString(ciphertext), nil
}

// IsExists 判断所给路径文件/文件夹是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
}

// IfNoFileToCreate 文件不存在就创建文件
func IfNoFileToCreate(fileName string) (file *os.File) {
	var f *os.File
	var err error
	if !IsExists(fileName) {
		f, err = os.Create(fileName)
		if err != nil {
			return
		}
		log.Printf("IfNoFileToCreate 函数成功创建文件:%s", fileName)
		defer f.Close()
	}
	return f
}
func WriteStringToFileMethod1(fileName string, writeInfo string) {
	_ = IfNoFileToCreate(fileName)
	info := []byte(writeInfo)
	if err := ioutil.WriteFile(fileName, info, 0666); err != nil {
		log.Printf("WriteStringToFileMethod1 写入文件失败:%+v", err)
		return
	}
	log.Printf("WriteStringToFileMethod1 写入文件成功")
}
func main() {
	key := []byte("KEY61736466617364666173646661712")
	data := []byte("data")
	encryData, err := Encrypt(key, data)
	log.Println("Encrypt", encryData, err)

	ori, err := Decrypt(key, encryData)
	fmt.Println("Decrypt", encryData, []byte(encryData), ori, err)

	ori, err = Decrypt(key, "1MzgwmqiHyFSPSoXSB3eQuGJrFVulxaO5wfD5GoMxfo")
	log.Println("Decrypt python ", ori, string(ori), err)
	// WriteStringToFileMethod1("./data.txt", encryData)
	// 	fmt.Printf("%v -> '%s'", encryData, encryData)
	// bytes := []byte("I am byte array !")
	// str := string(bytes)
	// bytes[0] = 'i' //注意这一行，bytes在这里修改了数据，可是str打印出来的依然没变化，
	// fmt.Println(str)
}
