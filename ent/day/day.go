// Code generated by ent, DO NOT EDIT.

package day

const (
	// Label holds the string label denoting the day type in the database.
	Label = "day"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldMonth holds the string denoting the month field in the database.
	FieldMonth = "month"
	// FieldDay holds the string denoting the day field in the database.
	FieldDay = "day"
	// EdgeActivity holds the string denoting the activity edge name in mutations.
	EdgeActivity = "activity"
	// Table holds the table name of the day in the database.
	Table = "days"
	// ActivityTable is the table that holds the activity relation/edge.
	ActivityTable = "activities"
	// ActivityInverseTable is the table name for the Activity entity.
	// It exists in this package in order to avoid circular dependency with the "activity" package.
	ActivityInverseTable = "activities"
	// ActivityColumn is the table column denoting the activity relation/edge.
	ActivityColumn = "day_activity"
)

// Columns holds all SQL columns for day fields.
var Columns = []string{
	FieldID,
	FieldYear,
	FieldMonth,
	FieldDay,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}