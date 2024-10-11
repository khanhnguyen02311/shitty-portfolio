// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package data

import (
	"database/sql"
)

type Blog struct {
	ID              int64
	Title           string
	Slug            string
	Summary         sql.NullString
	ThumbnailDir    sql.NullString
	ContentFilename string
	TimeCreated     sql.NullTime
}

type BlogTag struct {
	ID      int64
	IDBlog  int64
	TagName string
}