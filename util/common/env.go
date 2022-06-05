package common

import (
	"log"
	"os"
	"strconv"
)

func Env(key, defval string) string {
	v := os.Getenv(key)
	if v == "" {
		return defval
	}

	return v
}

func EnvInt(key string, defval int) int {
	v := os.Getenv(key)
	if v == "" {
		return defval
	}

	iv, err := strconv.Atoi(v)
	if err != nil {
		log.Panicln(err)
	}

	return iv
}

func EnvFloat(key string, defval float64) float64 {
	v := os.Getenv(key)
	if v == "" {
		return defval
	}

	iv, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Panicln(err)
	}

	return iv
}
