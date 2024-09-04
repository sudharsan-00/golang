package main

import "testing"

func TestGetUser(t *testing.T) {
	md := Mockdatastore{
	Users: map[int]User{
		2:{ID: 2, first: "Jenny"},
	},
}
s := &Service{
	ds : md,
}

u, err := s.GetUser(2)
if err != nil{
	t.Errorf("got error: %v",err)
}

if u.first != "Jenny"{
	t.Errorf("got: %s, want:%s", u.first,"Jenny")
}
}