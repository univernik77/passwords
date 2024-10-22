package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"
	"github.com/fatih/color"
)

var menuVariants = []string{
    "1. Создать аккаунт",
    "2. Найти аккаунт по URL",
    "3. Найти аккаунт по логину",
    "4. Удалить аккаунт",
    "5. Выход ",
    "Выберите вариант",
}

var menu = map[string]func(*account.VaultWithDb) {
    "1":createAccount,
    "2":findAccountbyUrl,
    "3":findAccountbyLogin,
    "4":deleteAccount,
}
func main(){
    fmt.Println("__Менеджер паролей__")
    vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
    for {
        variant := promtData(menuVariants...)
        menuFunc := menu[variant]
        if menuFunc == nil {
            break Menu
        }
        menuFunc(vault)
    }
}

func createAccount(vault *account.VaultWithDb){
    login := promtData("Введите логин")
    password := promtData("Введите пароль")
    url := promtData("Введите URL")
    
  
    myAccount, err := account.NewAccountwithTimeStamp(login, password, url)
    if err != nil {
        output.PrintError("Неверный формат URL или Логин")
        return
    }
    vault.AddAccount(*myAccount)
    
}

func findAccountbyUrl(vault *account.VaultWithDb) {
    url := promtData("Введите url для поиска")
    accounts := vault.FindAccount(url, func (acc account.AccountwithTimeStamp, str string) bool {
    return strings.Contains(acc.Url, str)
    })
    outputResult(&accounts)
} 


func findAccountbyLogin(vault *account.VaultWithDb) {
    login := promtData("Введите login для поиска")
    accounts := vault.FindAccount(login, func (acc account.AccountwithTimeStamp, str string) bool {
    return strings.Contains(acc.Login, str)
    })
    outputResult(&accounts)
}

func outputResult(accounts *[]account.AccountwithTimeStamp) {
     if len(*accounts) == 0 {
        output.PrintError("Аккаунтов не найдено")
    } 
    for _, account := range *accounts {
        account.Output()
    } 
}

func deleteAccount(vault *account.VaultWithDb){
    url := promtData("Введите url для удаления")
    isDelete := vault.DeleteAccountbyUrl(url)

    if isDelete {
        color.Green("Удалено")
    } 
    output.PrintError("Не найдено")

}

func promtData(promt ...string) string {
    for i, line := range promt {
        if i == len(promt)-1{
            fmt.Printf("%v: ", line)
        }
        fmt.Println(line)
    }
    var res string
    fmt.Scanln(&res)
    return res
}



