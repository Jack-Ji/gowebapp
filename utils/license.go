package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"runtime"
	"time"

	"github.com/denisbrodbeck/machineid"
)

type SystemInfo struct {
	OS      string
	ARCH    string
	CPU     int
	MAC     string
	MACHINE string
}

type UseInfo struct {
	GenTime    string
	ExpireTime string
	ProjNum    int
	UserNum    int
}

type License struct {
	Sys SystemInfo
	Use UseInfo
}

var privatePEM = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`

var privateKey *rsa.PrivateKey

func init() {
	block, _ := pem.Decode([]byte(privatePEM))
	if block == nil {
		panic(fmt.Sprintf("decode private pem failed"))
	}

	var err error
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("parse private key failed: %s", err))
	}
}

func GetSysInfo() *SystemInfo {
	info := SystemInfo{}
	info.OS = runtime.GOOS
	info.ARCH = runtime.GOARCH
	info.CPU = runtime.NumCPU()
	_, info.MAC = GetIP()
	id, err := machineid.ProtectedID("langmy.com")
	if err != nil {
		panic(err)
	}
	info.MACHINE = id
	return &info
}

func EncryptData(data []byte) ([]byte, error) {
	bs, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privateKey.PublicKey, data, nil)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func DecryptData(data []byte) ([]byte, error) {
	bs, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, data, nil)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func LoadLicense(rd io.Reader) *License {
	ebs := make([]byte, 10240)
	s, err := io.ReadFull(rd, ebs)
	if err != nil && err != io.ErrUnexpectedEOF {
		panic(fmt.Sprintf("read license data failed: %s", err))
	}
	if 0 == len(ebs) || s == len(ebs) {
		panic("abnormal license size!")
	}

	bs, err := DecryptData(ebs[:s])
	if err != nil {
		panic(fmt.Sprintf("decrypt license data failed: %s", err))
	}

	lic := License{}
	err = json.Unmarshal(bs, &lic)
	if err != nil {
		panic(fmt.Sprintf("decode license data failed: %s", err))
	}

	return &lic
}

func GenLicense(hdinfo []byte, expireDate time.Time, projNum, userNum int) ([]byte, error) {
	now := time.Now()
	if now.After(expireDate) {
		return nil, fmt.Errorf("invalid expire date: %s", expireDate)
	}

	bs, err := DecryptData(hdinfo)
	if err != nil {
		return nil, err
	}

	sys := SystemInfo{}
	err = json.Unmarshal(bs, &sys)
	if err != nil {
		return nil, err
	}

	lic := License{}
	lic.Sys = sys
	lic.Use.GenTime = now.Format(time.RFC3339)
	lic.Use.ExpireTime = expireDate.Format(time.RFC3339)
	lic.Use.ProjNum = projNum
	lic.Use.UserNum = userNum

	bs, err = json.MarshalIndent(&lic, "", "  ")
	if err != nil {
		return nil, err
	}

	ebs, err := EncryptData(bs)
	if err != nil {
		return nil, err
	}

	fmt.Println("successfully generated license:")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(string(bs))
	fmt.Println("--------------------------------------------------------------------------------")
	return ebs, nil
}
