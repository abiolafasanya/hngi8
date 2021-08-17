package config

// Connect database
db, err := sql.Open("mysql", "root@tcp(localhost)/golalng")
if err != nil {
	panic(err)
}

// defer db.Close()
// fmt.Println("Database connected successfullly")

// // Select from database
// results, err := db.Query("SELECT * FROM users");

// if err != nil {
// 	panic(err.Error())
// }
// for results.Next() {
//   var user users

//   err = results.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email)
//   if err != nil {
// 	  panic(err.Error())
//   }
//   fmt.Printf("Firstname: %s\n Lastname: %s\n Email: %s\n", user.Firstname, user.Lastname, user.Email)


//   fmt.Println("Data Selected Successfullly")
}