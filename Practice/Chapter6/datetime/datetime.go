package datetime

import (
	"errors"
	"unicode/utf8"
)

type Date struct {
	year  int
	month int
	day   int
}

type Event struct {
	title string
	Date
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("string limit exceeded")
	}
	e.title = title
	return nil
}

func (e *Event) Title() string {
	return e.title
}

func (d *Date) AddYear(year int) error {
	if year < 1 {
		return errors.New("invalid Year")
	}
	d.year = year
	return nil
}

func (d *Date) AddMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month given")
	}
	d.month = month
	return nil
}

func (d *Date) AddDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day given")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}
