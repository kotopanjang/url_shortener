package helper

import (
	"bufio"
	"encoding/json"

	// "crypto/rand"

	"log"
	"math/rand"
	"os"
	"strings"
)

var (
	wd = func() string {
		d, _ := os.Getwd()
		return d + "/"
	}()
)

func Deserialize(str string, res interface{}) error {
	err := json.Unmarshal([]byte(str), &res)
	return err
}

func ReadConfig() map[string]string {
	ret := make(map[string]string)
	file, err := os.Open(wd + "config/conf.conf")
	if err == nil {
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, _, e := reader.ReadLine()
			if e != nil {
				break
			}

			sval := strings.Split(string(line), "=")
			ret[sval[0]] = sval[1]
		}
	} else {
		log.Println(err.Error())
	}

	return ret
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomStr() string {
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func StringDiff(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// func EncryptAES(text, key string) (string, error) {
// 	plaintext := []byte(text)
// 	c, err := aes.NewCipher([]byte(key))
// 	if err != nil {
// 		return "", err
// 	}

// 	gcm, err := cipher.NewGCM(c)
// 	if err != nil {
// 		return "", err
// 	}

// 	nonce := make([]byte, gcm.NonceSize())
// 	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
// 		return "", err
// 	}

// 	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
// 	result := fmt.Sprintf("%x", ciphertext)
// 	return result, nil
// }

// func DecryptAES(text, key string) (string, error) {
// 	ciphertext, err := hex.DecodeString(text)
// 	if err != nil {
// 		return "", err
// 	}
// 	c, err := aes.NewCipher([]byte(key))
// 	if err != nil {
// 		return "", err
// 	}

// 	gcm, err := cipher.NewGCM(c)
// 	if err != nil {
// 		return "", err
// 	}

// 	nonceSize := gcm.NonceSize()
// 	if len(ciphertext) < nonceSize {
// 		return "", err
// 	}

// 	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
// 	res, err := gcm.Open(nil, nonce, ciphertext, nil)
// 	if err != nil {
// 		return "", err
// 	}

// 	result := fmt.Sprintf("%s", res)
// 	return result, nil
// }
