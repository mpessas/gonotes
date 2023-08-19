package libs

import "github.com/GoWebProd/uuid7"

func GenerateId(prefix string) string {
	v := uuid7.New().Next().String()
	return prefix + "_" + v
}
