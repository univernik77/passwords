package account

import (
	"time" 
	"encoding/json"
)


type Vault struct {
	Accounts []AccountwithTimeStamp`json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	return &Vault{
		Accounts: []AccountwithTimeStamp{},
		UpdatedAt: time.Now(),
	}
}

func (vault *Vault) AddAccount(acc AccountwithTimeStamp){
	vault.Accounts =  append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()

}

func (vault *Vault) ToBytes() ([]byte, error){
    file, err := json.Marshal(vault)
    if err != nil {
        return nil, err
    }
    return file, nil

}