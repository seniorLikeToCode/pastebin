package service

import (
	"database/sql"
	"fmt"
)

type Data struct {
	ID      int    `json:"id"`
	Uid     string `json:"uid"`
	Content string `json:"content"`
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Create(d Data) error {
	_, err := s.db.Exec("INSERT INTO data (uid, content) VALUES ($1, $2)", d.Uid, d.Content)
	if err != nil {
		return fmt.Errorf("failed to create data: %w", err)
	}
	return nil
}

func (s *Store) GetDatabyUid(uid string) (*Data, error) {
	row := s.db.QueryRow("SELECT id, uid, content FROM data WHERE uid = $1", uid)

	var d Data
	if err := row.Scan(&d.ID, &d.Uid, &d.Content); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("data not found")
		}
		return nil, fmt.Errorf("failed to get data by uid: %w", err)
	}

	return &d, nil
}
