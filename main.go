package main

import (
	"fmt"
    "errors"
	"math/rand/v2"
    "net/url"
)
       

type account struct {
    login string
    password string
    url string
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

func newAccount(login, password, urlString string) (*account, error) {
    _ , err := url.ParseRequestURI(urlString)
    if err != nil {
        return nil, errors.New("invalid URL")
    }
    return &account{
        login: login,
        password: password,
        url: urlString,
    }, nil

}

var letterRunes = []rune("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")

func main(){
    login := promtData("Введите логин:")
    password := promtData("Введите пароль:")
    url := promtData("Введите URL:")
    
  
    myAccount, err := newAccount(login, password, url)
    if err != nil {
        fmt.Println("Неверный формат URL")
        return
    }
    myAccount.generatePassword(12)
    myAccount.outputPassword()
	
}

func promtData(promt string) string {
    fmt.Print(promt + " ")
    var res string
    fmt.Scan(&res)
    return res
}



