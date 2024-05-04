package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Abdu0811"
	dbname   = "n9"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()


	names:=[]string{"Createuser", "createfriendship"}
	for _,name:=range names{
		str:="../DB/"+name+".sql"
		fmt.Println(str)
		sqlfile,err:=os.ReadFile(str)
		if err!=nil{
			log.Fatal(err)
		}
		_,err=db.Exec(string(sqlfile))
		if err!=nil{
			log.Fatal(err)
		}
	}
	
	userID := 1
	friendID := 2

	err = requestFriendship(db, userID, friendID)
	if err != nil {
		log.Fatal("Error requesting friendship: ", err)
	} else {
		log.Println("Friendship request sent successfully")
	}

	err = acceptFriendship(db, friendID, userID)
	if err != nil {
		log.Fatal("Error accepting friendship:", err)
	} else {
		log.Println("Friendship accepted successfully")
	}

	err = blockUser(db, userID, friendID)
	if err != nil {
		log.Fatal("Error blocking user: ", err)
	} else {
		log.Println("User blocked successfully")
	}

}

func requestFriendship(db *sql.DB, userID, friendID int) error {
	name:="../DB/insertfriendship.sql"
	sqlfile,err:=os.ReadFile(string(name))
	if err!=nil{
		log.Fatal(err)
	}
	_, err = db.Exec(string(sqlfile), userID, friendID)
	if err != nil {
		return fmt.Errorf("could not insert friendship request: %v", err)
	}
	return nil
}

func acceptFriendship(db *sql.DB, userID, friendID int) error {
	name:="../DB/updatefriendship.sql"
	sqlfile,err:=os.ReadFile(string(name))
	if err!=nil{
		log.Fatal(err)
	}
	res, err := db.Exec(string(sqlfile), friendID, userID)
	if err != nil {
		return fmt.Errorf("could not accept friendship: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no pending request to accept")
	}

	return nil
}

func blockUser(db *sql.DB, userID, friendID int) error {
	query := `INSERT INTO friendships (user_id, friend_id, status)
              VALUES ($1, $2, 'blocked')
              ON CONFLICT (user_id, friend_id) 
              DO UPDATE SET status = EXCLUDED.status;`
    _, err := db.Exec(query, userID, friendID)
    if err != nil {
        return fmt.Errorf("could not block user: %v", err)
    }
    return nil
}
