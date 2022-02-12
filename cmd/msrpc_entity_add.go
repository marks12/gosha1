package cmd

import (
	"github.com/abiosoft/ishell/v2"
	"errors"
	"strings"
	"regexp"
	"github.com/fatih/color"
)

func msrpcEntityAdd(c *ishell.Context) {

	yellow := color.New(color.FgYellow).SprintFunc()

	c.Println(yellow("Hello we start creating api for new entity"))

	entity, err := getName(c, false, "Entity")

	if err != nil {
		return
	}

	camelCase := strings.Title(entity)
	snakeCase := getLowerCase(entity)
	FirstLowerCase := GetFirstLowerCase(entity)

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
		AppendFile("./cnf/queue.go", "\n// generated code\nconst "+QueueConstant+" = \""+strings.ToLower(camelCase)+"\"")
	}

	c.Println("Created files:")

	c.Println(destinationFile1)
	c.Println(destinationFile2)

	c.Println("Конгретулэйшенс плеер уан, Ю уин. New entity " + camelCase + " was created. Bye")

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
