package main

import(
  // go basic dependencies 
  "fmt"
  "log"

  // imported dependencies
  "github.com/mohammadaziz313/p-go/greetings"
)

func main(){
  // keeping a log of events
  log.SetPrefix("greetings: ")
  // disabling certain properties
  log.SetFlags(0)
  // runs without error
  message,err := greetings.Hello("Lala land")
  if err !=nil {
	log.Fatal(err)
  }
  fmt.Println(message)

  // encounters an error
  message1 , err1 := greetings.Hello("")
  if err1 != nil{
	log.Fatal(err1)
  }
  fmt.Println(message1)
}
