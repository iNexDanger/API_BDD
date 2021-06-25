package main

func createPerson(person Person) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO Personas (ID, Name, DNI) VALUES (?, ?, ?)", person.ID, person.Name, person.DNI)
	return err
}
func deletePerson(id int64) error {

	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM Personas WHERE id = ?", id)
	return err
}
func updatePersons(person Person) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE Personas SET ID = ?, Name = ?, DNI = ? WHERE id = ?", person.ID, person.Name, person.DNI)
	return err
}

func getPersons() ([]Person, error) {
	//Declare an array because if there's error, we return it empty
	persons := []Person{}
	bd, err := getDB()
	if err != nil {
		return persons, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT ID, Name, DNI FROM Personas")
	if err != nil {
		return persons, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var person Person
		err = rows.Scan(&person.ID, &person.Name, &person.DNI)
		if err != nil {
			return persons, err
		}
		// and append it to the array
		persons = append(persons, person)
	}
	return persons, nil
}
func getPersonById(id int64) (Person, error) {
	var person Person
	bd, err := getDB()
	if err != nil {
		return person, err
	}
	row := bd.QueryRow("SELECT ID, Name, DNI FROM Personas WHERE id = ?", id)
	err = row.Scan(&person.ID, &person.Name, &person.DNI)
	if err != nil {
		return person, err
	}
	// Success!
	return person, nil
}
