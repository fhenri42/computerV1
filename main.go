package main

import (
  "os"
  "fmt"
  "strings"
  "strconv"
  "bytes"
)

func nub(before string, size int) []float64 {

  var u = 0
  var t = 0

  str := strings.Split(before,"")
  numberTab := make([]float64, size)

  for u < len(str) {

    if  u - 1 <= 0 || ((str[u] >= "0" && str[u] <= "9" && str[u - 1] != "^") || str[u] == "-" || str[u] == ".") {

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

func simplification(before string, after string) string {

  var t = 0
  var b = strings.Split(before, "")
  var a = strings.Split(after, "")
  var k = 0
  var conteur1 = 0
  var conteur = 0
  var signe = false
  numberTab := nub(before, 4)
  resultTab := nub(after, 4)

  for t < len(b) {
    if (b[t] == "X" && b[t + 1] ==  "^") {
      k = 0
      conteur1 = 0
      for k < len(a) {
        if a[k] == "X" && a[k + 1] ==  "^" {
          if t + 2 < len(b) && k + 2 < len(a) && (b[t + 2] == a[k + 2]) {
            numberTab[conteur] +=  resultTab[conteur1] * - 1
          }
          conteur1++
        }
        k++
      }
      conteur++
    }
    t++
  }
  t  = 0
  for t < len(numberTab) {
   fmt.Printf("%f\n",numberTab[t])
   t++
 }
  t = 0
  conteur = 0
  for t < len(b)  {
    signe = false
    if  t - 1 <= 0 || ((b[t] >= "0" && b[t] <= "9" && b[t - 1] != "^") || b[t] == "-" || b[t] == ".") {
      var tmp bytes.Buffer

      if b[t] == "-"  { signe = true }
      for t  < len(b) && ((b[t] >= "0" && b[t] <= "9") || b[t]  ==  "-" || b[t] == ".")  {
        tmp.WriteString(b[t])
        t++
      }

      nb := strconv.FormatFloat(numberTab[conteur], 'f', -1, 64)

      if signe == true && numberTab[conteur] >= 0 {
        s := []string{"+", ""}
        nb = strings.Join(s,nb)
      }
      before = strings.Replace(before, tmp.String(), nb, -1)
      conteur++
    }
    t++
  }
  return before
}

func parcing(str string) {

  var i = 0
  var degre2 = false
  str = strings.Replace(str, " ", "", -1)
  verife := strings.Split(str, "")
  splitEqual := strings.Split(str,"=")

    if (splitEqual[0] == splitEqual[1]) {
      fmt.Printf(" %s it s not a solube equation.\n", str)
      return
    }
    for i < len(verife) {
      if  i + 2 <= len(verife) && (verife[i] == "^" && (verife[i + 1] >= "3" && verife [i + 1] <= "9")) {
        fmt.Printf(" %s it s not a solube equation.\n", str)
        return
      }
      if i + 2 < len(verife) && (verife[i] == "^" && (verife[i + 1] >= "1" && verife [i + 1] <= "9" && verife[i + 2] >= "0" && verife[i + 2] <= "9")) {
        fmt.Printf(" %s it s not a solube equation.\n", str)
        return
      }
      if  i + 1 <= len(verife) && (verife[i] == "^" && (verife[i + 1] == "2")) {
        degre2 = true
      }

      i++
    }
    if  splitEqual[1] != "0" && degre2 == true  {
      fmt.Printf("Not reduce forme: %s\n", str)
      str = simplification(splitEqual[0],splitEqual[1])
      fmt.Printf("Reduce forme: %s = 0\n", str)
    } else {
      fmt.Printf("Reduce forme: %s\n", str)
  }
  if degre2 == true {
    fmt.Printf("Equation de degre 2")
  } else {
    fmt.Printf("Equation de degre 1\n")
  //  str = resolution1(splitEqual[0], splitEqual[1])
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
