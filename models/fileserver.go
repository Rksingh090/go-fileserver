package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ImageModel struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FileName         string             `json:"file_name,omitempty" bson:"file_name,omitempty"`
	OriginalFileName string             `json:"original_file_name,omitempty" bson:"original_file_name,omitempty"`
	FilePath         string             `json:"file_path,omitempty" bson:"file_path,omitempty"`
	UsedFor          string             `json:"used_for,omitempty" bson:"used_for,omitempty"`
	Url              string             `json:"url,omitempty" bson:"url,omitempty"`
	CreatedAt        primitive.DateTime `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt        primitive.DateTime `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
