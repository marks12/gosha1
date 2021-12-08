package cmd

import (
    "bufio"
    "bytes"
    "errors"
    "fmt"
    "gosha/mode"
    "gosha/settings"
    "gosha/types"
    "io"
    "io/ioutil"
    "log"
    "math/rand"
    "os"
    "path/filepath"
    "reflect"
    "regexp"
    "runtime"
    "strings"
    "time"
    "unicode"

    "github.com/fatih/color"
    "github.com/google/uuid"
    "gopkg.in/abiosoft/ishell.v2"
)

type RegularFind struct {
    BoolResult   bool
    StringResult string
    ArrayResult  []string
}

type ByCase []string

func (s ByCase) Len() int      { return len(s) }
func (s ByCase) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ByCase) Less(i, j int) bool {
    iRunes := []rune(s[i])
    jRunes := []rune(s[j])

    max := len(iRunes)
    if max > len(jRunes) {
        max = len(jRunes)
    }

    for idx := 0; idx < max; idx++ {
        ir := iRunes[idx]
        jr := jRunes[idx]

        lir := unicode.ToLower(ir)
        ljr := unicode.ToLower(jr)

        if lir != ljr {
            return lir < ljr
        }

        // the lowercase runes are the same, so compare the original
        if ir != jr {
            return ir < jr
        }
    }

    return false
}

func IsDirExists(path string) bool {

    if _, err := os.Stat(path); !os.IsNotExist(err) {
        return true
    }

    return false
}

func IsAllDirExists(dirs []string) bool {

    for _, d := range dirs {
        if !IsDirExists(d) {
            return false
        }
    }
    return true
}

func GetCurrentApp() types.App {

    app := types.App{}

    if IsAllDirExists(settings.UsualDefaultStructure) {
        app.IsAppExists = true
    } else {
        app.IsAppExists = false
    }

    return app
}

func GetOsArgument(arg string) (RegularFind, error) {

    arg = strings.TrimSpace(arg)
    arg = strings.Replace(arg, "--", "", -1)

    for _, a := range os.Args {

        a = strings.TrimSpace(a)
        a = strings.Replace(a, "--", "", -1)

        if a == arg || strings.ToLower(a) == strings.ToLower(arg+`=true`) {
            return RegularFind{BoolResult: true}, nil
        }

        re := regexp.MustCompile(arg + `=([a-zA-Z0-9а-яА-Я!#$%&()*+,./:;\<>?@ _{|}~\-"^=\[\]]*)`)
        as := re.FindSubmatch([]byte(a))

        if len(as) > 1 {

            arrs := []string{}

            for _, s := range as {
                arrs = append(arrs, string(s))
            }

            return RegularFind{
                StringResult: string(as[1]),
                ArrayResult:  arrs,
            }, nil
        }
    }

    return RegularFind{}, errors.New("Not found argument " + arg)
}

func InteractiveEcho(arrStr []string) {
    if mode.IsInteractive() {
        for _, t := range arrStr {
            fmt.Println(t)
        }
    }
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
    exists = false
    index = -1

    switch reflect.TypeOf(array).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(array)

        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
                index = i
                exists = true
                return
            }
        }
    }

    return
}

func getCurrentDirName() string {

    pwd, _ := os.Getwd()
    folders := strings.Split(pwd, "/")
    return folders[len(folders)-1]
}

func getGoModName() (name string, err error) {
    b, err := ioutil.ReadFile("./go.mod")
    if err != nil {
        return
    }
    firstLine := strings.Split(string(b), "\n")[0]
    name = strings.Replace(firstLine, "module ", "", -1)
    return strings.TrimSpace(name), err
}

func GetCurrentAppName() string {
    name, err := getGoModName()
    if err != nil || len(name) < 1 {
        return getCurrentDirName()
    }
    return name
}

func GetGoVersion() string {
    raw := runtime.Version()

    reg, err := regexp.Compile("[a-zA-Z]+")
    if err != nil {
        log.Fatal(err)
    }
    fullVersion := reg.ReplaceAllString(raw, "")

    versionArray := strings.Split(fullVersion, ".")
    if len(versionArray) < 3 {
        return fullVersion
    }
    return fmt.Sprintf("%s.%s", versionArray[0], versionArray[1])
}

func CreateFile(file, content string, c *ishell.Context) (err error) {

    var choice int
    var fmode os.FileMode = 0644

    if _, err := os.Stat(file); !os.IsNotExist(err) {

        if mode.IsInteractive() {
            choice = c.MultiChoice([]string{"No", "Yes"}, "file "+file+" already exists, rewrite?")
        } else {
            choice = 1
        }

        if choice == 0 {
            return errors.New("Cancel rewrite file: " + file)
        }
    }

    reg := regexp.MustCompile("\\.sh$")
    if reg.MatchString(file) {
        fmode = 0755
    }

    dir := filepath.Dir(file)
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        os.MkdirAll(dir, 0755)
    }

    content = assignMsName(content)

    err = ioutil.WriteFile(file, []byte(content), fmode)

    if err != nil {
        fmt.Println("Error creating", file)
        fmt.Println(err)
    }

    return
}

func CreateFileIfNotExists(file, content string, c *ishell.Context) (err error) {
    if _, err := os.Stat(file); !os.IsNotExist(err) {
        return errors.New("Cancel rewrite file: " + file)
    }
    return CreateFile(file, content, c)
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

func AssignVar(template string, variable string, value string) string {
    var microserviceNameRegexp = regexp.MustCompile(variable)
    return microserviceNameRegexp.ReplaceAllString(template, value)
}

func assignMsName(template string) string {
    var microserviceNameRegexp = regexp.MustCompile("{ms-name}")
    return microserviceNameRegexp.ReplaceAllString(template, GetCurrentAppName())
}

func assignPass(template, pass string) string {

    var newPassRegexp = regexp.MustCompile("{new-pass}")
    template = newPassRegexp.ReplaceAllString(template, pass)

    return template
}

func assignGuid(template string) string {

    var reg, _ = regexp.Compile("{new-guid}")
    template = reg.ReplaceAllString(template, getNewGuid())

    for i := 1; i < 100; i++ {
        template = strings.Replace(template, "{new-guid"+string(i)+"}", getNewGuid(), -1)
    }

    return template
}

func assignEntityName(template, entityName string) string {
    var entityNameRegexp = regexp.MustCompile("{entity-name}")
    return entityNameRegexp.ReplaceAllString(template, entityName)
}

func assignCurrentDateTime(template string) string {

    var format = "2006.01.02 15:04:05"

    var entityNameRegexp = regexp.MustCompile("{current-date-time}")
    res := entityNameRegexp.ReplaceAllString(template, (time.Now()).Format(format))

    entityNameRegexp = regexp.MustCompile("{current-date}")
    res = entityNameRegexp.ReplaceAllString(res, (time.Now()).Format(format))

    return res
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

func CopyFile(sourceFile, destinationFile string, replaceFrom []string, replaceTo []string, c interface{}) {

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

    err = ioutil.WriteFile(sourceFile, []byte(string(input)+appendString), 0644)

    if err != nil {
        fmt.Println("Error appending", sourceFile)
        fmt.Println(err)
        return
    }
}

func getLowerCase(ent string) string {

    var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
    var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

    snake := matchFirstCap.ReplaceAllString(ent, "${1}_${2}")
    snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
    return strings.ToLower(snake)
}

func GetFirstLowerCase(s string) string {

    if len(s) < 2 {
        return strings.ToLower(s)
    }

    bts := []byte(s)

    lc := bytes.ToLower([]byte{bts[0]})
    rest := bts[1:]

    return string(bytes.Join([][]byte{lc, rest}, nil))
}

func getName(c *ishell.Context, IsExistsStruct bool, targetName string) (string, error) {

    var selectedName string

    TargetName := strings.Title(targetName)

    green := color.New(color.FgGreen).SprintFunc()
    red := color.New(color.FgRed).SprintFunc()

    if IsExistsStruct {
        InteractiveEcho([]string{
            "Please type exists " + TargetName + ", like \"New" + TargetName + "\" or \"" + TargetName + "\" or exit return",
        })
    } else {
        InteractiveEcho([]string{
            "Please type new " + TargetName + ", like \"New" + TargetName + "\" or \"" + TargetName + "\" or exit return",
        })
    }

    char := c.ReadLine()

    selectedName = strings.Title(char)

    if selectedName == "Exit" {
        c.Println("Operation canceled!")
        return "", errors.New("exit")
    }

    if len(selectedName) > 0 {

        choice := c.MultiChoice([]string{
            "Yes",
            "No",
            "Cancel",
        }, "Correct "+targetName+" "+green(selectedName)+" ?")

        if choice == 0 {

            InteractiveEcho([]string{
                TargetName + " name is:", red(selectedName),
            })

            return selectedName, nil

        } else if choice == 2 {

            c.Println("Operation canceled!")
            return "", errors.New("exit")
        }

        return getName(c, IsExistsStruct, targetName)
    }

    _, b := c.ReadLineErr()
    if b != nil && b.Error() == "Interrupt" {
        c.Println("Operation canceled!")
        return "", errors.New("exit")
    }

    fmt.Println(TargetName + " name is empty. Please type CTRL + C for cancel operation!")

    return getName(c, IsExistsStruct, targetName)
}

func GetPKType(isUuidAsPk bool) string {
    if isUuidAsPk {
        return "uuid.UUID"
    }

    return "int"
}

func GetConfigConverter(isUuidAsPk bool) string {
    if isUuidAsPk {
        return ".Uuid()"
    }

    return ".Int()"
}

func GetPkAsString(varExp string, isUuidAsPk bool) string {
    if isUuidAsPk {
        return varExp + ".String()"
    }

    return "strconv.Itoa(" + varExp + ")"
}

func GetStrToIdFuncName(isUuidAsPk bool) string {
    if isUuidAsPk {
        return "uuid.Parse"
    }

    return "strconv.Atoi"
}

func GetIdIsValidExp(isUuidAsPk bool) string {
    if isUuidAsPk {
        return "!= uuid.Nil"
    }

    return "> 0"
}

func GetIdNil(isUuidAsPk bool) string {
    if isUuidAsPk {
        return "uuid.Nil"
    }

    return "0"
}

func GetIdIsNotValidExp(isUuidAsPk bool) string {
    if isUuidAsPk {
        return "== uuid.Nil"
    }

    return "< 1"
}
