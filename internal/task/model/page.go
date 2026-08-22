package model

import (
	"encoding/base64"
	"strconv"
)

type Page struct {
	Cursor uint
	Limit  int
}

type PageResult struct {
	Tasks      []*Task
	NextCursor uint
	HasMore    bool
}

func EncodeCursor(id uint) string {
	return base64.URLEncoding.EncodeToString([]byte(strconv.FormatUint(uint64(id), 10)))
}

func DecodeCursor(raw string) (uint, error) {
	if raw == "" {
		return 0, nil
	}
	b, err := base64.URLEncoding.DecodeString(raw)
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
