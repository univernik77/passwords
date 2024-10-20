package main

import (
    "github.com/fatih/color"
	"demo/password/account"
	"fmt"
)
       
func main(){
    fmt.Println("__Менеджер паролей__")
    vault := account.NewVault()
Menu:
    for {
        variant := getMenu()
        switch variant{
        case 1:
            createAccount(vault)
        case 2:
            findAccount(vault)
        case 3:
            deleteAccount(vault)
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

func createAccount(vault *account.Vault){

    login := promtData("Введите логин:")
    password := promtData("Введите пароль:")
    url := promtData("Введите URL:")
    
  
    myAccount, err := account.NewAccountwithTimeStamp(login, password, url)
    if err != nil {
        fmt.Println("Неверный формат URL или Логин")
        return
    }
    vault.AddAccount(*myAccount)
    
}

func findAccount(vault *account.Vault) {
    url := promtData("Введите url для поиска")
    accounts := vault.FindAccountbyUrl(url)
    if len(accounts) == 0 {
        color.Red("Аккаунтов не найдено")
    } 
    for _, account := range accounts {
        account.Output()
    } 
}

func deleteAccount(vault *account.Vault){
    url := promtData("Введите url для удаления")
    isDelete := vault.DeleteAccountbyUrl(url)
    if isDelete {
        color.Green("Удалено")
    } 
    color.Red("Не найдено")

}

func promtData(promt string) string {
    fmt.Print(promt + " ")
    var res string
    fmt.Scanln(&res)
    return res
}



