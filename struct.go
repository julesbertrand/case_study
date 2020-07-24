package main

import "time"

// Customer is the struct for a customer row in db
type Customer struct {
	// gorm.Model
	CustomerID  uint      `json:"customerId" gorm:"primary_key"`
	FirstName   string    `json:"firstName" binding:"required"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"createdAt"`
	CountryCode string    `json:"countryCode"`
	// Subscriptions []Subscription `json:"subscriptions" gorm:"foreignkey:CustomerID"`
}

// Plan is a plan that can be subscrpited to by a customer
type Plan struct {
	// gorm.Model
	PlanID    uint      `json:"planId" gorm:"primary_key"`
	PlanName  string    `json:"name" binding:"required"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	Frequency string    `json:"frequency"`
}

// Subscription is a link between a customer and a plan
type Subscription struct {
	SubscriptionID uint      `json:"subscriptionId" gorm:"primary_key"`
	CustomerID     uint      `json:"customerId"`
	PlanID         uint      `json:"planId"`
	CreatedAt      time.Time `json:"creatednAt"`
	EndedAt        time.Time `json:"endedAt"`
	ActiveStatus   bool      `json:"activeStatus"`
}
