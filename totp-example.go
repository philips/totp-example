// Implements an example command line application using the time based
// token library by balasanjay.
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/balasanjay/totp"
)

func main() {
	img, _ := totp.BarcodeImage("testing-totp", []byte("foobar"), nil)
	ioutil.WriteFile("barcode.png", img, 0777)
	fmt.Println("Barcode image written to barcode.png")
	fmt.Println("Enter your token and I will tell you whether it is right")

	in := bufio.NewReader(os.Stdin)
	input := ""

	for input != "q" {
		input, _ := in.ReadString('\n')
		input = strings.TrimRight(input, "\n")
		ok := totp.Authenticate([]byte("foobar"), input, nil)
		fmt.Println(fmt.Sprintf("%v", ok))
	}
}
