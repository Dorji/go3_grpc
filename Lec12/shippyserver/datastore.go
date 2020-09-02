//Создаем mongo-сессию
package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateClient(ctx context.Context, URI string, retry int32) (*mongo.Client, error) {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err := conn.Ping(ctx, nil); err != nil {
		if retry >= 3 {
			return nil, err
		}

		retry = retry + 1
		time.Sleep(time.Second * 2)
		return CreateClient(ctx, URI, retry)
	}

	return conn, err
}
