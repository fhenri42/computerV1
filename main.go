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

      nb, err := strconv.ParseFloat(tmp.String(), 64)
      if err != nil {
        fmt.Printf("an err occure %s", err)
      }
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
  var conteur = 0

  numberTab := nub(before, 3)
  resultTab := nub(after, 3)

  //  if a[k] >= "0" && a[k] <= "9" &&  a[k + 1] == "*" && a[]{

  //}
  for t < len(b) {
    if (b[t] == "X" && b[t + 1] ==  "^") {
      var k = 0
      var conteur1 = 0
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
    } else if (b[t] >= "0" && b[t] <=  "9" && b[k + 2] != "X") {
      var k = 0
      var conteur1 = 0
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
    }y
      t++
    }

    t = 0
    for t < len(numberTab) {
      fmt.Printf("%f\n",numberTab[t])
      t++
    }

    t = 0

    for t < len(resultTab) {
      fmt.Printf("%f\n",resultTab[t])
      t++
    }

    return "lol"
  }

  func parcing(str string) {

    //var i = 0

    str = strings.Replace(str, " ", "", -1)
  //  str = strings.Replace(str, "*X^0","", -1)
    fmt.Printf("str = %s\n", str)
    //  str = strings.Replace(str, "/X^0","", -1)
    splitEqual := strings.Split(str,"=")

    //  if splitEqual[1][0] != 0 && splitEqual[1][1] != nil  {
    str = simplification(splitEqual[0],splitEqual[1])
    //  }


    // splitStr:= strings.FieldsFunc(splitEqual[0], func(r rune) bool {
    // 	switch r {
    // 	case '+', '-':
    // 		return true
    // 	}
    // 	return false
    // })
    //
    //
    // for i < len(splitStr) {
    //   fmt.Printf("splitStr = %s\n", splitStr[i])
    //   i++
    // }
  }

  func main () {
    var i = 2

    if i == len(os.Args) {
      parcing(os.Args[1])
    }
    return
  }
