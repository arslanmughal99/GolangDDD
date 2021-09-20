package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"massivleads/domain/entity"
	serrors "massivleads/prototypes/errors"
)

const collection = "users"

type UserRepository struct {
	client *mongo.Client
}

// NewMongoUserRepository Create a new mongodb users repository
func NewMongoUserRepository() *UserRepository {
	repo := new(UserRepository)
	repo.client = newMongoClient()

	// Create mongo collection indexes
	idxs := []mongo.IndexModel{
		{
			Keys:    bson.D{{"username", 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{"email", 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := repo.client.Database(database).Collection(collection).Indexes().CreateMany(context.TODO(), idxs)

	if err != nil {
		panic("Failed to create user collection indexes")
	}

	return repo
}

func (r *UserRepository) CreateUser(user entity.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	_, err := r.client.Database(database).Collection(collection).InsertOne(ctx, &user)

	if mongo.IsDuplicateKeyError(err) {
		return serrors.ErrUserExist
	}

	return err
}

func (r *UserRepository) GetUserById(id string) (*entity.User, error) {
	user := new(entity.User)
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	err := r.client.Database(database).Collection(collection).FindOne(ctx, bson.M{"_id": id}).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByUsernameOrEmail(username string) (*entity.User, error) {
	user := new(entity.User)
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	err := r.client.Database(database).Collection(collection).FindOne(
		ctx,
		bson.D{
			{
				"$or", bson.A{
					bson.D{{"email", username}},
					bson.D{{"username", username}},
				},
			},
		},
	).Decode(user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {

			return nil, serrors.ErrUserNotExist
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*entity.User, error) {
	user := new(entity.User)
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	err := r.client.Database(database).Collection(collection).FindOne(
		ctx,
		bson.D{{"username", username}},
	).Decode(user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {

			return nil, serrors.ErrUserNotExist
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(user entity.User) error {
	upsert := true
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	_, err := r.client.Database(database).Collection(collection).UpdateOne(
		ctx, bson.M{"_id": user.ID}, user, &options.UpdateOptions{Upsert: &upsert},
	)

	return err
}
