// Package aggregates holds aggregates that combines many entities into a full object

package aggregate

import (
	"errors"

	"github.com/google/uuid"

	"github.com/santos/ddd-go/entity"
	"github.com/santos/ddd-go/valueobject"
)

var (
	// ErrInvalidPerson is returned when the person is not valid in the NewCustomer factory
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

// Customer is a aggregate that combines all entities needed to represent a customer
type Customer struct {
	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	person *entity.Person
	// a customer can hold many products
	products []*entity.Item
	// a customer can perform many transactions
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new Customer aggregate
// It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	// Validate that the Name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	customer := Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}

	return customer, nil
}
