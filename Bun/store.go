package main

import (
	"context"

	"github.com/uptrace/bun"
)

type memberStore struct {
	db *bun.DB
}

func newUserStore(db *bun.DB) *memberStore {
	return &memberStore{db: db}
}

//With this code, you can only call one member at a time
// func (s *memberStore) getByID(ctx context.Context, _ int64) (*Member, error) {
// 	user := new(Member)
// 	err := s.db.NewSelect().
// 		Model(user).
// 		Where("id = ?", 1).
// 		Limit(1).
// 		Scan(ctx)
// 	return user, err
// }

// With this code, you can call all members at once, and with a range clause on the main file, we get a neat row of the members
func (s *memberStore) getAllMembers(ctx context.Context) ([]*Member, error) {
	var users []*Member
	err := s.db.NewSelect().
		Model(&users).
		Scan(ctx)
	return users, err
}
