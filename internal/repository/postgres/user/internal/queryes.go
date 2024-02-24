package queryes

var (
	Create = "INSERT INTO users(email, password) VALUES($1, $2) RETURNING *;"
	GetByEmail = "SELECT id FROM users WHERE email = $1;"
	GetUserIDAndPasswordByEmail = "SELECT id, password FROM users WHERE email = $1;"
)