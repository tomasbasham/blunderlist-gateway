package entity

import "time"

// Task is a type representing a single task.
type Task struct {
	ID        uint       `jsonapi:"primary,tasks"`
	Title     string     `jsonapi:"attr,title"`
	Completed bool       `jsonapi:"attr,completed"`
	CreatedAt time.Time  `jsonapi:"attr,created_at,iso8601"`
	Comments  []*Comment `jsonapi:"relation,comments"`
}
