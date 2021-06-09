package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	// to consume messages
	mechanism := plain.Mechanism{
		Username: "producer",
		Password: "prod-sec@Gd~CTrH]-sV[g]h",
	}
	tlsConfig, err := NewTLSConfigC("kafka/certificate.pem", "kafka/key.pem", "kafka/CARoot.pem")
	if err != nil {
		log.Fatal("Certificate generation failed", err)
		return
	}
	// This can be used on test server if domain does not match cert:
	tlsConfig.InsecureSkipVerify = true

	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
		TLS:           tlsConfig,
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"192.168.20.31:9092"},
		GroupID: "iot_dev_group_dev_wy",
		Topic:   "iot_new_terminal",
		Dialer:  dialer,
	})
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Failed to read message", err)
			return
		}
		// TODO: process message
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}

// NewTLSConfigC NewTLSConfig generates a TLS configuration used to authenticate on server with
// certificates.
// Parameters are the three pem files path we need to authenticate: client cert, client key and CA cert.
func NewTLSConfigC(clientCertFile, clientKeyFile, caCertFile string) (*tls.Config, error) {
	tlsConfig := tls.Config{}

	// Load client cert
	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return &tlsConfig, err
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return &tlsConfig, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig.RootCAs = caCertPool
	tlsConfig.BuildNameToCertificate()
	return &tlsConfig, err
}
