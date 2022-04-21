package expenses

import (
	"errors"
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
	var recordResult []Record
	for _, record := range in {
		if predicate(record) {
			recordResult = append(recordResult, record)
		}
	}
	return recordResult
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		if (record.Day) <= p.To && (record.Day) >= p.From {
			return true
		}
		return false
	}

}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		if record.Category == c {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var totalByPeriod float64
	var recordResult []Record

	recordResult = Filter(in, ByDaysPeriod(p))
	for _, record := range recordResult {
		totalByPeriod = totalByPeriod + record.Amount
	}
	return totalByPeriod
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var total float64
	var recordResultByPeriod []Record
	var recordResultByCategory []Record

	// Error
	var ErrEmptyRecord = errors.New("No result for that category")

	recordResultByCategory = Filter(in, ByCategory(c))
	if len(recordResultByCategory) == 0 {
		return 0, ErrEmptyRecord
	}
	recordResultByPeriod = Filter(recordResultByCategory, ByDaysPeriod(p))
	for _, record := range recordResultByPeriod {
		total = total + record.Amount
	}
	return total, nil
}
