package main

import (
  "os"
  "fmt"
  "strings"
  "bytes"
  "strconv"
)
func nub(before string, after string) string {

  var u = 0
  var nb float64
  var nb1 float64
  str := strings.Split(before,"")

  for u < len(before) {

    if  u - 1 <= 0 || ((str[u] >= "0" && str[u] <= "9") || str[u] == "-" || str[u] == ".") {

      var tmp bytes.Buffer

      for u  < len(str) && ((str[u] >= "0" && str[u] <= "9") || str[u]  ==  "-" || str[u] == ".")  {
        tmp.WriteString(str[u])
        u++
      }
      s, _ := strconv.ParseFloat(tmp.String(), 64)
      nb = s
    }
    u++
  }

  for u < len(after) {

    if  u - 1 <= 0 || ((str[u] >= "0" && str[u] <= "9") || str[u] == "-" || str[u] == ".") {

      var tmp bytes.Buffer

      for u  < len(str) && ((str[u] >= "0" && str[u] <= "9") || str[u]  ==  "-" || str[u] == ".")  {
        tmp.WriteString(str[u])
        u++
      }
      s, _ := strconv.ParseFloat(tmp.String(), 64)
      nb1 = s
    }
    u++
  }
  fmt.Printf("%f \n %f",nb,nb1)

  return "lol"
}

func  resolution1(before string, after string) string {
  var i = 0

  splitbefore:= strings.FieldsFunc(before, func(r rune) bool {
    switch r {
    case '+', '-':
      return true
    }
    return false
  })
  splitAfter:= strings.FieldsFunc(after, func(r rune) bool {
    switch r {
    case '+', '-':
      return true
    }
    return false
  })
  i = 0
  for i < len(splitbefore) {
    var p = 0
    fmt.Printf("splitStr = %s\n", splitbefore[i])
    for p < len(splitAfter) {
      tmpB := strings.Split(splitbefore[i], "")
      tmpA := strings.Split(splitAfter[p], "")
      fmt.Printf("%s, %s", tmpA, tmpB)
      if tmpA[len(tmpA) - 1] == tmpB[len(tmpB) - 1] {
        nb := nub(splitbefore[i], splitAfter[p]) /*dois return le resulta bien formater*/
        splitbefore[i] = nb
      }
      p++
    }
    i++
  }
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
    i++
  }
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
  }

  func main () {
    var i = 2

    if i == len(os.Args) {
      parcing(os.Args[1])
    }
    return
  }
