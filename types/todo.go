package types

import "time"

type Todo struct {
	ID        string    `bson:"_id" json:"id"`
	UserID    string    `bson:"userID" json:"userID"`
	Completed bool      `bson:"completed" json:"completed"`
	Content   string    `bson:"body" json:"content"`
	CreateAt  time.Time `bson:"createAt" json:"createAt"`
	FinshAt   time.Time `bson:"finshAt" json:"finshAt"`
}
