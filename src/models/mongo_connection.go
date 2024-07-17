package models

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()
var client, _ = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
var mongoDB = client.Database(os.Getenv("DATABASE")) // TODO: afegir lo de las keys

type mongoModel[T any] struct {
	collection *mongo.Collection
	object     T
}

func CreateMongoModel[T any](modelName string, properties T) *mongoModel[T] {
	return &mongoModel[T]{
		collection: mongoDB.Collection(modelName),
		object:     properties,
	}
}

func (m *mongoModel[T]) setUniqueField(uniqueFields ...string) *mongoModel[T] {
	length := len(uniqueFields)

	if length > 1 {
		indexModels := []mongo.IndexModel{}

		for _, field := range uniqueFields {
			indexModels = append(indexModels, mongo.IndexModel{
				Keys:    bson.D{{Key: field, Value: 1}},
				Options: options.Index().SetUnique(true),
			})
		}

		m.collection.Indexes().CreateMany(ctx, indexModels)
	} else {
		m.collection.Indexes().CreateOne(
			ctx,
			mongo.IndexModel{
				Keys:    bson.D{{Key: uniqueFields[0], Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		)
	}

	return m
}

// Returns id, body and error
func (m *mongoModel[T]) FindOneAndId(filter bson.M, projection ...bson.M) (string, *T, error) { // OTHER: alomillor retornar un puntero a T
	var (
		res        bson.M
		returnable T
		err        error
	)

	if len(projection) >= 1 {
		projectionOptions := options.FindOne().SetProjection(projection[0])
		err = m.collection.FindOne(ctx, filter, projectionOptions).Decode(&res)
	} else {
		err = m.collection.FindOne(ctx, filter).Decode(&res)
	}

	if err != nil {
		return "", &returnable, err
	}

	byteRes, err := bson.Marshal(&res)

	if err != nil {
		return "", &returnable, err
	}
	bson.Unmarshal(byteRes, &returnable)

	return res["_id"].(primitive.ObjectID).Hex(), &returnable, nil
}

func (m *mongoModel[T]) FindOne(filter bson.M, projection ...T) (*T, error) {
	var (
		res T
		err error
	)

	if len(projection) >= 1 {
		projectionOptions := options.FindOne().SetProjection(projection[0])
		err = m.collection.FindOne(ctx, filter, projectionOptions).Decode(&res)
	} else {
		err = m.collection.FindOne(ctx, filter).Decode(&res)
	}

	return &res, err
}

func (m *mongoModel[T]) InsertOne(newDocument T) (primitive.ObjectID, error) {
	res, err := m.collection.InsertOne(ctx, newDocument)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), err
}

func (m *mongoModel[T]) UpdateOne(find bson.M, update bson.M) (*mongo.UpdateResult, error) {
	return m.collection.UpdateOne(ctx, find, primitive.M{"$set": update})
}

func (m *mongoModel[T]) UpdateMany(find bson.M, update bson.M) (*mongo.UpdateResult, error) {
	return m.collection.UpdateMany(ctx, find, update)
}

func (m *mongoModel[T]) DeleteOne(find bson.M) (*mongo.DeleteResult, error) {
	return m.collection.DeleteOne(ctx, find)
}

func (m *mongoModel[T]) DeleteMany(find bson.M) (*mongo.DeleteResult, error) {
	return m.collection.DeleteMany(ctx, find)
}
