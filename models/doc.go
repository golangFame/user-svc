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

//pointers are not scanned by bun by default instead use relation() to scan for *p
//
//Relation(name) is just an alias for new coming table

/*



has-one
belongs-to
has-many
polymorphic-has-many
many-to-many

	Author	 Author `bun:"rel:belongs-to,join:author_id=id"`


# Arrays

https://bun.uptrace.dev/postgres/postgres-arrays.html

https://github.com/oiime/logrusbun




*/
