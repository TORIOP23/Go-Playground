package main

import "time"

type User struct {
	ID string

	// SessionToken string // 1.
	// thay = Auth, we still can call method like usual
	// Auth // 2

	//Payment // 2

	Features *[]string
}

type PaidUser struct {
	User
	Auth
	Payment
}

type TrialUser struct {
	Auth
	User
	// doesn't make sense if Trial User has Payment
}

type Auth struct {
	SessionToken string
}
type Payment struct{}

func (u *Payment) GetLastPaymentDate(userId string) time.Time {
	return time.Now()
}

func (u *Payment) IsAccountPayed(userId string) bool {
	return true
}

func (u *Auth) IsLoggedIn(userId string) bool {
	return true
}

func (u *Auth) LoginUser(userId string, hashedPassword string) {

}

func (u *User) HasFeature(feature string) bool {
	return true
}
