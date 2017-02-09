package main

import "fmt"
import "log"
import "time"
import "math/rand"

import "io/ioutil"
import "strings"

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

func main() {
  separator := "#"
  s := []string{getRandomWord(), getRandomWord(), getRandomWord(), getRandomWord()}
  var output string = strings.Join(s, separator)
  fmt.Println(output)
}

