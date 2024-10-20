package main

import (
	"fmt"
	"math/rand/v2"
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

var letterRunes = []rune("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")

func main(){
    login := promtData("Введите логин:")
    //password := promtData("Введите пароль:")
    url := promtData("Введите URL:")
    
    myAccount := account{
        login: login,
        //password: password,
        url: url,
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



