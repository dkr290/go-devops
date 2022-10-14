package main

import (
	"fmt"
	"log"
	"os"

	sshlib "github.com/blacknon/go-sshlib"
	"golang.org/x/crypto/ssh"
)

var (
	// Target ssh server
	localAddrPort  = "localhost:8080"
	remoteHostPort = "target.com:8501"

	remoteServer = "10.17.8.4"
	remotePort   = "22"
	user         = "danssudo"
	password2    = "password"

	termlog = "./test_termlog"
)

func main() {

	// Create sshlib.Connect
	con := &sshlib.Connect{

		// 	// If you use ssh-agent forwarding, please set to true.
		// 	// And after, run `con.ConnectSshAgent()`.
		ForwardAgent: true,
	}

	readPrivateKey, err := LoadPrivate("C:\\Users\\dans\\.ssh\\id_rsa")
	readPublicKey, err := LoadPublic("C:\\Users\\dans\\.ssh\\id_rsa.pub")

	if err != nil {
		log.Fatalln("Cannot read the private key file", err)
	}

	fmt.Println(readPrivateKey)

	// Create ssh.AuthMethod
	//authMethod := sshlib.CreateAuthMethodPassword(password2)
	authMethod, _ := sshlib.CreateAuthMethodCertificate(readPublicKey, readPrivateKey)
	// // Connect ssh server
	err = con.CreateClient(remoteServer, remotePort, user, []ssh.AuthMethod{authMethod})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// con.TCPLocalForward(localAddrPort, remoteHostPort)

}

func LoadPrivate(filepath string) (ssh.Signer, error) {
	// Read the bytes of the PEM file, e.g. id_rsa
	pemData, e := os.ReadFile(filepath)

	if e != nil {
		return nil, e
	}

	signer, err := ssh.ParsePrivateKey(pemData)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	return signer, nil
}

func LoadPublic(filepath string) (string, error) {
	// Read the bytes of the PEM file, e.g. id_rsa
	pemData, e := os.ReadFile(filepath)

	if e != nil {
		return "", e
	}

	return string(pemData), nil
}
