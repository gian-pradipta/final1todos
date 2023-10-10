package repository

import "final_satu/internal/entity"

type Repository interface {
	/*
	* Sebelum membaca ini, silahkan dicek dulu folder entity. di sana ada struct todo
	* yang merepresentasikan sebuah record dari tabel todo di database
	 */

	/*
	* Fungsi GetAll mengembalikan array dari struct todos
	* yang berasal dari package entity di folder entity
	* Struct todos adalah representasi sebuah record dari tabel todo di database.
	* Memanggil fungsi ini akan mengembalikan semua record dari tabel todo
	* Fungsi juga mereturn error, jangan lupa di - handle
	 */
	GetAll() (error, []entity.Todos)

	/*
	* Fungsi GetOne(id int) menerima parameter berupa integer.
	* Fungsi ini akan mengembalikan satu entity struct todos
	* Fungsi juga mereturn error, jangan lupa di - handle
	 */
	GetOne(id int) (error, entity.Todos)

	/*
	* Fungsi ini menerima parameter berupa entity struct todo.
	* Fungsi ini akan membuat record baru di database sesuai dengan entity todo yang dipass ke
	* fungsi ini
	* Fungsi juga mereturn error, jangan lupa di - handle
	 */
	Create(newTodo entity.Todos) error

	/*
	* Fungsi Update akan menerima yang pertama adalah entity todo record yang baru
	* dan kedua adalah id berupa integer yang menandakan id dari record yang
	* ingin diupdate.
	* Fungsi juga mereturn error, jangan lupa di - handle
	 */
	Update(newTodo entity.Todos, id int) error

	/*
	* Fungsi delete menerima parameter id berupa integer yang akan menghapus
	* record di database sesuai dengan idnya
	* Fungsi juga mereturn error, jangan lupa di - handle
	 */
	Delete(id int) error
}
