package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)
       

type account struct {
    login string
    password string
    url string
}

type accountwithTimeStamp struct {
     createdAt time.Time
     updatedAt time.Time
     account
}

func (acc *account) outputPassword() {
    fmt.Println(acc.login, acc.password, acc.url)

}

func (acc *account) generatePassword(n int) {
    password := make([]rune, n)
    for i  := range password {
        password[i] = letterRunes[rand.IntN(len(letterRunes))]
    }
    acc.password = string(password)
}

func newAccountwithTimeStamp(login, password, urlString string) (*accountwithTimeStamp, error) {
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
        account: account{
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

var letterRunes = []rune("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")

func main(){
    login := promtData("Введите логин:")
    password := promtData("Введите пароль:")
    url := promtData("Введите URL:")
    
  
    myAccount, err := newAccountwithTimeStamp(login, password, url)
    if err != nil {
        fmt.Println("Неверный формат URL или Логин")
        return
    }
    myAccount.generatePassword(12)
    myAccount.outputPassword()
	
}

func promtData(promt string) string {
    fmt.Print(promt + " ")
    var res string
    fmt.Scanln(&res)
    return res
}



