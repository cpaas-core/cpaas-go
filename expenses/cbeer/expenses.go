package expenses

import "errors"

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
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(in Record) bool {
		return (in.Day >= p.From) && (in.Day <= p.To)
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(in Record) bool {
		return in.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var totalExpenses float64
	filteredRecords := Filter(in, ByDaysPeriod(p))
	for _, record := range filteredRecords {
		totalExpenses += record.Amount
	}
	return totalExpenses
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var totalExpenses float64
	categoryFilteredRecords := Filter(in, ByCategory(c))
	if len(categoryFilteredRecords) < 1 {
		return 0.0, errors.New("No matching categories")
	}
	filteredRecords := Filter(categoryFilteredRecords, ByDaysPeriod(p))
	for _, record := range filteredRecords {
		totalExpenses += record.Amount
	}
	return totalExpenses, nil
}
