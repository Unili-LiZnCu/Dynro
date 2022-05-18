package main

import (
	"github.com/Unili-LiZnCu/dynro/lib/cert"
)

func main(){
	err := cert.GenerateCert()
	if err != nil {
		panic(err)
	}
}
