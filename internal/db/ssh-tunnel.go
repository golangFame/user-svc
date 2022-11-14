package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/BzingaApp/user-svc/enums"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"strconv"
	"time"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Open(s string) (_ driver.Conn, err error) {
	return pq.DialOpen(self, s)
}

func (self *ViaSSHDialer) Dial(network, address string) (net.Conn, error) {
	return self.client.Dial(network, address)
}

func (self *ViaSSHDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return self.client.Dial(network, address)
}

func DialWithPublickey(addr string, port int, user, publickeyfile string) (*ssh.Client, error) {
	key, err := ioutil.ReadFile(publickeyfile)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	client, err := ssh.Dial("tcp", addr+":"+strconv.Itoa(port), &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.HostKeyCallback(func(string, net.Addr, ssh.PublicKey) error { return nil }),
	})
	if client == nil || err != nil {
		return nil, err
	}
	client.SendRequest(user+"@"+addr, true, nil) // keep alive
	return client, nil
}

func sshTunnelTheSQL(conf *viper.Viper) {

	sshHost := conf.GetString(enums.SSH_HOST)
	sshPort := conf.GetInt(enums.SSH_PORT)
	sshUser := conf.GetString(enums.SSH_USER)
	sshPrivateKeyFilePath := conf.GetString(enums.SSH_PRIVATE_KEY_FILE_PATH)

	// Connect to the SSH Server - ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshHost, sshPort), sshConfig);
	if sshcon, err := DialWithPublickey(sshHost, sshPort, sshUser, sshPrivateKeyFilePath); err == nil {

		log.Info(string(sshcon.ServerVersion())) //FIXME using genesis.log

		sql.Register("postgres+ssh", &ViaSSHDialer{sshcon})

	} else {
		fmt.Printf("ssh connection failed")
		panic(err)
	}

	return
}
