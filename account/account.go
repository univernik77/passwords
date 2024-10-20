package account

import (
	
	"errors"
	"math/rand/v2"
	"net/url"
	"time"
	"github.com/fatih/color"
)

var letterRunes = []rune("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")

type Account struct {
    Login string `json:"login"`
    Password string `json:"password"`
    Url string`json:"url"`
}

type AccountwithTimeStamp struct {
     CreatedAt time.Time `json:"createdAt"`
     UpdatedAt time.Time`json:"updatedAt"`
     Account
}

func (acc *Account) Output() {
    color.Yellow(acc.Login)
    color.Yellow(acc.Password)
    color.Yellow(acc.Url)

}

func (acc *Account) generatePassword(n int) {
    password := make([]rune, n)
    for i  := range password {
        password[i] = letterRunes[rand.IntN(len(letterRunes))]
    }
    acc.Password = string(password)
}

func NewAccountwithTimeStamp(login, password, urlString string) (*AccountwithTimeStamp, error) {
    if login == ""{
        return nil, errors.New("invalid Login")
    }
    _ , err := url.ParseRequestURI(urlString)
    if err != nil {
        return nil, errors.New("invalid URL")
    }
    newAccountwithTimeStamp := &AccountwithTimeStamp{
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        Account: Account{
        Login: login,
        Password: password,
        Url: urlString,
        },
    }

    if newAccountwithTimeStamp.Password == "" {
        newAccountwithTimeStamp.generatePassword(12)
    }
    return newAccountwithTimeStamp, nil

}