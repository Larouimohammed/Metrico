package hardware

import "go.mongodb.org/mongo-driver/bson/primitive"

type Route struct {
	Dst string `json:"dst"`
	Src string `json:"src"`
	Gw  string `json:"gw"`
}
type Iface struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `json:"name"`
	HardwareAddr string             `json:"Hardwareaddr"`
	Type         string             `json:"type"`
	Flags        string             `json:"flags"`
	Routes       []Route            `json:"routes"`
}
