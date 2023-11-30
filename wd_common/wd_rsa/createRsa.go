package wd_rsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"wd_common/wd_database"
)

var (
	PrivateKey *ecdsa.PrivateKey
	BucketName = "private"
	BucketKey  = "key"
)

func restorePrivateKey() {
	db := wd_database.ConnectPrivateVoltDB()
	var privateKey []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		if bucket == nil {
			return nil
		}
		privateKey = bucket.Get([]byte(BucketKey))
		return nil
	})
	fmt.Println(privateKey, err)

	if len(privateKey) == 0 || err != nil {
		pk, data, err := marshalPrivateKey()
		if err != nil {
			log.Fatal("fail to marshaling private key")
			return
		}
		if err := savePrivateKey(data); err != nil {
			log.Fatal("fail to restore private key")
			return
		}
		PrivateKey = pk
	} else {
		if err := parsePrivateKey(privateKey); err != nil {
			log.Fatal("fail to parsing private key")
		}
		return
	}
}
func marshalPrivateKey() (*ecdsa.PrivateKey, []byte, error) {
	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("fail to generate private, public key pair %+v", err)
		return nil, nil, err
	}
	data, err := x509.MarshalECPrivateKey(pk)
	if err != nil {
		log.Fatal("fail to convert private key")
		return nil, nil, err
	}

	return pk, data, nil
}

func parsePrivateKey(data []byte) error {
	key, err := x509.ParseECPrivateKey(data)
	if err != nil {
		return err
	}
	PrivateKey = key
	return nil
}

func savePrivateKey(data []byte) error {
	db := wd_database.ConnectPrivateVoltDB()

	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			return err
		}
		if err := bucket.Put([]byte(BucketKey), data); err != nil {
			return err
		}
		return nil
	})
}

func GetPrivateKey() *ecdsa.PrivateKey {
	restorePrivateKey()
	return PrivateKey
}
