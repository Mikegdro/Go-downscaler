package greetings

import (
	"testing"
	"regexp"
	"fmt"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)

	msg, err := Hello("Gladys")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {

	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
} 

func TestHellos(t *testing.T) {

	names := []string {
		"Mike",
		"Gladys",
	}

	msgs, err := Hellos([]string{"Mike", "Gladys"})

	if err != nil {
		t.Fatalf(`Hellos(%v) = %q, want "", error`, msgs, err)
	}

	// Loop to go through the map
	for key, val := range msgs {

		// Flags for the keys and values
		var keysOk bool = false
		var valsOk bool = false

		// Loop through the names list
		for _, name := range names {

			// We create a regexp to find the name
			want := regexp.MustCompile(`\b`+name+`\b`)

			// We try to find the name in the messages
			if want.MatchString(val) {
				valsOk = true
			}

			// We check for the keys to see if they are the names
			if name == key {
				keysOk = true
			}			
		}

		// In case one flag is set to false we throw the err
		if !keysOk || !valsOk {
			t.Fatalf(`What`)
		}

	}

	fmt.Println("names- ", names)
}