package databases
import "fmt"
func CreateAllTable() {
	err,session,keySpaceMeta := ConnectToDatabase()
	defer session.Close()
	if err != nil {
		fmt.Println("Could not connect")
	}
	if _,exits := keySpaceMeta.Tables["users"];exits != true {
		err:= session.Query("CREATE TABLE users (" +
		"id text, username text, password text, email text, birthdate text, phonenumber text,tokenuser text," +
		"PRIMARY KEY (id))").Exec()
		if err != nil {
			fmt.Println("Could not create table users")
		}
	}
}