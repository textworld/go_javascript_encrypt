package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)
type TextRequest struct {
	Text string `json:"text"`
}

var privateKey = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hello gin " + strings.ToLower(context.Request.Method) + " method",
	})
}

func AESCTR(text, key, iv string) (string, error) {
	keyBs, _ := hex.DecodeString(key)
	ivBs, _ := hex.DecodeString(iv)

	ciphertext, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBs)
	if err != nil {
		return "", fmt.Errorf("failed to new cipher: %v", err)
	}

	plaintext := make([]byte, len(ciphertext))

	stream := cipher.NewCTR(block, ivBs)
	stream.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}

func GetPrivateKeyFromString(key string)(*rsa.PrivateKey, error){
	var err error
	privPem, _ := pem.Decode([]byte(key))
	if privPem == nil {
		return nil, fmt.Errorf("failed to decode private content")
	}

	var privPemBytes []byte
	if privPem.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("no rsa private key found, but got %s", privPem.Type)
	}

	privPemBytes = privPem.Bytes

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPemBytes); err != nil {
		panic(err)
	}
	privPemBytes = privPem.Bytes

	var privateKey *rsa.PrivateKey
	var ok bool
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast key to type *rsa.PrivateKey")
	}

	return privateKey, nil
}

func AesDecrypt(c *gin.Context) {
	req := TextRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{"code": 4000, "message": err.Error()})
		return
	}
	text := strings.TrimSpace(req.Text)
	if text == "" {
		c.JSON(200, gin.H{"code": 4000, "message": "empty text"})
		return
	}

	decryptText, err := AESCTR(text, "31323334353637383930313233343536", "0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f")
	if err != nil {
		c.JSON(200, gin.H{"code": 5000, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 2000, "message": "success", "data": decryptText})
}

func RsaDecrypt(c *gin.Context) {

	privateKey, err := GetPrivateKeyFromString(privateKey)
	if err != nil {
		c.JSON(200, gin.H{"code": 5000, "message": err.Error()})
		return
	}

	req := TextRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{"code": 4000, "message": err.Error()})
		return
	}
	text := strings.TrimSpace(req.Text)
	if text == "" {
		c.JSON(200, gin.H{"code": 4000, "message": "empty text"})
		return
	}
	// 使用base64进行解码
	strReader := strings.NewReader(text)
	decoder := base64.NewDecoder(base64.StdEncoding, strReader)
	encrypted, err := ioutil.ReadAll(decoder)

	des, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encrypted)
	if err != nil {
		c.JSON(200, gin.H{"code": 5000, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 2000, "message": "success", "data": string(des)})
}

