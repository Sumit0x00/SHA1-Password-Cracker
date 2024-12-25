package main

import (
	"fmt"

	"github.com/Sumit0x00/SHA1-Password-Cracker/pass"
)

func main() {
	str := pass.CrackSHA1Hash("18c28604dd31094a8d69dae60f1bcd347f1afc5a", false)
	fmt.Println(str) // superman

	str = pass.CrackSHA1Hash("53d8b3dc9d39f0184144674e310185e41a87ffd5", true)
	fmt.Println(str) // superman
}
