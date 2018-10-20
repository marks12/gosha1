package cmd

import (
    "strings"
    "os"
    "io/ioutil"
    "fmt"
    "github.com/google/uuid"
    "math/rand"
    "time"
    "regexp"
)

func getCurrentDirName() string {

    pwd, _ := os.Getwd()
    folders := strings.Split(pwd, "/")
    return folders[len(folders)-1]
}

func CreateFile(file, content string) (err error) {

    err = ioutil.WriteFile(file, []byte(content), 0644)

    if err != nil {
        fmt.Println("Error creating", file)
        fmt.Println(err)
    }

    return
}

func getNewGuid() string {
    id := uuid.New()
    return id.String()
}

func generatePassword(n int) string {

    rand.Seed(time.Now().UnixNano())

    letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

    b := make([]rune, n)

    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

var newGuidRegexp = regexp.MustCompile("new-guid")
var newPassRegexp = regexp.MustCompile("new-pass")
