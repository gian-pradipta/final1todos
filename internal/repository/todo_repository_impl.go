package repository

import (
	"database/sql"
	"final_satu/internal/database"
	"final_satu/internal/entity"
	"log"
)

type todorepo struct {
	db *sql.DB
}

func New() Repository {
	var repo todorepo
	var err error
	repo.db, err = database.New()
	if err != nil {
		log.Fatal(err)
	}
	return &repo
}
func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func (t *todorepo) GetAll() (error, []entity.Todos) {
	var err error
	db := t.db

	query := "SELECT nama, tanggal_mulai, tanggal_selesai, deskripsi FROM todos"
	rows, err := db.Query(query)
	defer rows.Close()

	var todo entity.Todos
	var todos []entity.Todos = make([]entity.Todos, 0)
	for rows.Next() {
		err = rows.Scan(&todo.Nama, &todo.TanggalMulai, &todo.TanggalSelesai, &todo.Deskripsi)
		todos = append(todos, todo)
	}

	return err, todos
}
func (t *todorepo) GetOne(id int) (error, entity.Todos) {
	var err error
	db := t.db

	query := "SELECT nama, tanggal_mulai, tanggal_selesai, deskripsi FROM todos WHERE id = ?"
	rows, err := db.Query(query, id)
	defer rows.Close()

	var todo entity.Todos
	for rows.Next() {
		err = rows.Scan(&todo.Nama, &todo.TanggalMulai, &todo.TanggalSelesai, &todo.Deskripsi)
	}

	return err, todo
}
func (t *todorepo) Create(newTodo entity.Todos) error {
	var err error
	db := t.db

	query := "INSERT INTO todos (nama, tanggal_mulai, tanggal_selesai, deskripsi) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(query, newTodo.Nama, newTodo.TanggalMulai, newTodo.TanggalSelesai, newTodo.Deskripsi)

	return err

}
func (t *todorepo) Update(newTodo entity.Todos, id int) error {
	var err error
	db := t.db

	query := "UPDATE todos SET nama = ?, tanggal_mulai = ?, tanggal_selesai = ?, deskripsi = ?"
	_, err = db.Exec(query, newTodo.Nama, newTodo.TanggalMulai, newTodo.TanggalSelesai, newTodo.Deskripsi)

	return err
}
func (t *todorepo) Delete(id int) error {
	var err error
	db := t.db

	query := "DELETE FROM todos WHERE id = ?"
	_, err = db.Exec(query, id)

	return err
}
