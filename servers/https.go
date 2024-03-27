package servers

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	fib "github.com/gofiber/fiber/v2"
)

func ListenHTTPSServer(app *fib.App, address string, serverDesc string, httpsSecretName string) error {
	//LOAD CERTIFICATES
	path := fmt.Sprintf("/tls/%v/tls.crt", httpsSecretName)
	fmt.Printf(path)
	cerr, err := tls.LoadX509KeyPair(path, fmt.Sprintf("/tls/%v/tls.key", httpsSecretName))
	if err != nil {
		fmt.Printf("[%v] Server SSL Fatal error reading certs", serverDesc)
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cerr}}

	//SERVER
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("%s : %v", serverDesc, err))
	}
	lntls := tls.NewListener(ln, cfg)
	fmt.Printf("[%v] API server HTTPS listen to %v\n", serverDesc, address)
	return app.Listener(lntls)
}
