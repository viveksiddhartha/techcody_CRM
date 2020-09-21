package datastore

import (
	"SVCRM/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/mediocregopher/radix.v2/pool"
)

type RedisDatastore struct {
	*pool.Pool
}

func NewRedisDatastore() (*RedisDatastore, error) {

	connectionPool, err := pool.New("tcp", "127.0.0.1:6379", 10)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection Successful-Redis")

	return &RedisDatastore{
		Pool: connectionPool,
	}, nil

}

func CreateEntityRedis(entity *models.CoEntity) error {

	r, err := NewRedisDatastore()
	if err != nil {
		return err
	}

	entityJSON, err := json.Marshal(*entity)
	if err != nil {
		return err
	}

	if r.Cmd("SET", "CoEntityID:"+entity.CoEntityId, string(entityJSON)).Err != nil {
		return errors.New("Failed to execute Redis SET command")
	}

	return nil
}

func GetUserRedis(username string) (*models.CoEntity, error) {
	r, err := NewRedisDatastore()
	if err != nil {
		return nil, err
	}
	fmt.Println("BValue in GetUserRedis %v \n", username)

	exists, err := r.Cmd("EXISTS", "CoEntityID:"+username).Int()

	if err != nil {
		return nil, err
	} else if exists == 0 {
		return nil, nil
	}

	var u models.CoEntity

	entityJSON, err := r.Cmd("GET", "CoEntityID:"+username).Str()

	fmt.Println("userJSON: ", entityJSON)

	if err != nil {
		log.Print(err)

		return nil, err
	}

	if err := json.Unmarshal([]byte(entityJSON), &u); err != nil {
		log.Print(err)
		return nil, err
	}

	return &u, nil
}

func (r *RedisDatastore) Close() {
	r.Close()
}
