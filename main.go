package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)
       


func main(){
    files.WriteFile("Hello", "text.txt")
    login := promtData("Введите логин:")
    password := promtData("Введите пароль:")
    url := promtData("Введите URL:")
    
  
    myAccount, err := account.NewAccountwithTimeStamp(login, password, url)
    if err != nil {
        fmt.Println("Неверный формат URL или Логин")
        return
    }
    myAccount.OutputPassword()
    
	
}

func promtData(promt string) string {
    fmt.Print(promt + " ")
    var res string
    fmt.Scanln(&res)
    return res
}



