package account

import (
	"encoding/json"
	"strings"
	"time"
	"demo/password/output"
)

type ByteReader interface{
	Read() ([]byte, error)
}

type ByteWriter interface{
	Write([]byte)
}


type Db interface {
	ByteReader
	ByteWriter
}
type Vault struct {
	Accounts []AccountwithTimeStamp`json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db Db
}

func NewVault(db Db) *VaultWithDb {
	file, err :=  db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []AccountwithTimeStamp{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)  
	if err != nil {
		output.PrintError(err)
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []AccountwithTimeStamp{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDb{
			Vault: vault,
			db: db,
		}
}

func (vault *VaultWithDb) AddAccount(acc AccountwithTimeStamp){
	vault.Accounts =  append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) FindAccountbyUrl(url string) []AccountwithTimeStamp {
	var accounts []AccountwithTimeStamp
	for _, account := range vault.Accounts{
		isMatched := strings.Contains(account.Url, url)
		if isMatched{
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDb) DeleteAccountbyUrl(url string) bool {
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

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.PrintError(err)
	}
	vault.db.Write(data)
	
}