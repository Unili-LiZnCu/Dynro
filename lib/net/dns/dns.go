package dns

import (
	// "fmt"
	// "io/ioutil"
	"net"
	// "net/http"
)

func NSLookup(name string) ([]*net.NS , error) {
	return net.LookupNS(name)
}

func NSLookupH(name string) (ret []string, err error) {
	addrs, err := net.LookupNS(name)
	if err != nil {
		return nil, err
	}
	ret = make([]string, len(addrs))
	for i := range addrs {
		ret[i] = addrs[i].Host
	}
	return ret , nil
}