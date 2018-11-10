package cmd

import (
    "strings"
    "os"
    "github.com/google/uuid"
    "math/rand"
    "time"
    "regexp"
    "gopkg.in/abiosoft/ishell.v2"
    "io/ioutil"
    "fmt"
    "errors"
)

func getCurrentDirName() string {

    pwd, _ := os.Getwd()
    folders := strings.Split(pwd, "/")
    return folders[len(folders)-1]
}

func CreateFile(file, content string, c *ishell.Context) (err error) {

    if _, err := os.Stat(file); !os.IsNotExist(err) {

        choice := c.MultiChoice([]string{"No", "Yes"}, "file " + file + " already exists, rewrite?")

        if choice == 0 {
            return errors.New("Cancel rewrite file")
        }
    }

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

func assignMsName(template string) string {
    var microserviceNameRegexp = regexp.MustCompile("{ms-name}")
    return microserviceNameRegexp.ReplaceAllString(template, getCurrentDirName())
}

func assignPass(template string) string {

    var newPassRegexp = regexp.MustCompile("{new-pass}")
    template = newPassRegexp.ReplaceAllString(template, generatePassword(8))

    return template
}

func assignGuid(template string) string {

    var reg, _ = regexp.Compile("{new-guid}")
    template = reg.ReplaceAllString(template, getNewGuid())

    for i:=1; i < 100; i++ {
        template = strings.Replace(template, "{new-guid" + string(i) + "}", getNewGuid(), -1)
    }

    return template
}

func assignEntityName(template, entityName string) string {
    var entityNameRegexp = regexp.MustCompile("{entity-name}")
    return entityNameRegexp.ReplaceAllString(template, entityName)
}

