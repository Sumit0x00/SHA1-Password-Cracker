package pass

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadInPassword() []string {
	b, err := ioutil.ReadFile("top-10000-passwords.txt")

	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	return strings.Split(str, "\n")

}

func ReadInSalts() []string {
	b, err := ioutil.ReadFile("known-salts.txt")

	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	return strings.Split(str, "\n")

}

func HashString(str string) string {
	bs := sha1.Sum([]byte(str))
	return fmt.Sprintf("%x", bs)

}

func HashWithSalt(str string) []string {
	var s []string

	for _, salt := range ReadInSalts() {
		s = append(s, HashString(salt+str))
		s = append(s, HashString(str+salt))
	}
	return s
}

func CrackSHA1Hash(str string, useSalt bool) string {
	for _, pass := range ReadInPassword() {
		if useSalt {
			for _, salted := range HashWithSalt(pass) {
				if salted == str {
					return pass
				}
			}
		} else {
			if HashString(pass) == str {
				return pass
			}
		}
	}
	return "PASSWORD NOT IN DATABASE"
}
