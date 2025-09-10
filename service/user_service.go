package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"CRUD_operation/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	CreateUser(name string, age int, email string, college string, branch string, phone string, address string, language string, projects string, about string, role string) (models.Schema, error)

	GetUser(id string) (models.Schema, error)

	UpdateUser(id string, name string, age int, email string, college string, branch string, phone string, address string, language string, projects string, about string, role string) (models.Schema, error)

	DeleteUser(id string) error
}

type userService struct {
	collection *mongo.Collection
}

func NewUserService(collection *mongo.Collection) UserService {
	return &userService{collection: collection}
}

func (s *userService) CreateUser(name string, age int, email string, college string, branch string, phone string, address string, language string, projects string, about string, role string) (models.Schema, error) {
	user := models.Schema{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Age:       age,
		Email:     email,
		College:   college,
		Branch:    branch,
		Phone:     phone,
		Address:   address,
		Language:  language,
		Projects:  projects,
		About:     about,
		Role:      role,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		return models.Schema{}, err
	}

	return user, nil
}

func (s *userService) GetUser(id string) (models.Schema, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Schema{}, errors.New("invalid user ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.Schema
	err = s.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return models.Schema{}, err
	}
	return user, nil
}

func (s *userService) UpdateUser(id string, name string, age int, email string, college string, branch string, phone string, address string, language string, projects string, about string, role string) (models.Schema, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Schema{}, errors.New("invalid user ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":       name,
			"age":        age,
			"email":      email,
			"college":    college,
			"branch":     branch,
			"phone":      phone,
			"address":    address,
			"language":   language,
			"projects":   projects,
			"about":      about,
			"role":       role,
			"updated_at": time.Now().UTC(),
		},
	}

	result, err := s.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return models.Schema{}, err
	}

	if result.MatchedCount == 0 {
		return models.Schema{}, errors.New("user not found")
	}

	fmt.Println("User updated successfully")
	return s.GetUser(id)
}

func (s *userService) DeleteUser(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
