package queryes

var (
	Create = "INSERT INTO users(email, password) VALUES($1, $2) RETURNING id;"
	GetByEmail = "SELECT id FROM users WHERE email = $1;"
)