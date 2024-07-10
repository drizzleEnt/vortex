package model

import "time"

type Client struct {
	ID          int64
	ClientName  string  `json:"client_name"`
	Version     int     `json:"version"`
	Image       string  `json:"image"`
	CPU         string  `json:"cpu"`
	Memory      string  `json:"memory"`
	Priority    float64 `json:"priority"`
	NeedRestart bool    `json:"need_restart"`
	SpawnedAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
