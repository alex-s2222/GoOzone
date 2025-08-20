package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Person struct{
	Name string
	Amount int
}

func main() {
	persons, err := parseFile("./data.txt")
	if err != nil{
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", persons)

	fmt.Println(sumAmountsFromFile("./data.txt"))
}


func parseFile(filename string)([]Person, error) {
	f, err := os.Open(filename)
	if err != nil{
		return nil, errors.Wrap(err, "Open File")
	}
	defer f.Close()
	
	return parseReader(f)
}


func parseReader(rd io.Reader)([]Person, error){
	data, err := io.ReadAll(rd)
	if err != nil{
		return nil, errors.Wrap(err, "reading file")
	}

	lines := strings.Split(string(data), "\n")

	// что бы не было аллокаций определяем копасити
	persons := make([]Person, 0, len(lines))

	for _, line := range lines {
		person, err := parseLine(line)
		if err != nil {
			var myError incorrectLineError
			if errors.As(err, &myError){
				return nil, errors.New("incorect line: " + myError.line)
			}
			return nil, errors.Wrap(err, "parsing person")
		}
		
		persons = append(persons, person)
	}

	return persons, nil
}

func sumAmountsFromFile(filename string) (map[string]int, error){
	f, err := os.Open(filename)
	if err!=nil{
		return nil, errors.Wrap(err, "Opening File")
	}
	defer f.Close()

	return sumAmountsFromReader(f)
}

func sumAmountsFromReader(r io.Reader) (map[string]int, error){
	buf := bufio.NewReader(r)
	aggregation := make(map[string]int)
	
	for{
		line, err := buf.ReadString('\n')
		if err != nil{
			if err == io.EOF{
				break
			}
			return nil, errors.Wrap(err, "reading line")
		}

		line = strings.TrimSuffix(line, "\n")

		person, err := parseLine(line)
		if err != nil{
			return nil, errors.Wrap(err, "parsing line")
		}
		fmt.Println(person.Name, person.Amount)
		aggregation[person.Name] += person.Amount
	}
	return aggregation, nil
}

var lineRe = regexp.MustCompile(`^Name:([^,]+), Amount:(-?\d+)$`)
var (
	// errIncorectLine = errors.New("the line in incorrect")
	errCannotParseAmount = errors.New("cannot parse amount")
)

type incorrectLineError struct{
	line string
}

func (e incorrectLineError) Error() string{
	return "incorect line"
}

func parseLine(line string) (Person, error) {
	matches := lineRe.FindStringSubmatch(line)
	if len(matches) < 3 {
		return Person{}, incorrectLineError{
			line: line,
		}
	}

	name := matches[1]
	amountStr := matches[2]

	amount, err := strconv.Atoi(amountStr)
	if err != nil{
		return Person{}, errCannotParseAmount 
	}

	return Person{
		Name: name,
		Amount: amount,
	}, nil
}