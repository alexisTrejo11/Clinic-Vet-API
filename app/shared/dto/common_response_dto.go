package commondto

// CustomerDetails represents basic customer information
// swagger:model CustomerDetails
type CustomerDetails struct {
	// The unique identifier of the customer
	// Required: true
	// Example: 8
	ID uint `json:"id"`

	// The first name of the customer
	// Required: true
	// Max length: 50
	// Example: John
	FirstName string `json:"first_name"`

	// The last name of the customer
	// Required: true
	// Max length: 50
	// Example: Doe
	LastName string `json:"last_name"`
}

// EmployeeDetails represents basic employee (veterinarian) information
// swagger:model EmployeeDetails
type EmployeeDetails struct {
	// The unique identifier of the employee
	// Required: true
	// Example: 3
	ID uint `json:"id"`

	// The first name of the employee
	// Required: true
	// Max length: 50
	// Example: Jane
	FirstName string `json:"first_name"`

	// The specialty of the veterinarian
	// Required: true
	// Max length: 100
	// Example: Cardiology
	Specialty string `json:"specialty"`

	// The last name of the employee
	// Required: true
	// Max length: 50
	// Example: Smith
	LastName string `json:"last_name"`
}

// PetDetails represents basic pet information
// swagger:model PetDetails
type PetDetails struct {
	// The unique identifier of the pet
	// Required: true
	// Example: 5
	ID uint `json:"id"`

	// The name of the pet
	// Required: true
	// Max length: 100
	// Example: Max
	Name string `json:"name"`

	// The species of the pet
	// Required: true
	// Max length: 50
	// Example: Dog
	Species string `json:"species"`

	// The breed of the pet
	// Required: true
	// Max length: 50
	// Example: Golden Retriever
	Breed string `json:"breed"`

	// The age of the pet in years
	// Required: true
	// Minimum: 0
	// Maximum: 30
	// Example: 5
	Age int `json:"age"`

	// The weight of the pet in kilograms
	// Required: true
	// Minimum: 0.1
	// Maximum: 1000.0
	// Example: 25.5
	Weight float64 `json:"weight"`
}
