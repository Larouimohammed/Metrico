package hardware

import "go.mongodb.org/mongo-driver/bson/primitive"

type CPU struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Product string             `json:"product"`
	Vendor  string             `json:"vendor"`
	Width   string             `json:"width"`
	Version string             `json:"version"`
}
