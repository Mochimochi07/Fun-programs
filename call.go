package main

import "fmt"
import"time"

func newln(){
fmt.Println("\n")

}

func showcontacts(){

fmt.Println("1:Keybum,Lee")
fmt.Println("2:Hog rider")
fmt.Println("3:Krusty Krab")

}

func ridercall(){
for i := 1; i <= 5; i++ {
time.Sleep(time.Second)
    fmt.Println("HAAAWWWGGG RAAAADDDDAAAAHHHHH!!!!!")
  }
}



func main(){

var choice string


fmt.Println("Hi, who do you wanna call?")
newln()
showcontacts()
newln()
fmt.Scanln(&choice)

switch {

case choice == "1":
fmt.Println("*calls moe's tavern") 
time.Sleep(time.Second)
fmt.Println("Moe: Is there a lee keybum here anyone?")
time.Sleep(time.Second)
fmt.Println("Moe: I have a lee keybum on the phone?")
time.Sleep(time.Second)
fmt.Println("Everyone but moe: Then you should'nt be handling liquor")
time.Sleep(time.Second)
fmt.Println("Call ended")
break

case choice == "2":
fmt.Println("*calls hog rider")
time.Sleep(time.Second) 
fmt.Println("Archer:We need back up!:)
time.Sleep(time.Second)
fmt.Println("*Hogrider yells")
time.Sleep(time.Second)
ridercall()
fmt.Println("Call ended")
time.Sleep(time.Second)
fmt.Println(*A TEAM OF HOG RIDERS APPEAR!)
break

case choice =="3":

fmt.Println("*calls the Krusty Krab!") 
time.Sleep(time.Second)
fmt.Println("You:Is this the Krusty Krab?")
fmt.Println("No, this is Patrick")



Default:
fmt.Println("Calling...)
fmt.Println("No one answered")
fmt.Println("Call ended")
}
















}
