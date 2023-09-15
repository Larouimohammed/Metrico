package db

import (
	"context"
	"metriko/hardware"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IfaceMetricStor interface {
	InsertIface(ctx context.Context, ifaces []hardware.Iface)  error
	GetIfaces(ctx context.Context) ([]*hardware.Iface, error)
	GetIfaceByID(ctx context.Context, id string) (*hardware.Iface, error)
}

type MongoIfaceMetricstor struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoIfaceMetricStore(client *mongo.Client,db_name string) *MongoIfaceMetricstor {
	return &MongoIfaceMetricstor{
		client: client,
		coll:   client.Database(db_name).Collection("iface"),
	}
}
func (s *MongoIfaceMetricstor) InsertIface(ctx context.Context, ifaces []hardware.Iface)  error {
	for _, iface := range ifaces {
	resp, err := s.coll.InsertOne(ctx, iface)

	if err != nil {
		return err
	}
	iface.ID = resp.InsertedID.(primitive.ObjectID)}

	return  nil
}
func (s *MongoIfaceMetricstor) GetIfaces(ctx context.Context) ([]*hardware.Iface, error) {
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var ifaces []*hardware.Iface
	if err := cur.All(ctx, &ifaces); err != nil {
		return nil, err
	}

	return ifaces, nil
}

func (s *MongoIfaceMetricstor) GetIfaceByID(ctx context.Context, id string) (*hardware.Iface, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var iface hardware.Iface
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&iface); err != nil {
		return nil, err
	}
	return &iface, nil
}
