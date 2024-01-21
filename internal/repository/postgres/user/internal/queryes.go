package queryes

var (
	Create = "INSERT INTO users(email, password) VALUES($1, $2);"
)