package envutil

import (
	"os"
	"strconv"
	"strings"
)


func GetEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return strings.TrimSpace(value)
    }
    return defaultValue
}

func GetInt(key string, defaultValue int) int {
    strVal := GetEnv(key, "")
    if val, err := strconv.Atoi(strVal); err == nil {
        return val
    }
    return defaultValue
}
