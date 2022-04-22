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
	var filtered_records []Record
	for _, r := range in {
		if predicate(r) {
			filtered_records = append(filtered_records, r)
		}
	}
	return filtered_records
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}

}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}

}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	Total := 0.0
	filter := Filter(in, ByDaysPeriod(p))
	for _, r := range filter {
		Total += r.Amount
	}
	return Total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	category_wise_total := 0.0
	err := errors.New("Category Not Found")
	filter_day := Filter(in, ByDaysPeriod(p))
	filter_category := Filter(in, ByCategory(c))
	if len(filter_category) < 1 {
		return 0, err
	} else if len(filter_day) < 1 {
		return 0, nil
	} else {
		category_wise_total += TotalByPeriod(filter_category, p)
	}
	return category_wise_total, nil
}
