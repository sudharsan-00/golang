package main

import (
	"fmt"
	"log"
)

type User struct {
	ID    int
	first string
}

type Mockdatastore struct {
	Users map[int]User
}

func (md Mockdatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("user %v not found", id)
	}
	return user,nil
}

func (md Mockdatastore) SaveUser(u User) error{
	_,ok := md.Users[u.ID]
    if ok {
		return fmt.Errorf("user %v aldready exist",u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

type Datastore interface{
	GetUser(id int) (User,error)
    SaveUser(u User) error
}

type Service struct{
	ds Datastore
}

func (s Service) GetUser(id int) (User, error){
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error{
	return s.ds.SaveUser(u)
}

func main(){
	db := Mockdatastore{
		Users: make(map[int]User),
	}

	srvs := Service{
		ds: db,
	}

	u1 := User{
		ID : 1,
		first: "James",
	}

	err := srvs.SaveUser(u1)
	if err != nil{
		log.Fatalf("error %s", err)
	}

	u1Returned, err := srvs.GetUser(u1.ID)
	if err != nil{
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)
}
