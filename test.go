package main

import (
  "os"
  "fmt"
  "strings"
  "bytes"
  "strconv"
)
func nub(before string, size int) []float64 {

  var u = 0
  var t = 0

  str := strings.Split(before,"")
  numberTab := make([]float64, size)

  for u < len(str) {

    if  u - 1 <= 0 || ((str[u] >= "0" && str[u] <= "9") || str[u] == "-" || str[u] == ".") {

      var tmp bytes.Buffer

      for u  < len(str) && ((str[u] >= "0" && str[u] <= "9") || str[u]  ==  "-" || str[u] == ".")  {
        tmp.WriteString(str[u])
        u++
      }
      nb, _ := strconv.ParseFloat(tmp.String(), 64)
      numberTab[t] = nb
      t++
    }
    u++
  }
  return numberTab
}

func  resolution1(before string, after string) string {
  numberTab := nub(before, 4)
  resultTab := nub(after, 4)
  fmt.Printf("%f", numberTab)
  fmt.Printf("%f", resultTab)
  return "lol"

}
func parcing(str string) {

  var i = 0
  var degre2 = false

  str = strings.Replace(str, " ", "", -1)
  str = strings.Replace(str, "X^0","A", -1)
  str = strings.Replace(str, "X^1","B", -1)
  str = strings.Replace(str, "X^2","C", -1)
  str = strings.Replace(str, "X^","D", -1)
  verife := strings.Split(str, "")

  splitEqual := strings.Split(str,"=")

  for i < len(verife) {
    if  i  <= len(verife) && (verife[i] == "D") {
      fmt.Printf(" %s it s not a solube equation.\n", str)
      return
    }
    if  i  <= len(verife) && (verife[i] == "C") { degre2 = true}
  }

  splitbefore:= strings.FieldsFunc(splitEqual[0], func(r rune) bool {
    switch r {
    case '+', '-':
      return true
    }
    return false
  })
  splitAfter:= strings.FieldsFunc(splitEqual[1], func(r rune) bool {
    switch r {
    case '+', '-':
      return true
    }
    return false
  })

  if (splitEqual[0] == splitEqual[1]) {
    fmt.Printf(" %s it s not a solube equation.\n", str)
    return
  }

  if degre2 == true {
    fmt.Printf("Equation de degre 2")
  } else {
    fmt.Printf("Equation de degre 1\n")
    str = resolution1(splitEqual[0], splitEqual[1])
    fmt.Printf("Resolution: %s\n", str)
  }

  for i < len(splitbefore) {
    fmt.Printf("splitStr = %s\n", splitbefore[i])
    i++
  }
  i = 0
  for i < len(splitAfter) {
    fmt.Printf("splitAfter = %s\n", splitAfter[i])
    i++
  }
}


func main () {
  var i = 2

  if i == len(os.Args) {
    parcing(os.Args[1])
  }
  return
}
