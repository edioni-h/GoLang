package main

import (
	"context"
	"fmt"
	"time"
)

type Member struct {
	ID        int64     `bun:",pk,autoincrement"`
	Name      string    `bun:",notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func main() {
	ctx := context.Background()

	// user := &Member{Name: "Edion"}
	// _, err = db.NewInsert().
	// 	Model(user).
	// 	Exec(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	//Deleting specific rows by deciding the rows ourselves
	// ids := []int64{2, 4, 6}
	// _, err = db.NewDelete().
	// 	Model(&Member{}).
	// 	Where("id IN (?)", bun.In(ids)).
	// 	Exec(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	//Adding multiple rows with the table formation bsed on the struct
	// users := []*Member{
	// 	{ID: 2, Name: "Hashi"},
	// 	{ID: 3, Name: "Hashani"},
	// 	{ID: 4, Name: "Leonit"},
	// 	{ID: 5, Name: "Roi"},
	// }
	// _, err = db.NewInsert().
	// 	Model(&users).
	// 	Exec(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	db, err := DBconnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	storedMembers := NewUserStore(db)

	// member, err := storedMembers.getByID(ctx, 1)
	// if err != nil {
	// 	fmt.Println("Getting error:", err)
	// }

	// fmt.Println("Member found:", member.Name)

	member, err := storedMembers.getAllMembers(ctx)
	if err != nil {
		fmt.Println("Getting error:", err)
	}
	for _, member := range member {
		fmt.Println("Member found:", member.Name)
	}

}
