package infra

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnMongoDB(ctx context.Context, user string, pwd string, address string, port int) (*mongo.Client, error) {
	connStr := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, pwd, address, port)
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err = conn.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return conn, nil
}
