package main

import "fmt"
import "log"
import "time"
import "strings"
import "strconv"
import "os"

import "math/rand"
import "io/ioutil"

import "github.com/urfave/cli"

func readLines(path string) ([]string, error) {
  content, err := ioutil.ReadFile(path)
  if err != nil {
    return nil, err
  }
  lines := strings.Split(string(content), "\n")
  return lines, nil
}

func getWords() (int, []string) {
  var path string = "xkpasswd-words.txt"

  lines, err := readLines(path)
  if err != nil {
    log.Fatalf("readLines: %s", err)
    return 0, nil
  }

  count := len(lines)
  return count, lines
}

func getRandomWord() (string) {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  allWordsCount, allWords := getWords()
  var randomNumber int = r.Intn(allWordsCount);

  return allWords[randomNumber]
}

func getRandomDigit() (int) {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  var digit int = r.Intn(9)
  return digit
}

func generatePassword(pattern string, separator string) (string) {
  words := patternToArray(pattern, separator)
  return strings.Join(words, "")
}

func patternToArray(pattern string, separator string) ([]string) {
  array := make([]string, 0)

  for i := 0; i < len(pattern); i++ {
    if (string(pattern[i]) == "w") {
      array = append(array, getRandomWord())
    }

    if (string(pattern[i]) == "d") {
      array = append(array, strconv.Itoa(getRandomDigit()))
    }

    if (string(pattern[i]) == "s") {
      array = append(array, separator)
    }
  }

  return array
}

func getSeparatorForComplexity(level int) (string) {
  // enforce limits
  if level < 1 {
    level = 1
  }

  if level > 6 {
    level = 6
  }

  // define allowed separators
  var separatorPool string = "#.-=+_"

  if level == 5 {
    separatorPool = "#.-=+_!$*:~?"
  }

  if level == 6 {
    separatorPool = "#.-=+_!$*:~?%^&;"
  }

  rtn := separatorPool[getRandomDigit()]

  return string(rtn)
}

func getPatternForComplexity(level int) (string) {
  // enforce limits
  if level < 1 {
    level = 1
  }

  if level > 6 {
    level = 6
  }

  // define the pattern
  var rtn string

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
  // TODO flag to specify list of words, fallback to default (+download if missing)
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
