package db

import (
	"context"
	"metriko/hardware"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CpuMetricStor interface {
	Insertcpu(ctx context.Context, cpu *hardware.CPU) (*hardware.CPU, error)
	GetCpus(ctx context.Context) ([]*hardware.CPU, error)
	GetCpuByID(ctx context.Context, id string) (*hardware.CPU, error)
}

type MongoCpuMetricStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoCpuMetricStore(client *mongo.Client,db_name string) *MongoIfaceMetricstor {
	return &MongoIfaceMetricstor{
		client: client,
		coll:   client.Database(db_name).Collection("cpu"),
	}
}
func (s *MongoIfaceMetricstor) Insertcpu(ctx context.Context, cpu *hardware.CPU) (*hardware.CPU, error) {
	resp, err := s.coll.InsertOne(ctx, cpu)
	if err != nil {
		return nil, err
	}
	cpu.ID = resp.InsertedID.(primitive.ObjectID)

	return cpu, nil
}
func (s *MongoIfaceMetricstor) GetCpus(ctx context.Context) ([]*hardware.CPU, error) {
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var cpus []*hardware.CPU
	if err := cur.All(ctx, &cpus); err != nil {
		return nil, err
	}

	return cpus, nil
}

func (s *MongoIfaceMetricstor) GetCpuByID(ctx context.Context, id string) (*hardware.CPU, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user hardware.CPU
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
