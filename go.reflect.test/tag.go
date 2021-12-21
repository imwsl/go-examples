package main 
  
import ( 
    "fmt"
    "reflect"
) 
  
type Employee struct { 
    ID      string `auto_increment:"true" increment:"1"` 
    Name    string `varchar:"255"` 
    Surname string `"varchar:"255"` 
} 
  
// Main function 
func main() { 
    e:= Employee{} 
    // c variable represents table columns 
    c:= reflect.TypeOf(e).Field(0).Tag 
    fmt.Printf("%s\n\n", c) 
  
    // use of Get method 
    g:= c.Get("increment") 
    fmt.Printf("Get method:%s\n", g) 
}
