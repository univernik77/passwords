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
    login string
    password string
    url string
}

type accountwithTimeStamp struct {
     createdAt time.Time
     updatedAt time.Time
     Account
}

func (acc *Account) OutputPassword() {
    color.Yellow(acc.login)
    color.Yellow(acc.password)
    color.Yellow(acc.url)

}

func (acc *Account) generatePassword(n int) {
    password := make([]rune, n)
    for i  := range password {
        password[i] = letterRunes[rand.IntN(len(letterRunes))]
    }
    acc.password = string(password)
}

func NewAccountwithTimeStamp(login, password, urlString string) (*accountwithTimeStamp, error) {
    if login == ""{
        return nil, errors.New("invalid Login")
    }
    _ , err := url.ParseRequestURI(urlString)
    if err != nil {
        return nil, errors.New("invalid URL")
    }
    newAccountwithTimeStamp := &accountwithTimeStamp{
        createdAt: time.Now(),
        updatedAt: time.Now(),
        Account: Account{
        login: login,
        password: password,
        url: urlString,
        },
    }

    if newAccountwithTimeStamp.password == "" {
        newAccountwithTimeStamp.generatePassword(12)
    }
    return newAccountwithTimeStamp, nil

}