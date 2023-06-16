package models

import (
	"database/sql"
	"fmt"
	"gamex/config"
	"gamex/entities"
)

type UserModel struct {
	conn *sql.DB
}

type CommentModel struct {
	conn *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		conn: conn,
	}
}

func NewCommentModel() *CommentModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &CommentModel{
		conn: conn,
	}
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {

	row, err := u.conn.Query("select id, full_name, email, username, password, role from users where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&user.Id, &user.Full_name, &user.Email, &user.Username, &user.Password, &user.Role)
	}

	return nil
}

func (u UserModel) Create(user entities.User) (int64, error) {

	result, err := u.conn.Exec("insert into users (full_name, email, username, password) values(?,?,?,?)",
		user.Full_name, user.Email, user.Username, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil
}

func (u UserModel) CreateUser(user entities.User) (int64, error) {

	result, err := u.conn.Exec("insert into users (full_name, email, username, password, role) values(?,?,?,?,?)",
		user.Full_name, user.Email, user.Username, user.Password, user.Role)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil
}

func (u CommentModel) Createcomment(comment entities.Comment) (int64, error) {
	//var euser entities.User
	// user, err := user.Current()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// username := user.Username

	// fmt.Printf("Username is: %s\n", username)

	result, err := u.conn.Exec("insert into comment (title) values(?)",
		comment.Title)

	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("bad id")
	}

	return lastInsertId, nil

}

func (p *CommentModel) FindAllComment() ([]entities.Comment, error) {

	rows, err := p.conn.Query("select title from comment order by comment_id desc")
	if err != nil {
		return []entities.Comment{}, err
	}
	defer rows.Close()

	var dataComment []entities.Comment
	for rows.Next() {
		//var user entities.User
		var comment entities.Comment
		rows.Scan(&comment.Id,
			&comment.Title,
		)
		dataComment = append(dataComment, comment)

	}

	return dataComment, nil

}

//ADMIN PANEL

func (p *UserModel) FindAll() ([]entities.User, error) {

	rows, err := p.conn.Query("select * from users")
	if err != nil {
		return []entities.User{}, err
	}
	defer rows.Close()

	var dataUser []entities.User
	for rows.Next() {
		var user entities.User
		rows.Scan(&user.Id,
			&user.Full_name,
			&user.Email,
			&user.Username,
			&user.Password,
			&user.Role,
		)

		dataUser = append(dataUser, user)
	}

	return dataUser, nil

}

func (p *UserModel) Find(id int64, user *entities.User) error {

	return p.conn.QueryRow("select id, full_name, email,username, role from users where id = ?", id).Scan(
		&user.Id,
		&user.Full_name,
		&user.Email,
		&user.Username,
		&user.Role)
}

func (p *UserModel) Update(user entities.User) error {

	_, err := p.conn.Exec(
		"update users set full_name = ?, email = ?, username = ?, role = ? where id = ?",
		user.Full_name, user.Email, user.Username, user.Role, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *UserModel) Delete(id int64) {
	p.conn.Exec("delete from users where id = ?", id)
}

// func (p *UserModel) currentUser(id int64) {
// 	if()
// 	p.conn.Exec("select id from users where id = ?", id)
// }
