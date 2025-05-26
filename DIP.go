/*
	Dependency Inversion Principle says that the high level module should not depend on the low level module.
	Both should depend on abstractions. In simple words, don't depend on specific terms. Depend on general contract instead
*/

package main

/*
	-----------------------Bad Example--------------------------------
	In the below example, we have a UserService struct that has a MySQLDB struct.
	We have a SaveUser method that saves the user to the MySQL database.

	But, this violates the Dependency Inversion Principle because the UserService struct is dependent on the MySQLDB struct.

	Now, if we want to use a different database, we will have to modify the UserService struct.
*/

type MySQLDB struct {
	SaveToMySQL func(user string)
}

type UserService struct {
	mysqlDB *MySQLDB // stuck with MySQLDB
}

func (us *UserService) SaveUser(user string) {
	us.mysqlDB.SaveToMySQL(user) // can't use other DBs
}

/*
	-----------------------Good Example--------------------------------
	In the below example, we have a DB interface that has a Save method.
	We have a UserServiceGood struct that has a DB interface.

	Now, we can use any DB that implements the DB interface since we are not dependent on any particular DB.
*/

type DB interface {
	Save(data string) error
}

type UserServiceGood struct {
	db DB // Not dependent on any particular DB
}

func (us *UserServiceGood) SaveUser(user string) {
	us.db.Save(user) // can use any DB
}