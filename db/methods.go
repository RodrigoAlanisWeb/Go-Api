package db

import "fmt"

// User User
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone`
}

// Create Create A Contact
func Create(user User) {
	con := connect()

	defer con.Close()

	sql := fmt.Sprintf("INSERT INTO contacts VALUES(null,'%s','%s','%s')", user.Name, user.Email, user.Phone)

	query, err := con.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	defer query.Close()

}

// GetAll Get all the contacts
func GetAll() []User {

	con := connect()

	defer con.Close()

	results, err := con.Query("SELECT * FROM contacts")

	if err != nil {
		panic(err.Error())
	}

	var users []User

	for results.Next() {
		var user User

		err = results.Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users
}

// GetOne Get one contact
func GetOne(id int) []User {

	con := connect()

	defer con.Close()
	sql := fmt.Sprintf("SELECT * FROM contacts WHERE id = %d", id)
	results, err := con.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	var data []User

	for results.Next() {
		var user User

		err = results.Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
		if err != nil {
			panic(err.Error())
		}

		data = append(data, user)
	}

	return data
}

// Update update a contact
func Update(id int, body User) {

	con := connect()

	defer con.Close()

	sql := fmt.Sprintf("UPDATE contacts SET name = '%s', email = '%s', phone = '%s' WHERE id = %d",
		body.Name, body.Email, body.Phone, id)
	results, err := con.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	defer results.Close()
}

// Delete delete a contact
func Delete(id int) {
	con := connect()

	defer con.Close()

	sql := fmt.Sprintf("DELETE FROM contacts WHERE id = %d", id)
	results, err := con.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	defer results.Close()
}
