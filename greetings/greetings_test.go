package greetings

import(
  "testing"
  "regexp"
)

// TestHelloName calls greeting.Hello with a name,
// checking for valid return value.
func TestHelloName(t *testing.T){
  name:= "Goldy"
  want := regexp.MustCompile(`\b`+name+`\b`)
  msg, err := Hello("Goldy")
  if !want.MatchString(msg) || err != nil{
    t.Fatalf(`Hello("Goldy")= %q, %v, want match for %#q,nil`,msg,err,want)
  }
}

// TestHelloEmtoy calls greeting.Hello with an empty string,
// checking for an error
func TestHelloEmpty(t *testing.T){
  msg , err := Hello("")
  if msg!= "" || err ==nil {
    t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
  }
}
