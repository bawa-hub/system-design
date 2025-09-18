package domain

import (
	"errors"
	"regexp"
	"strings"
)

// Email value object
type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	if !isValidEmail(value) {
		return Email{}, errors.New("invalid email format")
	}
	return Email{value: strings.ToLower(value)}, nil
}

func (e Email) String() string {
	return e.value
}

func (e Email) Equals(other Email) bool {
	return e.value == other.value
}

func (e Email) Domain() string {
	parts := strings.Split(e.value, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched && len(email) > 5 && len(email) < 100
}

// Name value object
type Name struct {
	first string
	last  string
}

func NewName(first, last string) (Name, error) {
	first = strings.TrimSpace(first)
	last = strings.TrimSpace(last)
	
	if first == "" {
		return Name{}, errors.New("first name cannot be empty")
	}
	
	if last == "" {
		return Name{}, errors.New("last name cannot be empty")
	}
	
	if len(first) < 2 || len(first) > 50 {
		return Name{}, errors.New("first name must be between 2 and 50 characters")
	}
	
	if len(last) < 2 || len(last) > 50 {
		return Name{}, errors.New("last name must be between 2 and 50 characters")
	}
	
	return Name{first: first, last: last}, nil
}

func (n Name) First() string {
	return n.first
}

func (n Name) Last() string {
	return n.last
}

func (n Name) Full() string {
	return n.first + " " + n.last
}

func (n Name) Equals(other Name) bool {
	return n.first == other.first && n.last == other.last
}

// Money value object
type Money struct {
	amount   int64
	currency string
}

func NewMoney(amount int64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, errors.New("amount cannot be negative")
	}
	
	if !isValidCurrency(currency) {
		return Money{}, errors.New("invalid currency")
	}
	
	return Money{
		amount:   amount,
		currency: currency,
	}, nil
}

func (m Money) Amount() int64 {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, errors.New("cannot add different currencies")
	}
	
	return NewMoney(m.amount+other.amount, m.currency)
}

func (m Money) Multiply(factor float64) (Money, error) {
	if factor < 0 {
		return Money{}, errors.New("factor cannot be negative")
	}
	
	newAmount := int64(float64(m.amount) * factor)
	return NewMoney(newAmount, m.currency)
}

func (m Money) IsZero() bool {
	return m.amount == 0
}

func (m Money) IsPositive() bool {
	return m.amount > 0
}

func isValidCurrency(currency string) bool {
	validCurrencies := []string{"USD", "EUR", "GBP", "JPY", "CAD", "AUD"}
	for _, valid := range validCurrencies {
		if currency == valid {
			return true
		}
	}
	return false
}

// UserID value object
type UserID struct {
	value string
}

func NewUserID(value string) (UserID, error) {
	if value == "" {
		return UserID{}, errors.New("user ID cannot be empty")
	}
	return UserID{value: value}, nil
}

func (id UserID) String() string {
	return id.value
}

func (id UserID) Equals(other UserID) bool {
	return id.value == other.value
}

// ProductID value object
type ProductID struct {
	value string
}

func NewProductID(value string) (ProductID, error) {
	if value == "" {
		return ProductID{}, errors.New("product ID cannot be empty")
	}
	return ProductID{value: value}, nil
}

func (id ProductID) String() string {
	return id.value
}

func (id ProductID) Equals(other ProductID) bool {
	return id.value == other.value
}

// OrderID value object
type OrderID struct {
	value string
}

func NewOrderID(value string) (OrderID, error) {
	if value == "" {
		return OrderID{}, errors.New("order ID cannot be empty")
	}
	return OrderID{value: value}, nil
}

func (id OrderID) String() string {
	return id.value
}

func (id OrderID) Equals(other OrderID) bool {
	return id.value == other.value
}

// CustomerID value object
type CustomerID struct {
	value string
}

func NewCustomerID(value string) (CustomerID, error) {
	if value == "" {
		return CustomerID{}, errors.New("customer ID cannot be empty")
	}
	return CustomerID{value: value}, nil
}

func (id CustomerID) String() string {
	return id.value
}

func (id CustomerID) Equals(other CustomerID) bool {
	return id.value == other.value
}
