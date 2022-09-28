package stuff 

import {
	"errors"
	"strconv"
	"time"
}

var Name string = "Omotoye"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}

type Date struct {
	day int 
	month int 
	year int 
}

func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("incorrect day value")
	}
	d.day = day
	return nil 
}

func (d *Date) SetMonth(m int) error {
	if (m < 1) || (m > 31) {
		return errors.New("incorrect day value")
	}
	d.month = month 
	return nil 
}


func (d *Date) SetYear(y int) error {
	if (y < 1875) || (y > time.Now().Year()) {
		return errors.New("incorrect day value")
	}
	d.year = y 
	return nil 
}

// creating getter functions
func (d *Date) Day() int {
	return d.day
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Year() int {
	return d.year
}

