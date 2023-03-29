package main

import (
	"errors"
	"fmt"
)

func chackUserName(username string) error {
	if username == "admin" {
		return errors.New("Username can not be admin")
	}
	return nil
}

func main() {
	userNames := []string{"user1", "admin", "user2", "user3"}
	for _, userName := range userNames {
		if err := chackUserName(userName); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Username is OK")
		}
	}
}
