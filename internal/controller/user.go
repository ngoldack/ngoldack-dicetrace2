package controller

import (
	"context"
	"github.com/ngoldack/dicetrace/internal/models"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	userCollection *mongo.Collection
}

func CreateUserController(userCollection *mongo.Collection) *UserController {
	return &UserController{
		userCollection: userCollection,
	}
}

func (u *UserController) GetUser(ctx context.Context, _ int, _ int) (users []models.User, err error) {
	cursor, err := u.userCollection.Find(ctx, nil)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return
}

func (u *UserController) PostUser(ctx context.Context, username, email, name string) (user *models.User, err error) {
	user = &models.User{
		Username: "",
	}

	// Inserts the user
	res, err := u.userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// Gets the newly created user
	user, err = u.GetUserWithUserId(ctx, res.InsertedID.(primitive.ObjectID).String())
	if err != nil {
		return nil, err
	}

	return
}

func (u *UserController) GetUserWithUserId(ctx context.Context, userId string) (user *models.User, err error) {

	err = u.userCollection.FindOne(ctx, bson.D{{
		Key:   "_id",
		Value: userId,
	}}).Decode(user)

	if err != nil {
		log.Error().Err(err).Msgf("error while getting user with id '%s'", userId)
	}

	return

}
