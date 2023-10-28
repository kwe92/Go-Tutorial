package main

import "fmt"

type Department int

const (
	Management Department = iota
	Sales
	Development
	HelpDesk
)

func getTitle(title int) string {

	var titleString string

	switch {
	case title == 0:
		titleString = "Management"
	case title == 1:
		titleString = "Sales"
	case title == 2:
		titleString = "Development"
	case title == 3:
		titleString = "HelpDesk"
	default:
		titleString = "Unknown"

	}

	return titleString

}

type Employee struct {
	Fname      string
	Lname      string
	Age        int
	Salary     float64
	Department int
}

func (e *Employee) Description() string {
	return fmt.Sprintf("Department: %s Name: %s Salary: %0.2f", getTitle(e.Department), e.Fname+" "+e.Lname, e.Salary)
}

type Manager struct {
	// the user defined type Employee embedded in the user defined type Manager
	Employee
	Reports []Employee
}

// func (m *Manager) Description() string {
// 	return "I WILL BE CALLED FIRST"
// }

func main() {
	var emp00 = Employee{
		Fname:      "Shikamaru",
		Lname:      "Nara",
		Age:        14,
		Salary:     110_000,
		Department: int(Development),
	}

	var emp01 = Employee{
		Fname:      "Baki",
		Lname:      "Hanma",
		Age:        16,
		Salary:     45_000,
		Department: int(HelpDesk),
	}

	empolyees := []Employee{emp00, emp01}

	var mgr00 = Manager{
		Employee: Employee{
			Fname:      "Kweayon",
			Lname:      "Clark",
			Age:        31,
			Salary:     275_000,
			Department: int(Management),
		},
		Reports: empolyees,
	}

	fmt.Println(mgr00.Description())

	for _, emp := range mgr00.Reports {
		fmt.Println(emp.Description())
	}

}

// Object Composition Over CLass Inheritance

// Embedding For Composition

//   - GO does not support class inheritance
//   - embedding is a way of composition to encourage code reuse and keep code DRY
//   - embedded fields are promoted to the containing struct

// Embedded Field

//   - embedded types have their fields and methods promoted to the containing struct
//   - any type can be embedded in a field, including built-in types
//   - embedded fields are unnamed and take on the name of the type
//   - you must instantiate the embedded type and then the containing type will
//     have all fields and methods

// Overriding Embedded Field Methods and Field Names

//   - methods can not be overridden
//   - if the containing struct and embedded field have the same method name
//     the containing field method will take precedence and the embedded field must be
//     explicitly called with dot notation
//  - the same applies if the enbedded field and containing field share the same
//    name for a field

// Embedding vs Inheritance

//   - embedding is still composition by instantiation and no fields or methods are overridden
