package link

import (
	"gorm.io/gorm"
	"math/rand"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex:idx_hash;not null" `
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: GenerateHash(6),
	}
}

func GenerateHash(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
