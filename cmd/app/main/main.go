package main

import (
	"fmt"

	"github.com/ark-go/arkCertbotDns/internal"
)

func main() {
	fmt.Printf("hello, world\n")

	// internal.GetSubDomain()
	// internal.GetAccountInfoReq()
	internal.GetDnsGetData("anisoftware.ru")
	//internal.SetDnsTxtData("")
}
