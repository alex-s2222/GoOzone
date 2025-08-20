package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseLine_ShoildFillPersonFields(t *testing.T){
	line := "Name:Иван, Amount:10"

	person, err := parseLine(line)
	assert.NoError(t, err)
	assert.Equal(t, Person{
		Name: "Иван", 
		Amount: 10,
	}, person)
}

func Test_ParseLine_ShouldFillPersonFilds_WhenAnountIsNegative(t *testing.T){
	line := "Name:Иван, Amount:-10"

	person, err := parseLine(line)

	assert.NoError(t,err)
	assert.Equal(t, Person{
		Name:"Иван",
		Amount: -10,
	}, person)
}

func Test_ParseLine_ShouldReturnError_WhenAmountIsNotNumber(t *testing.T){
	line := "Name:Иван, Amount:fejs21"
	_, err := parseLine(line)
	assert.Error(t, err)
}

func Test_ParseLine_ShouldReturnError_WhenAmountIsEmty(t *testing.T){
	line := "Name:Иван, Amount:"
	_, err := parseLine(line)
	assert.Error(t, err)
}

func Test_ParseLine_ShouldReturnError_WhenAmountIsTooBig(t *testing.T){
	line := "Name:Иван, Amount:111111111111111111111111111111"
	_, err := parseLine(line)
	assert.Error(t, err)
	assert.Equal(t, errCannotParseAmount, err)
}

func Test_ParseReader_WrongLineshouldGiveErr(t *testing.T){
	data := "Name:Сергей, Amount:49\n" +
			"Name:Петр, Amount:65"
	buf := bytes.NewBufferString(data)
	_, err := parseReader(buf)
	
	assert.Error(t, err)

}

func Test_ParseReader_CorrentLineShouldGivePersonsList(t *testing.T){
	data := "Name:Сергей, Amount:49\n" +
			"Name:Петр, Amount:65"
	buf := bytes.NewBufferString(data)

	person, err := parseReader(buf)

	assert.NoError(t, err)
	assert.Equal(t,
				[]Person{
				{
					Name:"Сергей",
					Amount:49,
				},
				{
					Name:"Петр",
					Amount:65,
				},
			}, person)

}