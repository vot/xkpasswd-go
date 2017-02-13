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

func generatePassword (compl int) (string) {
  separator := "#"
  s := []string{getRandomWord(), getRandomWord(), getRandomWord(), getRandomWord()}
  var output string = strings.Join(s, separator)
  return output
}

// TODO download list of words if missing
func main() {
  app := cli.NewApp()
  app.Name = "xkpasswd"
  app.Usage = "Generate memorable passwords"

  var complexity int

  app.Flags = []cli.Flag {
    cli.IntFlag{
      Name:        "complexity, c",
      Value:       3,
      Usage:       "Define complexity (1-5)",
      Destination: &complexity,
    },
  }

  app.Action = func(c *cli.Context) error {
    var output string = generatePassword(complexity)
    fmt.Println(output)
    return nil
  }

  app.Run(os.Args)
}
