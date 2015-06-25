package globalquran

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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
