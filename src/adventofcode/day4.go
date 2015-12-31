package adventofcode

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func ExecuteDay4(input string) {
	result := 0
	var data []byte
	var ashex string

	for {
		data = []byte(fmt.Sprintf("%v%v", input, result))
		//		hashed := md5.Sum(data)
		hasher := md5.New()
		hasher.Write(data)
		//		hashed := md5.Sum(data)
		ashex = hex.EncodeToString(hasher.Sum(nil))

		if result%10000000 == 0 {
			fmt.Printf("Data: %v, Hashed: %v\n", string(data), ashex)
		}

		if ashex[0:6] == "000000" {
			break
		}
		result++
	}

	fmt.Printf("Data: %v, Hashed: %v\n", string(data), ashex)
	fmt.Printf("Result: %v\n", result)
}
