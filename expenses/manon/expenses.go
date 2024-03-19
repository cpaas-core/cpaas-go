package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	filteredRecords := []Record{}
	for _, record := range in {
		if predicate(record) {
			filteredRecords = append(filteredRecords, record)
		}
	}
	return filteredRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return p.From <= r.Day && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var total float64
	byDaysPeriod := ByDaysPeriod(p)
	records := Filter(in, byDaysPeriod)
	for _, record := range records {
		total += record.Amount
	}
	return total
}

type UnknownCategoryError struct {
	category string
}

func (err *UnknownCategoryError) Error() string {
	return fmt.Sprintf("unkown category %s", err.category)
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	byCategory := ByCategory(c)
	recordsFilteredByCategory := Filter(in, byCategory)
	if len(recordsFilteredByCategory) == 0 {
		return 0, &UnknownCategoryError{category: c}
	}
	return TotalByPeriod(recordsFilteredByCategory, p), nil
}
