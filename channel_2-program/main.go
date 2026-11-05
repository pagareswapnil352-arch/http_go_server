package main

import "fmt"


func main(){


	channel1 :=make(chan int)

     	channel2 :=make(chan int)


		go func(){

			channel1 <- 100
		}()

       go func(){

			channel2 <- 1000
		}()


		for i:=1;i<=2;i++{

			select{
             
				case value1 := <- channel1
				fmt.Println("value is:",value1)

				case value2 :=<- channel2
				fmt.Println("the value is:",value2)
			}
		}
}
