package globalquran

import "testing"

func TestConfiguration(t *testing.T) {
	if len(APIKEY) == 0 {
		t.Error("APIKEY needed to work with GlobalQuran API")
	}

	if len(BASEURL) == 0 {
		t.Error("BASEURL needed to work with GlobalQuran API")
	}

	if len(READTYPE) == 0 {
		t.Error("No READTYPE specified for API, (quran-simple) May be ?")
	}
}

func TestGetApiRequest(t *testing.T) {
	result := GetAyah("116:1")

	if len(result) == 0 {
		t.Error("Error in request")
	}
}
