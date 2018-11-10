package cmd

import (
    "gopkg.in/abiosoft/ishell.v2"
    "errors"
    "strings"
    "regexp"
    "io/ioutil"
    "fmt"
    "github.com/fatih/color"
    "bytes"
    "os"
    "log"
    "bufio"
    "io"
)

func msrpcEntityAdd(c *ishell.Context) {

    yellow := color.New(color.FgYellow).SprintFunc()

    c.Println(yellow("Hello we start creating api for new entity"))

    entity, err := getName(c)

    if err !=nil {
        return
    }

    camelCase := strings.Title(entity)
    snakeCase := getLowerCase(entity)
    FirstLowerCase := getFirstLowerCase(entity)

    sourceFile := "./api/new_entity.go"
    destinationFile1 := "./api/" + snakeCase + ".go"


    QueueConstant, isNew, err := getQueueName(c)

    if err != nil {
        return
    }

    CopyFile(
        sourceFile,
        destinationFile1,
        []string{"NewEntity", "newEntity"},
        []string{camelCase, FirstLowerCase},
        c)

    sourceFile = "./cnf/route_new_entity.go"
    destinationFile2 := "./cnf/route_" + snakeCase + ".go"

    CopyFile(
        sourceFile,
        destinationFile2,
        []string{"NewEntity", "newEntity", "QueueNameCore"},
        []string{camelCase, FirstLowerCase, QueueConstant},
        c)

    if isNew {
        AppendFile("./cnf/queue.go", "\n// generated code\nconst " + QueueConstant + " = \"" + strings.ToLower(camelCase) + "\"")
    }

    c.Println("Created files:" )

    c.Println(destinationFile1)
    c.Println(destinationFile2)

    c.Println("Конгретулэйшенс плеер уан, Ю уин. New entity " + camelCase + " was created. Bye" )

}

func getQueueName(c *ishell.Context) (name string, IsNew bool, err error) {


    queues := getAvalableQeues()

    queues = append([]string{"Create new queue name"}, queues...)

    choice := c.MultiChoice(queues, "Please select exists queue or create new in file ?")

    if choice == 0 {

        c.Println("Please, enter new queueName:")
        queue := strings.Title(c.ReadLine())

        reg, _ := regexp.Compile("^QueueName")
        queue = reg.ReplaceAllString(queue, "")

        queue = "QueueName" + queue
        c.Println("New Queue name is", queue)

        IsNew = true

        return queue, IsNew, nil

    } else {

        if choice < 1 {

            err = errors.New("cancel creating queue")
            return

        }

        return queues[choice], IsNew, nil
    }

    return
}

func getAvalableQeues() (qArr []string) {

    for _, line := range getFileLines("./cnf/queue.go") {

        arr := strings.Split(line, " ")

        if arr[0] == "const" {
            qArr = append(qArr, arr[1])
        }
    }

    return
}

func getFileLines(filepath string) (strArr []string) {

    f, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
    if err != nil {
        log.Fatalf("open file error: %v", err)
        return
    }
    defer f.Close()

    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                break
            }

            log.Fatalf("read file line error: %v", err)
            return
        }

        strArr = append(strArr, line)
    }

    return
}

func CopyFile(sourceFile, destinationFile string, replaceFrom []string, replaceTo []string, c *ishell.Context) {

    input, err := ioutil.ReadFile(sourceFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    newCode := string(input)

    for key, textReplaceFrom := range replaceFrom {

        reg, _ := regexp.Compile(textReplaceFrom)
        newCode = reg.ReplaceAllString(newCode, replaceTo[key])
    }

    err = ioutil.WriteFile(destinationFile, []byte(newCode), 0644)

    if err != nil {
        fmt.Println("Error creating", destinationFile)
        fmt.Println(err)
        return
    }
}

func AppendFile(sourceFile, appendString string) {

    input, err := ioutil.ReadFile(sourceFile)
    if err != nil {
        fmt.Println(err)
        return
    }

    err = ioutil.WriteFile(sourceFile, []byte(string(input) + appendString), 0644)

    if err != nil {
        fmt.Println("Error appending", sourceFile)
        fmt.Println(err)
        return
    }
}

func getLowerCase(ent string) string {

    var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
    var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

    snake := matchFirstCap.ReplaceAllString(ent, "${1}_${2}")
    snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
    return strings.ToLower(snake)
}

func getFirstLowerCase(s string) string {

    if len(s) < 2 {
        return strings.ToLower(s)
    }

    bts := []byte(s)

    lc := bytes.ToLower([]byte{bts[0]})
    rest := bts[1:]

    return string(bytes.Join([][]byte{lc, rest}, nil))
}

func getName(c *ishell.Context) (string, error) {

    green := color.New(color.FgGreen).SprintFunc()
    red := color.New(color.FgRed).SprintFunc()

    c.Println("Please type new entityName, like \"NewEntity\" or \"Entity\" or exit return")

    entityName := strings.Title(c.ReadLine())

    if entityName == "exit" {
        return "", errors.New("exit")
    }

    if len(entityName) > 0 {

        choice := c.MultiChoice([]string{
            "Yes",
            "No",
            "Cancel creating new entity",
        }, "Correct name " + green(entityName) + " ?")

        if choice == 0 {

            c.Println("New entity name is", red(entityName))

            return entityName, nil

        } else if choice == 2 {

            return "", errors.New("exit")
        }
    }

    return getName(c)
}
