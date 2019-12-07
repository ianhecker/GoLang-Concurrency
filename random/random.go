package random

import (
	"fmt"
	"math/rand"
	"os"
)

func RandomFile(count int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < count; i++ {
		num := rand.Intn(count)
		_, err := file.WriteString(fmt.Sprintf("%d,", num))
		if err != nil {
			panic(err)
		}
	}
}
