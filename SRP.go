/*
	Single Responsibility Principle says that a class should have only one responsibility.
*/

package main

import (
	"fmt"
	"log"
	"strings"
)

/*
	-----------------------Bad Example--------------------------------
	The below UserManagement struct and the CreateUser method are -
	1. Creating a user
	2. Validating the email
	3. Saving the user to the database
	4. Sending a welcome email

	They are doing too many things. It violates the Single Responsibility Principle.

	Problems with this:
	1. If email validation rules change → modify this code
	2. If database changes → modify this code
	3. If email service changes → modify this code
	4. Too many reasons to change!
*/

type UserManagement struct{}

func (um *UserManagement) CreateUser(name, email string) {
	fmt.Printf("Creating User: %s\n", name)

	if !strings.Contains(email, "@") {
		log.Println("Invalid Email - ", email)
		return
	}

	fmt.Println("Saving User to the database")

	fmt.Printf("Sending Welcome Email to %s\n", email)
}

/*
	-----------------------Good Example--------------------------------
	We will refactore the above code such that each responsibility is in a separate struct.
	1. UserValidator - Validates the email
	2. SaveUser - Saves the user to the database
	3. EmailService - Sends the welcome email

	Now, we can use these structs to create a UserManagementService that will create a user.

	In this way, we have separated the responsibilities of the UserManagement struct into three separate structs.

	The Simple Rules
	1. One job per struct/function
	2. One reason to change
	3. If you can describe what your code does with "AND", it's probably doing too much
*/

type User struct {
	Name  string
	Email string
}

type UserValidator struct{}

func (uv *UserValidator) ValidateEmail(email string) bool {
	return strings.Contains(email, "@") && len(email) > 5
}

type SaveUser struct{}

func (su *SaveUser) SaveUser(user User) error {
	fmt.Println("Saving User to the database")
	return nil
}

type EmailService struct{}

func (es *EmailService) SendWelcomeEmail(email string) error {
	fmt.Printf("Sending Welcome Email to %s\n", email)
	return nil
}

type UserManagementService struct {
	validator    *UserValidator
	saveUser     *SaveUser
	emailService *EmailService
}

func NewUserManagementService() *UserManagementService {
	return &UserManagementService{
		validator:    &UserValidator{},
		saveUser:     &SaveUser{},
		emailService: &EmailService{},
	}
}

func (ums *UserManagementService) CreateUser(name, email string) error {
	if !ums.validator.ValidateEmail(email) {
		return fmt.Errorf("invalid email")
	}

	user := User{Name: name, Email: email}

	if err := ums.saveUser.SaveUser(user); err != nil {
		return err
	}

	return ums.emailService.SendWelcomeEmail(email)
}

func demo() {
	fmt.Println("Single Responsibility Principle")

	userManagementService := NewUserManagementService()
	err := userManagementService.CreateUser("Abhishek", "example@gmail.com")
	if err != nil {
		log.Println("Error creating user - ", err)
	}

	fmt.Println("User created successfully")
}
