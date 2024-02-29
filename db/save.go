package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
)

func Insert(collection string, data interface{}) error {
    client, ctx := getConnection()
    defer client.Disconnect(ctx)

    c := client.Database("webform").Collection(collection)
    
    _, err := c.InsertOne(context.Background(), data)
    return err 
}

