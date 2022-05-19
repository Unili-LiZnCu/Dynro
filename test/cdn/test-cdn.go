package main

import (
	"fmt"

	"github.com/Unili-LiZnCu/dynro/lib/net/dns"
)

func main(){
	n , _ := dns.NSLookupH("github.com")
	
	for i := range n {
		fmt.Println(n[i])
	}
}