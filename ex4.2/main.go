package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var algo = flag.Int("algo", 256, "Type of SHA algorithm")

	flag.Parse()
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	switch *algo {
	case 256:
		fmt.Println(sha256.Sum256([]byte(stdin)))
		return
	case 384:
		fmt.Println(sha512.Sum384([]byte(stdin)))
		return
	case 512:
		fmt.Println(sha512.Sum512([]byte(stdin)))
		return

	default:
		log.Fatal("Unexpected Algo")
	}

}
