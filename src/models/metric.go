package models

import (
	"time"
)

type Metric struct {
	Name        string    `json:"name"`
	Value       float64   `json:"value"`
	Timestamp   time.Time `json:"timestamp"`
}