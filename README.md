# GolangRestSwag

## Go

go mod init xxx/xxx (launch go project)
go mod tidy  
go build  
go run .  

## Go release

To do the command in git bash  
https://goreleaser.com/quick-start/

## Run docker

Build image with => docker build . -t golangrestswag  
Run image in a container => docker run --publish 8080:8080  golangrestswag  

## Go routine

**Sample :** 

func main() {  
  go print1()  
  fmt.Println("0")  
  time.Sleep(2 * time.Second)  
}  

func print1() {  
	go print2()  
	fmt.Println("1")  
  }  
  
  func print2() {  
	time.Sleep(time.Second)  
	fmt.Println("2")  
  }  
  
  **OUTPUT :**  
  0  
  1  
  2  
