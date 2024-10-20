package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)
       
func main(){
    fmt.Println("__Менеджер паролей__")
Menu:
    for {
        variant := getMenu()
        switch variant{
        case 1:
            createAccount()
        case 2:
            findAccount()
        case 3:
            deleteAccount()
        default:
            break Menu
        }
    }
    
    //myAccount.OutputPassword()
}

func getMenu() int {
    var variant int 
    fmt.Println("Выберите вариант: ")
    fmt.Println("1. Создать аккаунт: ")
    fmt.Println("2. Найти аккаунт: ")
    fmt.Println("3. Удалить аккаунт: ")
    fmt.Println("4. Выход ")
    fmt.Scan(&variant)
    return variant
}

func createAccount(){

    login := promtData("Введите логин:")
    password := promtData("Введите пароль:")
    url := promtData("Введите URL:")
    
  
    myAccount, err := account.NewAccountwithTimeStamp(login, password, url)
    if err != nil {
        fmt.Println("Неверный формат URL или Логин")
        return
    }
    vault := account.NewVault()
    vault.AddAccount(*myAccount)
    data, err := vault.ToBytes()
    if err != nil {
        fmt.Println("Не удалось преобразовать в JSON")
        return
    }
    files.WriteFile(data, "data.json")
}

func findAccount() {

}

func deleteAccount(){

}

func promtData(promt string) string {
    fmt.Print(promt + " ")
    var res string
    fmt.Scanln(&res)
    return res
}



