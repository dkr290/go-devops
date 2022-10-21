package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

// Get default location of a private key
func privateKeyPath(priv string) string {

	return priv
}

// Get private key for ssh authentication
func parsePrivateKey(keyPath string) (ssh.Signer, error) {
	buff, _ := ioutil.ReadFile(keyPath)
	return ssh.ParsePrivateKey(buff)
}

// Get ssh client config for our connection
// SSH config will use 2 authentication strategies: by key and by password
func makeSshConfig(user, password, privateKey string) (*ssh.ClientConfig, error) {
	config := ssh.ClientConfig{}
	fmt.Println("Using the following private key", privateKeyPath(privateKey))
	key, err := parsePrivateKey(privateKeyPath(privateKey))

	if err != nil {
		return nil, err
	}

	if password == "" {

		config = ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(key),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	} else {
		config = ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{

				ssh.Password(password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}

	return &config, nil
}

func forward(localConn net.Conn, config *ssh.ClientConfig, serverAddrString, remoteAddrString string) {
	// Setup sshClientConn (type *ssh.ClientConn)
	sshClientConn, err := ssh.Dial("tcp", serverAddrString, config)
	if err != nil {
		log.Fatalf("ssh.Dial failed: %s", err)
	}

	// Setup sshConn (type net.Conn)
	sshConn, err := sshClientConn.Dial("tcp", remoteAddrString)
	if err != nil {
		log.Fatalln("Error establishing connection with", remoteAddrString, err)
	}
	// Copy localConn.Reader to sshConn.Writer
	go func() {
		_, err = io.Copy(sshConn, localConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()

	// Copy sshConn.Reader to localConn.Writer
	go func() {
		_, err = io.Copy(localConn, sshConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()
}

func main() {
	var (
		sshAddr       string
		localAddr     string
		remoteAddr    string
		user          string
		sshPrivateKey string
		FlagError     = errors.New("please chek all paramenter or --help for help")
	)
	flag.StringVar(&localAddr, "l", "localhost:8080", "Specify local server and port. Default is localhost:8080")
	flag.StringVar(&sshAddr, "s", "", "Specify the ssh server <ip:port> or <host:port>")
	flag.StringVar(&remoteAddr, "r", "", "remote server <host:port> or <ip:port>, it can be the same like ssh server")
	flag.StringVar(&user, "u", "", "<username>")
	flag.StringVar(&sshPrivateKey, "k", "~/.ssh/id_rsa", "<private key path and file name>")

	flag.Parse()

	if localAddr == "" || sshAddr == "" || remoteAddr == "" || user == "" || sshPrivateKey == "" {

		flag.Usage()
		fmt.Println("Example usage sshtunnel.exe or sshtunnel -l localhost:5000 -r 10.1.1.1:9999 -u username -s 10.1.1.1:22")
		log.Fatalln(FlagError)
	}

	// Build SSH client configuration
	cfg, err := makeSshConfig(user, "", sshPrivateKey)
	if err != nil {
		log.Fatalln("Error makeSshConfig", err)
	}

	// Establish connection with SSH server
	conn, err := ssh.Dial("tcp", sshAddr, cfg)
	if err != nil {
		log.Fatalln("Error establishing ssh connection", err)
	}
	defer conn.Close()

	local, err := net.Listen("tcp", localAddr)
	if err != nil {
		log.Fatalln("Error forward trafic", err)
	}
	defer local.Close()

	for {
		// Setup localConn (type net.Conn)
		localConn, err := local.Accept()
		if err != nil {
			log.Fatalf("listen.Accept failed: %v", err)
		}
		go forward(localConn, cfg, sshAddr, remoteAddr)
	}

}
