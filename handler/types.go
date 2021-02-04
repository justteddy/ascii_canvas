package handler

import (
	"time"
)

type storage interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl time.Duration) error
}

type RectangleParams struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Fill    string `json:"fill"`
	Outline string `json:"outline"`
}

type FloodFillParams struct {
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Fill string `json:"fill"`
}

type Handler struct {
	storage storage
	ttl     time.Duration
}

func New(s storage, ttl time.Duration) *Handler {
	return &Handler{
		storage: s,
		ttl:     ttl,
	}
}
