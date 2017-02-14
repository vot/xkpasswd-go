package main

import "fmt"
import "log"
import "time"
import "strings"
import "os"

import "math/rand"
import "io/ioutil"

import "github.com/urfave/cli"

func readLinesSimple(path string) ([]string, error) {
  content, err := ioutil.ReadFile(path)
  if err != nil {
    return nil, err
  }
  lines := strings.Split(string(content), "\n")
  return lines, nil
}

func getWords() (int, []string) {
  var url string = "words.txt"

  lines, err := readLinesSimple(url)
  if err != nil {
    log.Fatalf("readLines: %s", err)
    return 0, nil
  }

  count := len(lines)
  return count, lines
}

func getRandomWord () (string) {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  allWordsCount, allWords := getWords()
  var randomNumber int = r.Intn(allWordsCount);

  return allWords[randomNumber]
}

func generatePassword (pattern string, separator string) (string) {
  words := patternToArray(pattern)
  return strings.Join(words, separator)
}

func patternToArray(pattern string) ([]string) {
  // TODO iterate over each character in a pattern
  array := []string{getRandomWord(), getRandomWord(), getRandomWord()}

  return array
}

func getSeparatorForComplexity (level int) (string) {
  // TODO get a random character from string
  return "#"
}

func getPatternForComplexity (level int) (string) {
  rtn := "wswsw"

  // enforce limits
  if level < 1 {
    level = 1
  }

  if level > 6 {
    level = 6
  }

  // define level patterns
  if level == 1 {
    rtn = "wsw"
  }

  if level == 2 {
    rtn = "wswsw"
  }

  if level == 3 {
    rtn = "wswswsdd"
  }

  if level == 4 {
    rtn = "wswswswsdd"
  }

  if level == 5 {
    rtn = "wswswswswsd"
  }

  if level == 6 {
    rtn = "ddswswswswswsdd"
  }

  return rtn
}

func main() {
  // TODO download list of words if missing
  app := cli.NewApp()
  app.Name = "xkpasswd"
  app.Version = "0.0.1"
  app.Usage = "Memorable password generator"

  var inputComplexity int
  var inputPattern string
  var inputSeparator string
  var inputNumber int

  app.Flags = []cli.Flag {
    cli.IntFlag{
      Name:        "complexity, c",
      Value:       2,
      Usage:       "Define complexity (1-6)",
      Destination: &inputComplexity,
    },
    cli.StringFlag{
      Name:        "pattern, p",
      Value:       "",
      Usage:       "Define pattern (w = word, d = digit, s = separator)",
      Destination: &inputPattern,
    },
    cli.StringFlag{
      Name:        "separator, s",
      Value:       "",
      Usage:       "Define separator character",
      Destination: &inputSeparator,
    },
    cli.IntFlag{
      Name:        "number, n",
      Value:       1,
      Usage:       "Define number of passwords to generate",
      Destination: &inputNumber,
    },
  }

  app.Action = func(c *cli.Context) error {
    var pattern string
    var separator string

    if len(inputPattern) > 0 {
      pattern = inputPattern
    } else {
      pattern = getPatternForComplexity(inputComplexity)
    }

    if len(inputSeparator) > 0 {
      separator = inputSeparator
    } else {
      separator = getSeparatorForComplexity(inputComplexity)
    }

    for i := 0; i < inputNumber; i++ {
      fmt.Println(generatePassword(pattern, separator))
    }

    return nil
  }

  app.Run(os.Args)
}
