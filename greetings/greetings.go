package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// This is a function
// we have variable followed by type
// at the end we have return type
func Hello(name string) (string,error){
  // If no name is provided, we return an error with message
  if name == ""{
	return "", errors.New("empty name")
  }
  // message variable is dynamically typeset
  // := operator is used for declaring and initializing
  //  message := fmt.Sprintf(selectRandomFormat(),name)
  message := fmt.Sprintf(selectRandomFormat())
  // so normally this would be the case for variable declaration
  var tmp string
  tmp = "GOはなんですか。"
  // interestingly code will not compile if there are unused variables
  fmt.Println(tmp)
  // the nil specifies the error if it did not occur
  return message,nil
}

func Hellos(names []string)(map[string]string , error){
	// mapping names to messages
	messages := make(map[string]string)
	// impressive for loop structure
	// Java doesn't allow sending back the index of array
	for _,name := range names{
		message , err := Hello(name)
		if(err != nil){
			return nil, err
		}
		messages[name] = message
	}
	return messages,nil
}


// sets initial value of rand
func init(){
  rand.Seed(time.Now().UnixNano())
}

// returns random quote
func selectRandomFormat() string{
  formats := []string{
	"Hi, %v. Welcome!",
	"Great to see you, %v!",
	"Whatever %v!",
  }
  return formats[rand.Intn(len(formats))]
}
