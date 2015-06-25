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

func GetAyah(surahNumber string, ayahNumber string) map[string]interface{} {
	surahAyahComb := strings.Join([]string{surahNumber, ayahNumber}, ":")
	uri := strings.Join([]string{BASEURL, endPoints["ayah"], surahAyahComb, READTYPE}, "/")

	result := getApiRequest(uri)
	return result
}
