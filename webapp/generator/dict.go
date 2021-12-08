package generator

// DONT EDIT THIS FILE IF YOU WILL RUN GENERATOR

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"math"
	"strconv"
)

var Babbler1 = NewBabbler(1)
var Babbler2 = NewBabbler(2)
var Babbler3 = NewBabbler(3)
var Babbler100 = NewBabbler(100)

func randate() time.Time {
	
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func randDuration() time.Duration {

	d := 24 * rand.Intn(364)
	h := strconv.Itoa(rand.Intn(23) + d)
	m := strconv.Itoa(rand.Intn(60))
	s := strconv.Itoa(rand.Intn(60))

	dur, _ := time.ParseDuration(h + "h" + m + "m" + s + "s")

	return dur
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func Float64n(precision int) float64 {

	output := math.Pow(10, float64(precision))
	rnd := rand.Float64()
	return float64(round(rnd * output)) / output
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Babbler struct {
	Count int
	Separator string
	Words []string
}

func NewBabbler(count int) (b Babbler) {
	b.Count = count

	if b.Count < 1 {
		b.Count = 1
	}

	b.Separator = " "
	b.Words = readAvailableDictionary()
	return
}

func (this Babbler) Babble() string {
	pieces := []string{}
	for i := 0; i < this.Count ; i++ {
		pieces = append(pieces, this.Words[rand.Int()%len(this.Words)])
	}

	return strings.Join(pieces, this.Separator)
}

func readAvailableDictionary() (words []string) {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		panic("Dictionary file for generate string examples not found. " +  err.Error())
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	words = strings.Split(string(bytes), "\n")
	return
}
