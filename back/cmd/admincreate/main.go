package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"temsys"
	"temsys/configuration"
	"temsys/hash"
	"temsys/mysql"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	adminUser := readUser()
	userRepo, hasher := createRepoAndHasher()
	if userRepo.ExistsWithName(adminUser.Name) {
		fmt.Println("El usuario ya existe")
		return
	}
	adminUser.Password = hasher.Hash(adminUser.Password)
	userRepo.Save(adminUser)
	fmt.Println("User created!")
}

func readUser() temsys.User {
	fmt.Print("Enter the admin name: ")
	name := readOrPanic()
	fmt.Print("Enter the password name: ")
	password := readOrPanic()
	return temsys.User{
		Name:     name,
		Password: password,
		Role:     temsys.AdminRole,
	}
}

func readOrPanic() string {
	data, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(data)
}

func createRepoAndHasher() (temsys.UserRepository, temsys.PasswordHasher) {
	config := configuration.Load()
	db := mysql.Connect(config)
	userRepo := mysql.NewUserRepo(db)
	hasher := hash.BcryptPasswordHasher{}
	return userRepo, hasher
}
