# GoBinClassify
---
GoBinClassify is a library that makes it easy to classify into groups. The main idea of the library is to classify as many groups as possible using the minimum number of questions.The library can be useful when we need to divide users by interests.
I also apologize for my crooked English translation.
[Author's website](https://tatem.pythonanywhere.com)
---
For example, we need to divide users into 10 groups. And in order for us to do this, we just need to ask 4 questions. And for 300 groups there are only 9 questions.
The questions should be binary. For example: yes and no. We will be able to find out the number of questions using the method from our library. 
A good example is my soul color(RGB) detection test - [Color Soul](https://tatem.pythonanywhere.com/project/colorpeople/)


# A simple example:
## Installation
```bash
go get github.com/Daniil-7/GoBinClassify
```
## Let's find out how many questions we need:
```Go
package main

// Import library
import (
 "github.com/Daniil-7/GoBinClassify"
 "fmt"
 "log"
)


/*
CoutQustion accepts the number of groups you want to get.
The number should be:
1. The whole
2. More than 1
If the transfer is successful, it returns the number of responses that you must pass to the Answer method.
*/

func main() {
  qcout, err := GoBinClassify.CoutQustion(10)
  
  if err != nil {
    log.Fatal(err)
  }
  
  fmt.Println(qcout)
  // Outputs: 4
}
```
## Then suppose we conduct a survey on our website: 
1) Do you have many friends? 
1. Yes (save as 1) 
2. No (save as 0) 
2) Do you like busy places? 
1. Yes (save as 1) 
2. No (save as 0) 
3) Do you see your friends often? 
1. Yes (save as 1) 
2. No (save as 0) 
4) Are you popular with people? 
1. Yes (save as 1) 
2. No (save as 0) 


## Classification by groups:
```Go
package main

// Import library
import (
  "github.com/Daniil-7/GoBinClassify"
  "fmt"
  "log"
)

/*
Answer(x, y)
The Answer method accepts: number - number of groups and array - answers to questions. 
Array Criteria:
1. The length of the array must be equal to CoutQustion(x) - the number returned by the CoutQustion method with the first argument of the Answer method passed. 
2. The answers must be binary (0 or 1)
*/
func main() {
  ans1, err := GoBinClassify.Answer(10, []int{1, 0, 1, 0})
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(ans1)
  // Outputs: 6
  // This means that my answer puts me in group 6.
  // Starts from 1.
  // Let's try a few more times:
  
  ans2, err := GoBinClassify.Answer(10, []int{0, 0, 0, 0})
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(ans2) // 0
  
  ans3, err := GoBinClassify.Answer(10, []int{1, 1, 1, 1})
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(ans3) // 10
  
  ans4, err := GoBinClassify.Answer(10, []int{1, 1, 1, 0})
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(ans4) // 9
}
```
## Find out how many groups can be divided into by the number of questions
```Go
package main

// Import library
import (
  "github.com/Daniil-7/GoBinClassify"
  "fmt"
  "log"
)

/*
CoutGroups accepts the number of questions. Returns an array with a range of the number of groups.
The first element is the minimum number of groups for such a number of questions, the second is the maximum.
The minimum transmitted number is 1.
*/

func main() {
  gcout, err := GoBinClassify.CoutGroups(4)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(gcout)
  // Outputs: [9, 16]
}
```