// https://www.tutorialspoint.com/pascal/pascal_program_structure.htm
package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "regexp"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("pascal file?")
    return
  }
  fileName := os.Args[1]
  file, err := os.Open(fileName)
  if err != nil {
    fmt.Println("Open: Error!")
    return
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  i := 0

  var programName string
  var block bool = false

  for scanner.Scan() {
    currentLine := scanner.Text()
    currentLine = strings.TrimSpace(currentLine)
    // fmt.Printf("Line %d: %s\n", i, currentLine)

    // line 1
    if strings.HasPrefix(currentLine, "program") {
      fields := strings.Fields(currentLine)
      if len(fields) == 2 {
        programName = strings.TrimSuffix(fields[1], ";")
        fmt.Println("ProgramName:", programName)
      } else {
        fmt.Printf("ProgramName: Error! %s\n", currentLine)
        return
      }
    } else if currentLine == "begin" {
      block = true
    } else if strings.Contains(currentLine, "writeln") {
      re := regexp.MustCompile(`\((.*?)\)`)
      match := re.FindStringSubmatch(currentLine)

      if len(match) > 1 {
        arg := match[1]
        fmt.Println(arg[1 : len(arg)-1])
      } else {
        fmt.Println("WriteLn: Error! %s\n", currentLine)
        return
      }
    } else if currentLine == "end." {
      if !block {
        fmt.Println("No begin: Error!")
        return
      }
      block = false
      break
    }
    i++
  }
  if err := scanner.Err(); err != nil {
    fmt.Println("Scanner: Error!")
  }
}
