package userservice

import (
	"bytes"
	"chat-chat-go/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUserData(userId string) (user *models.UserDto, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/getUserData/%s", os.Getenv("USER_SERVICE_URL"), userId), http.NoBody)
	if err != nil {
		log.Println("UserService: user not found")
		return user, err
	}

	req.Header.Add("X-API-Key", os.Getenv("GET_USER_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("UserService: failed to get user")
		return user, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("UserService: failed to parse user")
		return user, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("UserService: failed to parse user")
		return user, err
	}

	return user, nil
}

func (u *UserService) GetUsersData(ids []string) (users []models.UserDto, err error) {
	reqbody, err := json.Marshal(ids)
	if err != nil {
		log.Println("UserService: can not make request to get users")
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/getUsersData/", os.Getenv("USER_SERVICE_URL")), bytes.NewReader(reqbody))
	if err != nil {
		log.Println("UserService: user not found")
		return nil, err
	}

	req.Header.Add("X-API-Key", os.Getenv("GET_USER_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("UserService: failed to get users")
		return users, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("UserService: failed to parse users")
		return users, err
	}

	err = json.Unmarshal(body, &users)
	if err != nil {
		log.Println("UserService: failed to parse users")
		return users, err
	}

	return users, nil
}
