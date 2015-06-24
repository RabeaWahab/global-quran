package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const BASEURL string = "http://api.globalquran.com"
const APIKEY string = "945c29860c47be403296178612fb467d"
const READTYPE string = "quran-simple"

var endPoints = map[string]string{"ayah": "ayah", "surah": "surah", "juz": "juz", "page": "page"}

func main() {
	var result map[string]interface{}
	result = GetAyah(110, 6)

	fmt.Println(result)
}

func GetAyah(surahNumber int, ayahNumber int) map[string]interface{} {
	surahAyahComb = strings.Join([]string{surahNumber, ayahNumber}, ":")
	uri = strings.Join([]string{BASEURL, ENDPOINTS["ayah"], surahAyahComb, READTYPE}, "/")

	fmt.Println(uri)
	panic()

	result := getApiRequest(uri)

	return result
}

func getApiRequest(Uri string) map[string]interface{} {
	params := []string{Uri, APIKEY}
	Uri = strings.Join(params, "?key=")
	response, error := http.Get(Uri)
	var result map[string]interface{}

	if error != nil {
		fmt.Println("%s", error)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, error := ioutil.ReadAll(response.Body)
		if error != nil {
			fmt.Println("%s", error)
			os.Exit(1)
		} else {
			error := json.Unmarshal(contents, &result)

			if error != nil {
				fmt.Println("%s", error)
				os.Exit(1)
			}
		}
	}

	return result
}
