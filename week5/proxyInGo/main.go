package main

import "fmt"

type Connected interface { // interface
	Connect()
}

type Server struct{}

func (s *Server) Connect() {
	fmt.Println("connected server") // this will be printed if user licence count is under the limit
}

type User struct {
	LicenseID int // users have licence IDs for each difference company that uses this server system
}

type ServerProxy struct { // proxy struct
	server Server
	user   *User
}

func (s ServerProxy) Connect(countA int, countB int, limitA int, limitB int) (a int, b int) {
	// check for new user is licence and increase the active users in that licenceID
	if s.user.LicenseID == 1 {
		countA++
		if limitA >= countA { // if active users still lower than limit, let user connect the server
			s.server.Connect()
		} else { // if else don't let them join
			fmt.Println(" too much user in licence a")
		}
	}
	if s.user.LicenseID == 2 {
		countB++
		if limitB >= countB {
			s.server.Connect() // if active users still lower than limit, let user connect the server
		} else { // if else don't let them join
			fmt.Println(" too much user in licence b")
		}
	}
	a = countA
	b = countB
	return
}

func NewServerProxy(user *User) *ServerProxy {
	return &ServerProxy{Server{}, user}
}
func main() {
	// let's say we have two different company which are company a and company b
	// company a registered with maximum 3 user at the same time and company b registered with maximum 2 user at the same time
	// we will need to store active user count in variables

	//licenceID : 1 is the licence of company a and licenceID : 2 is the licence of company b
	countA := 0 // active a users
	countB := 0 // active b users
	limitA := 3 // limit of active a users
	limitB := 2 // limit of active b users

	server := NewServerProxy(&User{1}) // someone connecting with licence of company a
	countA, countB = server.Connect(countA, countB, limitA, limitB)

	server = NewServerProxy(&User{1}) // someone connecting with licence of company a
	countA, countB = server.Connect(countA, countB, limitA, limitB)

	server = NewServerProxy(&User{1}) // someone connecting with licence of company a
	countA, countB = server.Connect(countA, countB, limitA, limitB)

	// at this moment active user of company a is 3 which is limit for them so people can't connect with licenceID 1 anymore

	server = NewServerProxy(&User{1}) // someone connecting with licence of company a
	countA, countB = server.Connect(countA, countB, limitA, limitB)

	// output :
	// connected server
	// connected server
	// connected server
	// too much user in licence a

}
