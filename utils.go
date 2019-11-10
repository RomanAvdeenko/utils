package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
)

// CheckErr print Error and break the program
func CheckErr(msg string, err error) {
	if err != nil {
		log.Fatalln("ERROR!!! ", msg, err.Error())
	}
}

// GetFullFileName возвращает полное имя файла по именя файла
func GetFullFileName(fileName string) string {
	//Сначала откуда запускаем программу
	ex, _ := os.Executable()
	extFileName := filepath.Dir(ex) + "/" + fileName
	if _, err := os.Stat(extFileName); os.IsNotExist(err) {
		// пробуем из текуoей dir
		extFileName = "./" + fileName
		if _, err := os.Stat(extFileName); os.IsNotExist(err) {
			CheckErr("Config file not found:", err)
		}
	}
	return extFileName
}

// Содержит ли слайс элемент
func Contains(s interface{}, elem interface{}) bool {
	sVal := reflect.ValueOf(s)
	if sVal.Kind() == reflect.Slice {
		for i := 0; i < sVal.Len(); i++ {
			// !!! panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if sVal.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

// сделать GET
func RunGet(url string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("admin", "Nfhfc,ekm,f#")

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("Error: RunGet(), Do()")
		return err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%#v", string(respBody))

	return nil

}
