package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

const (
	keyValSeparator = ":"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func getClient() *redis.Client {
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			//Addr: "0.0.0.0:6379",
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		rdb.FlushAll(ctx)
	}
	return rdb
}

func CloseDB() {
	getClient().Close()
}

func uploadDataFor(langName string) error {
	rdata, err := getClient().HGetAll(ctx, langName).Result()
	if err != redis.Nil && len(rdata) > 0 { //data already there
		return nil
	}

	fileName := "./lang" + langName + ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("cannot upload data from file %v", fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curString := scanner.Text()
		if !strings.Contains(curString, keyValSeparator) {
			continue
		}

		curKeyValue := strings.Split(curString, keyValSeparator)
		if len(curKeyValue) < 2 {
			continue
		}

		curValue := strings.Trim(strings.TrimSpace(curKeyValue[1]), "\t")
		err = getClient().HSetNX(ctx, langName, curKeyValue[0], curValue).Err()
		if err != nil {
			return fmt.Errorf("error creating redis key-value %v-%v \n Reason: %v", curKeyValue[0], curKeyValue[1], err.Error())
		}
	}

	return nil
}

func GetValueByKey(langName, key string) (string, error) {
	err := uploadDataFor(langName)
	if err == redis.Nil || err != nil {
		return "", err
	}
	val, err := getClient().HGet(ctx, langName, key).Result()
	if err == redis.Nil || err != nil {
		return "", fmt.Errorf("error getting value for language %v key %v", langName, key)
	}
	return val, nil
}
