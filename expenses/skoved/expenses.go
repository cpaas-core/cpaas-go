package expenses

import "fmt"

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

var errorMsg string = "The record slice provided had no records in the provided " +
	"category %s"

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var filtered []Record
	for _, record := range in {
		if predicate(record) {
			filtered = append(filtered, record)
		}
	}
	return filtered
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		return record.Day >= p.From && record.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		return record.Category == c
	}
}

// TotalExpenses returns the sum of Amounts from each record in the list in
func TotalExpenses(in []Record) float64 {
	total := 0.0
	for _, record := range in {
		total += record.Amount
	}
	return total
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	return TotalExpenses(Filter(in, ByDaysPeriod(p)))
}

// CombinePredicatesAnd returns a predicate function that returns true when all
// provided predicate function return true
/*func CombinePredicatesAnd(predicates ...func(Record) bool) func(Record) bool {
	return func(record Record) bool {
		for _, predicate := range predicates {
			if !predicate(record) {
				return false
			}
		}
		return true
	}
}*/

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	filtered := Filter(in, ByCategory(c))
	if len(filtered) == 0 {
		return 0.0, fmt.Errorf(errorMsg, c)
	}
	return TotalExpenses(Filter(filtered, ByDaysPeriod(p))), nil
}
