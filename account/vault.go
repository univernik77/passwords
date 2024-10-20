package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"
	"github.com/fatih/color"
)


type Vault struct {
	Accounts []AccountwithTimeStamp`json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	file, err :=  files.ReadFile("data.json")
	if err != nil {
		return &Vault{
		Accounts: []AccountwithTimeStamp{},
		UpdatedAt: time.Now(),
	}

	}
	var vault Vault
	err = json.Unmarshal(file, &vault)  
	if err != nil {
		color.Red(err.Error())
	}
	return &vault
}

func (vault *Vault) AddAccount(acc AccountwithTimeStamp){
	vault.Accounts =  append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) FindAccountbyUrl(url string) []AccountwithTimeStamp {
	var accounts []AccountwithTimeStamp
	for _, account := range vault.Accounts{
		isMatched := strings.Contains(account.Url, url)
		if isMatched{
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *Vault) DeleteAccountbyUrl(url string) bool {
	var accounts []AccountwithTimeStamp
	isDelete := false
	for _, account := range vault.Accounts{
		isMatched := strings.Contains(account.Url, url)
		if !isMatched{
			accounts = append(accounts, account)
		}
		isDelete = true
	}
	vault.Accounts = accounts
	vault.save()
	return isDelete
}

func (vault *Vault) ToBytes() ([]byte, error){
    file, err := json.Marshal(vault)
    if err != nil {
        return nil, err
    }
    return file, nil
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
	}
	files.WriteFile(data, "data.json")
}