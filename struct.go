package main

import "time"

type Customer struct {
	// gorm.Model
	CustomerID  uint      `json:"customerId" gorm:"primary_key"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"createdAt"`
	CountryCode string    `json:"countryCode"`
}

type Plan struct {
	// gorm.Model
	PlanId       uint      `json:"planId" gorm:"primary_key"`
	Name         string    `json:"name"`
	Price        float32   `json:"price"`
	CreationDate time.Time `json:"creationDate"`
	Frequency    string    `json:"frequency"`
}

type Subscription struct {
	CustomerID uint      `json:"customerId" gorm:"primary_key"`
	PlanId     uint      `json:"planId" gorm:"primary_key"`
	CreatedAt  time.Time `json:"creationAt"`
	EndedAt    time.Time `json:"endedAt"`
	Status     bool      `json:"status"`
}
