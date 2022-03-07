package data

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"

	utilTypes "github.com/phobos42/passgo/utils"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

func SaveData(root *utilTypes.Container) {
	var bytes = createJSONBytes(root)
	encryptedString, err := encrypt(bytes, MySecret)
	if err != nil {

	}
	saveBytesToFile(encryptedString)
}
func RetreiveData(root *utilTypes.Container) {
	var fileData = getBytesFromFile()
	textBytes, err := decrypt(fileData, MySecret)
	if err != nil {
		panic(err)
	}
	ingestJSONFromBytes(textBytes, root)
}

// func Encode(b []byte) string {
// 	return base64.StdEncoding.EncodeToString(b)
// }
// func Decode(s string) []byte {
// 	data, err := base64.StdEncoding.DecodeString(s)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return data
// }

// Encrypt method is to encrypt or hide any classified text
func encrypt(text []byte, MySecret string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		panic(err)
	}
	//plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(text))
	cfb.XORKeyStream(cipherText, text)
	return cipherText, nil
}

// Decrypt method is to extract back the encrypted text
func decrypt(cipherText []byte, MySecret string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainBytes := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainBytes, cipherText)
	return plainBytes, nil
}

func saveBytesToFile(bytes []byte) {
	var err error
	err = ioutil.WriteFile("encryptedFile", bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func getBytesFromFile() []byte {
	var err error
	var bytes []byte

	bytes, err = ioutil.ReadFile("encryptedFile")
	if err != nil {
		panic(err)
	}

	return bytes

	//var myfolders allfolders
	// err = json.Unmarshal(jsonB, &root)
	// if err != nil {
	// 	panic(err)
	// }
}
