package models

// Varun: Struct to define the subscription plan details
type SubscriptionPlan struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Varun: Struct to define user subscriptions
type Subscription struct {
	ID              int    `json:"id"`
	SubscriptionID  string `json:"subscription_id"`
	CustomerID      string `json:"customer_id"`
	PlanID          string `json:"plan_id"`
	Status          string `json:"status"`
	NextBillingDate string `json:"next_billing_date"`
	PaymentMethodID string `json:"payment_method_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// Varun: Struct for representing a payment
type Payment struct {
	ID         int     `json:"id"`
	UserID     string  `json:"user_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
	Date       string  `json:"date"`
	PaymentID  string  `json:"payment_id"`
	InvoiceURL string  `json:"invoice_url"`
}

// Varun: Struct for representing an invoice
type Invoice struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id"`
	PaymentID string `json:"payment_id"`
	Content   string `json:"content"`
}

// Varun: Struct for representing a payment method
type PaymentMethod struct {
	ID              int    `json:"id"`
	UserID          string `json:"user_id"`
	MethodType      string `json:"method_type"`
	CardLastFour    string `json:"card_last_four"`
	ExpiryMonth     string `json:"expiry_month"`
	ExpiryYear      string `json:"expiry_year"`
	PaymentMethodID string `json:"payment_method_id"`
}

// Varun: Struct to handle change plan request
type ChangePlanRequest struct {
	SubscriptionID string `json:"subscription_id"`
	NewPlanID      string `json:"new_plan_id"`
}

// Varun: Struct to handle change payment method request
type ChangePaymentRequest struct {
	SubscriptionID  string `json:"subscription_id"`
	PaymentMethodID string `json:"payment_method_id"`
}
