// mongo.go
package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"user"
)

type Store struct {
	session *mgo.Session
	db      string
}

func New(mgoSession *mgo.Session, db string) *Store {
	return &Store{mgoSession, db}
}

func (s *Store) c() *mgo.Collection {
	return s.session.DB(s.db).C("user")
}

func (s *Store) GetUser(id string) (*user.User, error) {
	u := &user.User{}
	if err := s.c().FindId(id).One(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Store) CreateUser(u *user.User) error {
	if err := s.c().Insert(u); err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateUser(id, email, name string) (*user.User, error) {
	u, err := s.GetUser(id)
	if err != nil {
		return nil, err
	}

	u.Email = email
	u.Name = name

	if err := s.c().UpdateId(id, u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Store) ChangePassword(id, password string) error {
	if err := s.c().UpdateId(id, bson.M{"$set": bson.M{"password": password}}); err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteUser(id string) error {
	if err := s.c().RemoveId(id); err != nil {
		return err
	}

	return nil
}

func (s *Store) All() ([]user.User, error) {
	users := []user.User{}
	if err := s.c().Find(nil).All(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Store) GetUserByUsername(username string) (*user.User, error) {
	u := &user.User{}
	if err := s.c().Find(bson.M{"username": username}).One(u); err != nil {
		return nil, err
	}

	return u, nil
}
