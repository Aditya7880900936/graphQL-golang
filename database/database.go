package database

import (
	"context"
	"log"
	"time"

	"github.com/Aditya7880900936/graphQL-golang/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = "mongodb://localhost:27017"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return &DB{client: client}
}

func (db *DB) GetJob(id string) *model.JobListing {
	var jobListing model.JobListing
	collection := db.client.Database("graphql-golang").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	err := collection.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		log.Fatal(err)
	}
	return &jobListing
}

func (db *DB) GetJobs() []*model.JobListing {
	var jobListings []*model.JobListing
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.client.Database("graphql-golang").Collection("jobs")
	cursor, err := collection.Find(ctx, map[string]string{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		var jobListing model.JobListing
		err := cursor.Decode(&jobListing)
		if err != nil {
			log.Fatal(err)
		}
		jobListings = append(jobListings, &jobListing)
	}
	return jobListings
}

func (db *DB) CreateJobLsiting(jobInfo model.CreateJobListingInput) *model.JobListing {
	var returnJobListing model.JobListing
	return &returnJobListing
}

func (db *DB) UpdateJobListing(id string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	var jobListing model.JobListing
	return &jobListing
}

func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	return &model.DeleteJobResponse{DeleteJobID: jobId}
}
