package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"web/model"
)

const (
	ROOT          = "database"
	PASSWORD_ROOT = ROOT + "/password"
)

func SavePassword(password model.TsPassword) {
	// jsondata, _ := json.Marshal(password)
	os.MkdirAll(filepath.FromSlash(PASSWORD_ROOT), 0777)
	path := filepath.FromSlash(PASSWORD_ROOT + "/" + "password.json")
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File does not exist")
		var passwordData model.TsPasswordList
		passwordData.Index++
		passwordData.Name = "password"
		password.Id = passwordData.Index
		passwordData.List = append(passwordData.List, password)
		jsondata, _ := json.Marshal(passwordData)
		ioutil.WriteFile(path, []byte(jsondata), 0777)
	} else {
		fmt.Println("File exist")
		var passwordData model.TsPasswordList
		json.Unmarshal(content, &passwordData)
		if password.Id == -1 {
			passwordData.Index++
			password.Id = passwordData.Index
			passwordData.List = append(passwordData.List, password)
			jsondata, _ := json.Marshal(passwordData)
			ioutil.WriteFile(path, []byte(jsondata), 0777)
			fmt.Println(string(content))
		} else {
			fmt.Println("haru001")
			fmt.Println(password)
			fmt.Println(passwordData)

			for i := 0; i < len(passwordData.List); i++ {
				if password.Id == passwordData.List[i].Id {
					fmt.Println("updating=" + password.Name)
					passwordData.List[i].Name = password.Name
					passwordData.List[i].Password = password.Password
					passwordData.List[i].Details = password.Details
					fmt.Println(passwordData.List[i])
				}
			}

			jsondata, _ := json.Marshal(passwordData)
			ioutil.WriteFile(path, []byte(jsondata), 0777)
			fmt.Println(string(content))
		}
	}
}

func fileExists(filename string) bool {
	_, err := os.Open(filename)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func ReadPassword() *model.TsPasswordList {
	if fileExists(PASSWORD_ROOT + "/password.json") {
		path := filepath.FromSlash(PASSWORD_ROOT + "/" + "password.json")
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("File does not exist")
			var listOfpassword model.TsPasswordList
			return &listOfpassword
		} else {
			var listOfpassword model.TsPasswordList
			json.Unmarshal(content, &listOfpassword)
			return &listOfpassword
		}
	} else {
		var listOfpassword model.TsPasswordList
		return &listOfpassword
	}
}

func DeletePassword(id int64) {
	fmt.Println("DeletePassword id=" + string(id))
	list := ReadPassword()
	var tempList []model.TsPassword
	for _, p := range list.List {
		if id != p.Id {
			tempList = append(tempList, p)
		}
	}
	list.List = tempList
	jsondata, _ := json.Marshal(list)
	path := filepath.FromSlash(PASSWORD_ROOT + "/" + "password.json")
	ioutil.WriteFile(path, []byte(jsondata), 0777)
}
