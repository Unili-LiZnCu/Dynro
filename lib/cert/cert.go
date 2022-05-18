package cert


import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func GenerateCert() error {
	key , err := rsa.GenerateKey(rand.Reader , 2048)
	if err != nil {
		return err
	}
	sn , err := rand.Int(rand.Reader , big.NewInt(10))
	if err != nil {
		return err
	}
	nm := pkix.Name{
		Country:            []string{"CN"},
		Organization:       []string{"DYNRO"},
		OrganizationalUnit: []string{"DYNRO"},
		Province:           []string{"Rhodes Island"},
		CommonName:         "DYNRO",
		Locality:           []string{"Rhodes Island"},
	}
	temp := x509.Certificate{
		SerialNumber: sn,
		Issuer:       nm,
		Subject:      nm,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Hour * 24 * 365 * 10),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign | x509.KeyUsageKeyAgreement | x509.KeyUsageContentCommitment | x509.KeyUsageDataEncipherment ,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		BasicConstraintsValid: true,
		IsCA:                  true,
		IPAddresses: []net.IP{[]byte("127.0.0.1")},
	}
	cert , err := x509.CreateCertificate(rand.Reader , &temp ,&temp , &key.PublicKey , key)
	if err != nil {
		return err
	}
	block := pem.Block{
		Type:    "CERTIFICATE",
		Headers: nil,
		Bytes:   cert,
	}
	file, err := os.Create("cert/DYNRO.crt")
	if err != nil {
		return err
	}
	defer file.Close()
	pem.Encode(file, &block)
	block = pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(key),
	}
	file, err = os.Create("cert/DYNRO.key")
	if err != nil {
		return err
	}
	pem.Encode(file, &block)
	file, err = os.Create("cert/README.txt")
	file.Write([]byte("自签名证书生成完成，请保护好您的证书，不要泄露给他人,证书文件名为DYNRO.crt,密钥文件名为DYNRO.key，RSA密钥为随机生成，删除后无法找回。\n"))
	file.Write([]byte("请双击.crt文件，并将其安装到“受信任的根证书颁发机构”下。否则DYNRO可能无法正常运行。\n"))
	file.Write([]byte("该证书生成于"+time.Now().String()+"\n"))
	return nil
}