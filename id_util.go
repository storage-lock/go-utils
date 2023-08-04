package utils

import (
	"github.com/google/uuid"
	"strings"
)

// RandomID 生成一个随机的ID，如果指定了前缀的话，则拼接给定前缀到ID
func RandomID(prefix ...string) string {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")
	if len(prefix) > 0 {
		return prefix[0] + id
	} else {
		return id
	}
}
