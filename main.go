// abhinav OBLLnnzk3BGqwlgl
//"mongodb+srv://abhinav:OBLLnnzk3BGqwlgl@cluster0.snfeuii.mongodb.net/?retryWrites=true&w=majority"

package main

import (
	"context"
	"demo/database"
	"demo/handlers"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	dsn="mongodb+srv://abhinav:abhinav@cluster0.snfeuii.mongodb.net/?retryWrites=true&w=majority"
	PORT   string
)

func main() {
	// Create a new client and connect to the server
	client, err := database.GetConnection("mongodb+srv://abhinav:abhinav@cluster0.snfeuii.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}

	//Disconnect function to disconnect connection after the work is done
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	
	err=database.Ping(client,"demo")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	}

	flag.StringVar(&PORT, "port", "50080", "--port=50080 or -port=50080")
	flag.Parse()

	router := mux.NewRouter()

	srv := http.Server{
		Addr:         ":" + PORT,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}
	empDB:=new(database.Employee)
	empDB.Client=client
	empDB.Dbname="employeedb"
	empDB.Collection="employees"

	eh:=new(handlers.EmployeeHandler)
	eh.DB=empDB

	router.HandleFunc("/employee/add",eh.Add)
	router.HandleFunc("/employee/delete/{id}",eh.Delete)



	srv.ListenAndServe()


	//Create a employee object
	/* e1:=new (models.Employee)
	e1.Email="ayush.ranjan275@gmail.com"
	e1.Mobile="12312312"
	e1.Name="Abhinav"
	e1.Status="Active"
	e1.LastModified=time.Now().Unix()

	empDB:=new(database.Employee)
	empDB.Client=client
	empDB.Dbname="employeedb"
	empDB.Collection="employees"

	result,err:=empDB.Insert(context.Background(),e1)
	fmt.Println(result,err) */



}


