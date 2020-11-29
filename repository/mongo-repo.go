package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kennykim11/study-buddy-backend/graph/model"
)

type Repository interface {
	SaveUser(ctx context.Context, user *model.User)
	FindUser(ctx context.Context, googleid int) *model.User
	EnrollInSection(ctx context.Context, userID int, sectionID int)
	RemoveFromSection(ctx context.Context, userID int, sectionID int)
	RegisterContactInfo(ctx context.Context, userID int, contact model.ContactInput)
	FindSection(ctx context.Context, sectionID int) *model.Section
}

type Database struct {
	client *mongo.Client
}

func GetCollection(db *Database, name string) *mongo.Collection {
	return db.client.Database("studybuddy").Collection(name)
}

func New() Repository {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return &Database{
		client: db,
	}
}

func (db *Database) SaveUser(ctx context.Context, user *model.User) {
	collection := GetCollection(db, "users")
	_, err := collection.InsertOne(ctx, user) //passes context, error
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Database) FindUser(ctx context.Context, googleid int) *model.User {
	collection := GetCollection(db, "users")
	var result model.User
	err := collection.FindOne(ctx, bson.M{"googleid": googleid}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return &result
}

func (db *Database) EnrollInSection(ctx context.Context, userID int, sectionID int) {
	collection := GetCollection(db, "users")
	_, err := collection.UpdateOne(ctx,
		bson.D{{"googleid", userID}},
		bson.D{
			{"$addToSet", bson.D{
				{"taking", sectionID},
			}},
		})
	if err != nil {
		log.Fatal(err)
	}
	collection = GetCollection(db, "sections")
	_, err = collection.UpdateOne(ctx,
		bson.D{{"sectionid", sectionID}},
		bson.D{
			{"$addToSet", bson.D{
				{"userstaking", userID},
			}},
		})
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Database) RemoveFromSection(ctx context.Context, userID int, sectionID int) {
	collection := GetCollection(db, "users")
	_, err := collection.UpdateOne(ctx,
		bson.D{{"googleid", userID}},
		bson.D{
			{"$pull", bson.D{
				{"taking", sectionID},
			}},
		})
	if err != nil {
		log.Fatal(err)
	}
	collection = GetCollection(db, "sections")
	_, err = collection.UpdateOne(ctx,
		bson.D{{"sectionid", sectionID}},
		bson.D{
			{"$pull", bson.D{
				{"userstaking", userID},
			}},
		})
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Database) RegisterContactInfo(ctx context.Context, userID int, contact model.ContactInput) {
	collection := GetCollection(db, "users")
	_, err := collection.UpdateOne(ctx,
		bson.D{{"googleid", userID}},
		bson.D{
			{"$pull", bson.D{
				{"contacts", contact},
			}},
		})
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Database) FindSection(ctx context.Context, sectionID int) *model.Section {
	collection := GetCollection(db, "sections")
	var result model.Section
	err := collection.FindOne(ctx, bson.M{"sectionid": sectionID}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return &result
}
