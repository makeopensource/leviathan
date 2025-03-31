package docker

// ssh utils for docker connection

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	com "github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"time"
)

func publicKeyToString(publicKey ssh.PublicKey, comment string) (string, error) {
	authorizedKey := ssh.MarshalAuthorizedKey(publicKey)
	if comment != "" {
		return string(authorizedKey[:len(authorizedKey)-1]) + " " + comment, nil
	}
	return string(authorizedKey), nil
}

func stringToPublicKey(publicKeyString string) (ssh.PublicKey, error) {
	publicKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKeyString))
	if err != nil {
		return nil, fmt.Errorf("unable to parse public key %v", err)
	}
	return publicKey, nil
}

// sshDialer custom dialer that uses the SSH connection.
func sshDialer(sshClient *ssh.Client) func(ctx context.Context, network string, addr string) (net.Conn, error) {
	return func(ctx context.Context, network string, addr string) (net.Conn, error) {
		dockerUnix := "/var/run/docker.sock"
		dial, err := sshClient.Dial("unix", dockerUnix)
		if err == nil {
			return dial, nil
		}

		dockerTcp := "127.0.0.1:2375"
		log.Warn().
			Err(err).
			Msgf("failed to dial remote docker client at %s, using fallingback at %s", dockerUnix, dockerTcp)

		return sshClient.Dial("tcp", dockerTcp)
	}
}

func saveHostKey(machine models.MachineOptions) func(hostname string, remote net.Addr, key ssh.PublicKey) error {
	return func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		log.Debug().Msgf("Empty public key for %s, public key will be saved on connect", machine.Name())

		comment := fmt.Sprintf("added by leviathan for machine %s on %s", machine.Name(), time.Now().String())
		stringKey, err := publicKeyToString(key, comment)
		if err != nil {
			return fmt.Errorf("unable to convert public key for machine %s ", machine.Name())
		}

		machine.RemotePublickey = stringKey
		writeMachineToConfigFile(machine)
		return nil
	}
}

func writeMachineToConfigFile(machine models.MachineOptions) {
	machineKey := fmt.Sprintf("%s.%s", com.ClientsSSH.ConfigKey, machine.Name())
	viper.Set(machineKey, machine)
	err := viper.WriteConfig()
	if err != nil {
		log.Warn().Err(err).Msgf("failed to update machine %s public key to config", machine.Name())
	}
}

// GenerateKeyPair creates a new SSH key pair
func GenerateKeyPair() (privateKey []byte, publicKey []byte, err error) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Convert private key to PEM format
	privateKeyDER := x509.MarshalPKCS1PrivateKey(rsaKey)
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyDER,
	}
	privatePEM := pem.EncodeToMemory(&privateKeyBlock)

	// Generate public key from private key
	rsaPublicKey, err := ssh.NewPublicKey(&rsaKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	pubKeyBytes := ssh.MarshalAuthorizedKey(rsaPublicKey)

	return privatePEM, pubKeyBytes, nil
}

// initKeyPairFile creates RSA key-pair files,
// if they do not exist, otherwise skips generation
//
// the generated keys can be found in common.SSHConfigFolder
func initKeyPairFile() {
	basePath := com.SSHConfigFolder.GetStr()
	privateKeyPath := fmt.Sprintf("%s/%s", basePath, "id_rsa")
	publicKeyPath := fmt.Sprintf("%s/%s", basePath, "id_rsa.pub")

	defer log.Info().
		Msgf("to add the public key to other hosts use\nssh-copy-id -i %s <user>@<remote_host>\n", publicKeyPath)

	logF := log.Info().
		Str("private_key_file", privateKeyPath).
		Str("public_key_file", publicKeyPath)

	if com.FileExists(privateKeyPath) && com.FileExists(publicKeyPath) {
		logF.Msg("found existing keys... skipping generation")
		return
	}

	privateKey, publicKey, err := GenerateKeyPair()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate key pair")
	}

	if err := os.WriteFile(privateKeyPath, privateKey, 0600); err != nil {
		log.Fatal().Err(err).Msg("Failed to save private key")
	}
	if err := os.WriteFile(publicKeyPath, publicKey, 0644); err != nil {
		log.Fatal().Err(err).Msg("Failed to save public key")
	}

	logF.Msg("Generated new SSH key pair")
}

func LoadPrivateKey() ([]byte, error) {
	return os.ReadFile(fmt.Sprintf(
		"%s/%s",
		com.SSHConfigFolder.GetStr(),
		"id_rsa",
	))
}
