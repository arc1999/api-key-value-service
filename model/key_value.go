package model

type KeyValue struct {
	Key string `json:"key" validate:"required"`
	Value string `json:"value"`
}
