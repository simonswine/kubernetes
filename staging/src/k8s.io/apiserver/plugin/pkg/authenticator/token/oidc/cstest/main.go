package main

import (
	"k8s.io/kubernetes/staging/src/k8s.io/apiserver/plugin/pkg/authenticator/token/oidc"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("%s <issuerURL> <bearerToken>", os.Args[0])
	}

	o, err := oidc.New(oidc.OIDCOptions{
		IssuerURL:     os.Args[1],
		UsernameClaim: "sub",
	})
	if err != nil {
		log.Fatal("error oidc: ", err)
	}
	log.Printf("oidc create success: %+v", o)

	user, valid, err := o.AuthenticateToken(os.Args[2])
	if err != nil {
		log.Fatal("error oidc validate: ", err)
	}

	if valid {
		log.Printf("oidc token valid: %+v", user)
	} else {
		log.Print("oidc token invalid")
	}

}
