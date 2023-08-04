package sample_data

import "testing"

func TestGetSampleData(t *testing.T) {
	data := GetSampleData()
	if len(data) == 0 {
		t.Errorf("Error retrieving data")
	}
}
