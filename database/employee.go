package database

import (
	"context"
	"demo/models"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Employee struct {
	Client *mongo.Client
	Dbname string
	Collection string
}

var (
	ErrNil=errors.New("nil connection")
)

//Insert into the database

func (e *Employee) Insert(ctx context.Context,employee *models.Employee) (any,error) {
		if e.Client==nil{
			return nil,ErrNil
		}

		result,err:=e.Client.Database(e.Dbname).Collection(e.Collection).InsertOne(ctx,employee)
		return	result.InsertedID,err
}


func (e *Employee) Delete(ctx context.Context, id string) (int64, error) {
	if e.Client == nil {
		return 0, ErrNil
	}
	fmt.Println(id)
	
	objid, err := primitive.ObjectIDFromHex(id)
	fmt.Println(objid)

	if err != nil {
		return 0, err
	}
	result, err := e.Client.Database(e.Dbname).Collection(e.Collection).DeleteOne(ctx, bson.D{{"_id", objid}})

	return result.DeletedCount, err
}






