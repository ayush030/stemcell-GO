package models

import "github.com/google/uuid"

type Resource struct {
	Id      string `json:"id" gorm:"primary_key"`
	Payload string `json:"payload"`
}

func NewResource(payload string) *Resource {
	return &Resource{
		Id:      uuid.New().String(),
		Payload: payload,
	}
}

type CreateUpdateResource struct {
	Payload string `json:"payload"`
}
