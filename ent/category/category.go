// Code generated by ent, DO NOT EDIT.

package category

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCategoryId holds the string denoting the categoryid field in the database.
	FieldCategoryId = "category_id"
	// FieldCategoryName holds the string denoting the categoryname field in the database.
	FieldCategoryName = "category_name"
	// FieldCategoryUrl holds the string denoting the categoryurl field in the database.
	FieldCategoryUrl = "category_url"
	// Table holds the table name of the category in the database.
	Table = "categories"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldCategoryId,
	FieldCategoryName,
	FieldCategoryUrl,
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

var (
	// DefaultCategoryId holds the default value on creation for the "categoryId" field.
	DefaultCategoryId string
	// DefaultCategoryName holds the default value on creation for the "categoryName" field.
	DefaultCategoryName string
	// DefaultCategoryUrl holds the default value on creation for the "categoryUrl" field.
	DefaultCategoryUrl string
)