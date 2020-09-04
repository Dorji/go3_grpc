package main

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	pb "github.com/vlasove/Lec13/userserver/proto/user"
)

// message User {
//     string id = 1;
//     string name = 2;
//     string email = 3;
//     string company = 4;
//     string password = 5;
// }

type User struct {
	ID       string `sql:"id"`
	Name     string `sql:"name"`
	Email    string `sql:"email"`
	Company  string `sql:"company"`
	Password string `sql:"password"`
}

//Интерфейс модели
type Repository interface {
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
}

//Представитель модели
type PostgresRepository struct {
	db *sqlx.DB
}

//Создает нового представителя модели
func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

//MarshallUser ...
func MarshallUser(user *pb.User) *User {
	return &User{
		ID:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Company:  user.Company,
		Password: user.Password,
	}
}

//UnmarshallUser ....
func UnmarshallUser(user *User) *pb.User {
	return &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Company:  user.Company,
		Password: user.Password,
	}
}

//MarshallUserColelction ...
func MarshallUserCollection(users []*pb.User) []*User {
	us := make([]*User, len(users))
	for _, val := range users {
		us = append(us, MarshallUser(val))
	}
	return us
}

//UnmarshallUserCollection ...
func UnmarshallUserCollection(users []*User) []*pb.User {
	us := make([]*pb.User, len(users))
	for _, val := range users {
		us = append(us, UnmarshallUser(val))
	}
	return us
}

// //Интерфейс модели
// type Repository interface {
// 	Get(ctx context.Context, id string) (*User, error) +
// 	GetAll(ctx context.Context) ([]*User, error) +
// 	GetByEmail(ctx context.Context, email string) (*User, error)
// 	Create(ctx context.Context, user *User) error +
// }

func (r *PostgresRepository) Create(ctx context.Context, user *User) error {
	user.ID = uuid.NewV4().String()
	log.Println(user)

	query := "insert into users (id, name, email, company, password) values($1, $2, $3, $4, $5)"
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Company, user.Password)
	return err
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err := r.db.GetContext(ctx, users, "select * from users"); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *PostgresRepository) Get(ctx context.Context, id string) (*User, error) {
	query := "select * from users where id=$1"
	var user *User
	if err := r.db.GetContext(ctx, &user, query, id); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := "select * from users where email=$1"
	var user *User
	if err := r.db.GetContext(ctx, &user, query, email); err != nil {
		return nil, err
	}
	return user, nil
}
