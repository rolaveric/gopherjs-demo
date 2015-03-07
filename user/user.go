package user

// Interface for a database result row
type DBRow interface {
	GetInt(colnum int) int
	GetString(colnum int) string
}

// Interface for a database result
type DBResult interface {
	NextRow() DBRow
	RowCount() int
}

// Interface for an object which can be used to make database queries
type DB interface {
	Query(query string, params ...interface{}) DBResult
}

// Private package variable for the registered DB interface
var db DB

// Method for registering a DB interface
func RegisterDB(newDb DB) {
	db = newDb
}

// User type
type User struct {
	Name string
	ID   int
}

// Save method for the User type
func Save(u *User) {
	db.Query("UPDATE User SET name = ? WHERE id = ?", u.Name, u.ID)
}

// Function for creating a new User
func New(name string) *User {
	db.Query("INSERT INTO User (name) VALUES (?)", name)
	id := db.Query("SELECT @@IDENTITY").NextRow().GetInt(0)
	return &User{name, id}
}

// Function for getting a single User
func Get(id int) *User {
	result := db.Query("SELECT name FROM User WHERE id = ?", id)
	if result.RowCount() == 0 {
		return nil
	}
	name := result.NextRow().GetString(0)
	return &User{name, id}
}

// Function for getting all users
func All() []*User {
	result := db.Query("SELECT name, id FROM User")
	users := make([]*User, result.RowCount())
	for x, c := 0, result.RowCount(); x < c; x++ {
		row := result.NextRow()
		users[x] = &User{row.GetString(0), row.GetInt(1)}
	}
	return users
}
