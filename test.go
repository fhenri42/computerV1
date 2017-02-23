package main

import (
  "os"
  "fmt"
  "strings"
  "bytes"
  "strconv"
  "math"
)

func nub(before string, after string, last string, r bool, div bool) string {

  var u = 0
  var nb float64
  var nb1 float64
  str := strings.Split(before,"")
  str1 := strings.Split(after,"")

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
  u = 0
  for u < len(after) {

    if  u - 1 <= 0 || ((str1[u] >= "0" && str1[u] <= "9") || str1[u] == "-" || str1[u] == ".") {

      var tmp bytes.Buffer

      for u  < len(str1) && ((str1[u] >= "0" && str1[u] <= "9") || str1[u]  ==  "-" || str1[u] == ".")  {
        tmp.WriteString(str1[u])
        u++
      }
      s, _ := strconv.ParseFloat(tmp.String(), 64)
      nb1 = s
    }
    u++
  }
  var sum float64
  if div == true {
    sum = nb  / nb1
  } else {
    if r == true  {
    sum = nb + (nb1 * -1)
    } else {
      sum = nb + nb1
    }
  }

    sumString := strconv.FormatFloat(sum, 'f', -1, 64)

    if sum >= 0 {
      s := []string{"+", ""}
      sumString = strings.Join(s,sumString)
    }
    end := []string{sumString, ""}

    if last == "A"  {
      return strings.Join(end, "A")
      } else if last == "B" {
        return strings.Join(end, "B")
      } else if last == "C" {
        return strings.Join(end, "C")
      }
      return "NAN"
    }

    func  reduc(str []string) []string  {
      var i = 0
      var div = false
      fmt.Printf("%s\n",str)
      for i < len(str) {
        var p = i + 1
        for p < len(str) {
        div = false
          tmpB := strings.Split(str[i], "")
          tmpA := strings.Split(str[p], "")
          if tmpA[0] == "/" {
            div = true
          }
          if tmpA[len(tmpA) - 1] == tmpB[len(tmpB) - 1] {
            nb := nub(str[i], str[p], tmpB[len(tmpB) - 1], false, div)
            str[i] = nb
            str[p] = "0"
          }
          p++
        }
        i++
      }
      return str
    }

    func  reduction(before string, after string) []string {
      var i = 0

      splitbefore := strings.FieldsFunc(before, func(r rune) bool {
        switch r  {
        case '_':
          return true
        }
        return false
      })

      splitAfter:= strings.FieldsFunc(after, func(r rune) bool {
        switch r  {
        case '_':
          return true
        }
        return false
      })
      splitbefore = reduc(splitbefore)
      splitAfter = reduc(splitAfter)

      fmt.Printf("SplitStr = %s\nSplitAfter = %s\n", splitbefore, splitAfter)
      for i < len(splitbefore) {
        var p = 0
        for p < len(splitAfter) {
          tmpB := strings.Split(splitbefore[i], "")
          tmpA := strings.Split(splitAfter[p], "")
          if tmpA[len(tmpA) - 1] == tmpB[len(tmpB) - 1] {
            nb := nub(splitbefore[i], splitAfter[p], tmpB[len(tmpB) - 1], true, false)
            splitbefore[i] = nb
            splitAfter[p] = "L"
          }
          p++
        }
        i++
      }
      i = 0

      for i < len(splitAfter) {

        if splitAfter[i] != "L" {
          splitbefore = append(splitbefore, splitAfter[i])
        }
        i++
      }
      fmt.Printf("splitbefore = %s \n splitStr = %s\n",  splitAfter, splitbefore)
      return splitbefore
    }

    func resolution1(str []string) {
      var i = 0
      var p = 0

      for i < len(str) {
        tmpB := strings.Split(str[i], "")
        var u = 0
        for  u < len(tmpB) {
          if tmpB[u] == "+"  && tmpB[len(tmpB) - 1] == "A"{
            str[i] = strings.Replace(str[i], "+", "-", -1)
            u++
          }
          if tmpB[u] == "-" && tmpB[len(tmpB) - 1] == "A" {
            str[i] = strings.Replace(str[i], "-", "+", -1)
            u++
          }
          u++
        }
        i++
      }
      i = 0

      fmt.Printf("Resolution:")
      for  i < len(str) {
        tmpB := strings.Split(str[i], "")
        if tmpB[len(tmpB) - 1] == "A" {
           p = 0
          for p < len(tmpB) - 1 {
            fmt.Printf("%s",tmpB[p])
            p++
          }
        }
        i++
      }
      i = 0
      fmt.Printf("/")

    for  i < len(str) {
      tmpB := strings.Split(str[i], "")

        if tmpB[len(tmpB) - 1] == "B" {
          p = 0
          for p  < len(tmpB) - 1 {
            fmt.Printf("%s",tmpB[p])
            p++
          }
          fmt.Printf(" = X ")
        }
        i++
      }
    }

    func findNum(toFind string, str []string ) float64  {
      var i = 0
      for i < len(str) {
        tmp := strings.Split(str[i], "")
        fmt.Printf("%s",tmp)
        if tmp[len(tmp) - 1] == toFind {
          var u = 0
          for u < len(tmp) {

            if  u - 1 <= 0 || ((tmp[u] >= "0" && tmp[u] <= "9") || tmp[u] == "-") {

              var tmp1 bytes.Buffer

              for u  < len(tmp) && ((tmp[u] >= "0" && tmp[u] <= "9") || tmp[u]  ==  "-" || tmp[u] == "." || tmp[u]== " " ||  tmp[u] == "+")  {
      //          fmt.Printf("ZZZZZ%s",tmp[u])
                if tmp[u] != " " { tmp1.WriteString(tmp[u]) }
                u++
              }
              s, _ := strconv.ParseFloat(tmp1.String(), 64)
              return s
            }
            u++
          }
        }
        i++
      }
      return 1

    }
    func resolution2(str []string){
      var a float64 = findNum("C", str)
      var b float64 = findNum("B",str)
      var c float64 = findNum("A",str)
      fmt.Printf("\nA = %f ", a)
      fmt.Printf("\nB = %f ", b)
      fmt.Printf("\nC = %f ", c)

      var delta = math.Pow(b, 2) - (4 * a * c)
      fmt.Printf("\ndelta = %f ", delta)
      if (delta < 0) {
        fmt.Printf("\nil n y pas de solution dans R")
      }
      if (delta == 0) {
        var x0 float64 = -(b / (2 * a))
        fmt.Printf("\nil y a une solution x0 = %f", x0)
      }
      if (delta > 0 ) {
        var x1 float64 = (-b - math.Sqrt(delta)) / (2 * a)
        var x2 float64 = (-b + math.Sqrt(delta)) / (2 * a)
        fmt.Printf("\nil y a deux solution x1 = %f et x2 = %f", x1, x2)
      }

    }

    func parcing(str string) {

      var i = 0
      var degre2 = false

      str = strings.Replace(str, " ", "", -1)
      str = strings.Replace(str, "-", "_-", -1)
      str = strings.Replace(str, "+", "_+", -1)
      str = strings.Replace(str,"/", "_/", -1)
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
        fmt.Printf("Equation de degre 2\n")
        str1 := reduction(splitEqual[0], splitEqual[1])
        fmt.Printf("simplification : %s = 0\n", str1)
        resolution2(str1)
        } else {
          fmt.Printf("Equation de degre 1\n")
          str1 := reduction(splitEqual[0], splitEqual[1])
          resolution1(str1)
        }
      }

      func main () {
        var i = 2

        if i == len(os.Args) {
          parcing(os.Args[1])
        }
        return
      }
