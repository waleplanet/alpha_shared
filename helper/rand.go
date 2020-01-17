package helper

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var seedRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const intset = "0123456789"

func RandStringWithCharSet(length int, charset string) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seedRand.Intn(len(charset))]
	}
	return string(b)
}

func RandInt(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = intset[seedRand.Intn(len(intset))]
	}
	return string(b)
}

func RandString(length int) string {
	return RandStringWithCharSet(length, charset)
}

func GenerateToken(word string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(word), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}
