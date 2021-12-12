package json_encoding

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type test struct {
	input    Manager
	expected string
}

func TestEncodeManager(t *testing.T) {
	cases := []test{
		{Manager{
			FullName:       "Chris Smith",
			Position:       "CISO",
			Age:            32,
			YearsInCompany: 5,
		}, `{"full_name":"Chris Smith","position":"CISO","age":32,"years_in_company":5}`},
		{Manager{
			FullName:       "John Locke",
			Position:       "Employee",
			Age:            26,
			YearsInCompany: 0,
		}, `{"full_name":"John Locke","position":"Employee","age":26}`},
	}

	for _, tt := range cases {
		name := fmt.Sprintf("Manager %s", tt.input.FullName)
		t.Run(name, func(t *testing.T) {
			r, err := EncodeManager(&tt.input)
			if err != nil {
				t.Errorf("error not nil, got %v", err)
			}
			s, err := ioutil.ReadAll(r)
			if err != nil {
				t.Errorf("error not nil, got %v", err)
			}
			actual := string(s)
			if actual != tt.expected {
				t.Errorf("expected %s, actual %s", tt.expected, actual)
			}
		})
	}
}
