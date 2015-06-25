package globalquran

import (
	"strings"
)

// Constants for the service
const BASEURL string = "http://api.globalquran.com"
const APIKEY string = "945c29860c47be403296178612fb467d"
const READTYPE string = "quran-simple"

// Endpoints in a map, just in case the syntax changes
var endPoints = map[string]string{"ayah": "ayah", "surah": "surah", "juz": "juz", "page": "page"}

// surahAyahComb is either a string that includes the verse number in the Quran,
// or a combination of Surah:Ayah

func GetAyah(surahAyahComb string) map[string]interface{} {
	uri := strings.Join([]string{BASEURL, endPoints["ayah"], surahAyahComb, READTYPE}, "/")

	result := getApiRequest(uri)
	return result
}

func GetJuz(juzNumber string) map[string]interface{} {
	uri := strings.Join([]string{BASEURL, endPoints["juz"], juzNumber, READTYPE}, "/")

	result := getApiRequest(uri)
	return result
}

func GetSurah(surahNumber string) map[string]interface{} {
	uri := strings.Join([]string{BASEURL, endPoints["surah"], surahNumber, READTYPE}, "/")

	result := getApiRequest(uri)
	return result
}

func GetPage(pageNumber string) map[string]interface{} {
	uri := strings.Join([]string{BASEURL, endPoints["page"], pageNumber, READTYPE}, "/")

	result := getApiRequest(uri)
	return result
}
