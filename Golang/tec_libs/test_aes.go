package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func Decrypt(ikey, encrypted string) (string, error) {
	key := []byte(ikey)
	ciphertext := []byte(encrypted)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	fmt.Println("Decrypt in", ciphertext, string(ciphertext))
	return string(ciphertext), nil
}

func Encrypt(ikey, idata string) (string, error) {
	key := []byte(ikey)
	data := []byte(idata)
	// fmt.Print(key, len(key), len(data))
	block, _ := aes.NewCipher(key)
	// if err != nil {
	// 	return "", err
	// }
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	fmt.Println("Encrypt in", ciphertext, string(ciphertext))
	return string(ciphertext), nil
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
	encryData, err := Encrypt("KEY61736466617364666173646661712", "data")
	log.Println("Encrypt", encryData, err)

	ori, err := Decrypt("KEY61736466617364666173646661712", encryData)
	fmt.Println("Decrypt", encryData, []byte(encryData), ori, err)
	// WriteStringToFileMethod1("./data.txt", encryData)
	fmt.Printf("array: %s %v (%T)", encryData, encryData, encryData)
	fmt.Printf("%v -> '%s'", encryData, encryData)
	// bytes := []byte("I am byte array !")
	// str := string(bytes)
	// bytes[0] = 'i' //注意这一行，bytes在这里修改了数据，可是str打印出来的依然没变化，
	// fmt.Println(str)
}
