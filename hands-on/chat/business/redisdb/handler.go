package redisdb

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

func RegisterNewUser(username string, password string) error {
	err := redisClient.Set(context.Background(), username, password, 0).Err()
	if err != nil {
		log.Println("error while adding new user", err)
		return err
	}
	err = redisClient.SAdd(context.Background(), userSetKey(), username).Err()
	if err != nil {
		log.Println("error while adding user in set", err)
		redisClient.Del(context.Background(), username)
		return err
	}
	return nil
}

/** Verify if the username exists or not. */
func IsUserExist(username string) bool {
	return redisClient.SIsMember(context.Background(), userSetKey(), username).Val()
}

/** Validate if the username and password provided are correct or not. */
func IsUserAuthentic(username, password string) error {
	p := redisClient.Get(context.Background(), username).Val()
	if !strings.EqualFold(p, password) {
		return fmt.Errorf("invalid username or password")
	}
	return nil
}

func UpdateContactList(username, contact string) error {
	zs := &redis.Z{Score: float64(time.Now().Unix()), Member: contact}
	err := redisClient.ZAdd(context.Background(),
		contactListZKey(username), zs).Err()
	if err != nil {
		log.Println("error while updating contact list. username: ",
			username, "contact:", contact, err)
		return err
	}
	return nil
}
