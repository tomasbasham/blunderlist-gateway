package entity

import "time"

// Comment is a type representing a single comment associated with a task.
type Comment struct {
	ID        uint      `jsonapi:"primary,comments"`
	TaskID    uint      `jsonapi:"attr,task_id"`
	Text      string    `jsonapi:"attr,text"`
	Author    string    `jsonapi:"attr,author"`
	CreatedAt time.Time `jsonapi:"attr,created_at,iso8601"`
}
