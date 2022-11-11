//https://bun.uptrace.dev/guide/models.html#mapping-tables-to-structs

// Package models /*
package models

//You can add/remove fields to/from an existing model by using extend tag option. The new model will inherit the table name and the alias from the original model.
//
//type UserWithCount struct {
//	User `bun:",extend"`
//
//	Name		string `bun:"-"` // remove this field
//	AvatarCount int				 // add a new field
//}

//err := db.NewSelect().
//Model(book).
//Column("book.id").
//Relation("Author", func (q *bun.SelectQuery) *bun.SelectQuery {
//	return q.Column("id")
//}).
//Where("id = 1").
//Scan(ctx)

//Use returning even in insert to get fields of default-id,created_at etc..

//Relation(name) is just an alias for new coming table
