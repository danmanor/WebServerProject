package main

import (
	"fmt"

	sw "github.com/danmanor/WebServerProject/go-client/go"
)

func addUserClient() {
	var name string
	var email string

	fmt.Println("Please Enter Your Name")
	fmt.Scanln(&name)
	fmt.Println("Please Enter Your email")
	fmt.Scanln(&email)

	cfg := sw.NewConfiguration()
	apiService := sw.NewAPIClient(cfg)

	newUser := sw.NewUser{Name: name, Email: email}

	user, _, err := apiService.DefaultApi.AddUser(nil, newUser)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Success! %s Was Added To The System And His Id Is: %s\n", user.Name, user.Id)
	}
}

func deleteUserClient() {
	var userId string
	fmt.Println("Please Enter the user's Id you want to delete")
	fmt.Scanln(&userId)

	cfg := sw.NewConfiguration()
	apiService := sw.NewAPIClient(cfg)

	user, _, err := apiService.DefaultApi.DeleteUser(nil, userId)
	fmt.Println(user)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Success! %s Was Deleted\n", user.Name)
	}
}

func getUsersClient() {
	cfg := sw.NewConfiguration()
	apiService := sw.NewAPIClient(cfg)

	users, _, err := apiService.DefaultApi.FindUsers(nil)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("ID", "NAME", "EMAIL")
		for _, v := range users {
			fmt.Println(v.Id, v.Name, v.Email)
		}
	}
}

func main() {
	for true {

		var proceed string

		fmt.Println("Do You Wish To Register A User (1), Delete A User (2), Get All Users (3) or Exit (4)?  (1/2/3/4)")
		fmt.Scanln(&proceed)

		switch proceed {
		case "4":
			break
		case "1":
			addUserClient()
		case "2":
			deleteUserClient()
		case "3":
			getUsersClient()
		default:
			fmt.Println("Wrong Key, Please Try Again")
		}

		if proceed == "4" {
			break
		}
	}
}
