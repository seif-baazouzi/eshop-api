package models

import "time"

type Comment struct {
	CommentID    uint      `json:"commentID"`
	CommentValue string    `json:"commentValue"`
	CommentDate  time.Time `json:"commentDate"`
	Username     string    `json:"username"`
}
