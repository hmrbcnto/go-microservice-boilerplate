package user_repo

import (
	"context"
	"time"

	"github.com/hmrbcnto/go-net-http/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	CreateUser(*entities.User) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
	GetUserById(userId string) (*entities.User, error)
}

type userRepo struct {
	db *mongo.Collection
}

func NewRepo(db *mongo.Client) UserRepo {
	return &userRepo{
		db: db.Database("leniApi").Collection("users"),
	}
}

func (ur *userRepo) CreateUser(user *entities.User) (*entities.User, error) {

	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	// Creating user
	insertionResult, err := ur.db.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	// Getting inserted data
	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}

	// Query
	createdRecord := ur.db.FindOne(ctx, filter)

	// Decode to user entity
	createdUser := &entities.User{}
	createdRecord.Decode(createdUser)

	// returning
	return createdUser, nil
}

func (ur *userRepo) GetAllUsers() ([]entities.User, error) {

	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	/// Getting all users
	// Creating query
	query := bson.D{{}}

	cursor, err := ur.db.Find(ctx, query)

	if err != nil {
		return nil, err
	}

	var users []entities.User

	// Iterate and decode
	err = cursor.All(ctx, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepo) GetUserById(userId string) (*entities.User, error) {

	// Creating context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()
	// Creating query
	query := bson.D{{Key: "_id", Value: userId}}

	// Querying
	result := ur.db.FindOne(ctx, query)

	// Converting to user entity
	foundUser := new(entities.User)
	result.Decode(foundUser)

	// Returning
	return foundUser, nil
}
