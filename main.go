package main

import (
	"fmt"
	"log"
	"os"

	"github.com/epointpayment/evault/app/controllers"
	"github.com/epointpayment/evault/app/database"
	"github.com/epointpayment/evault/app/router"
	"github.com/epointpayment/evault/app/services"
	User "github.com/epointpayment/evault/app/services/user"

	"github.com/namsral/flag"
)

var (
	dsn        string
	createUser bool
	username   string
	password   string
	programID  int
)

func init() {
	flag.StringVar(&dsn, "dsn", "", "path to sample data i.e. user:pass@/example")

	flag.BoolVar(&createUser, "createUser", false, "create a new user")
	flag.StringVar(&username, "username", "", "username for new user")
	flag.StringVar(&password, "password", "", "password for new user")
	flag.IntVar(&programID, "programID", 0, "")

	flag.Parse()
}

func main() {
	var err error

	// Create new connection handler for database
	db := database.NewDatabase("mysql", dsn)
	err = db.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Setup services
	err = services.New(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a new user
	if createUser {
		us := User.New()
		user, err := us.Create(username, password, programID)
		if err != nil {
			fmt.Println("Unable to create user: " + err.Error())
			os.Exit(1)
		}

		fmt.Println(fmt.Sprintf("User has been created:  %s ", user.Username))
		os.Exit(0)
	}

	// Setup router and run
	c := controllers.NewControllers()
	r := router.NewRouter(c)
	err = r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
