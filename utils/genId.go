package utils

import "math/rand"

func GenerateId(length int, chars string) string {
	var id string = ""
	for i := 0; i <= length-1; i++ {
		num := rand.Intn(len(chars))
		id = id + string(chars[num])
	}

	return id
}
