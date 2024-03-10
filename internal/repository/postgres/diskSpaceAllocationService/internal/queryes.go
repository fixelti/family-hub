package queryes

var (
	GetUserServices = "SELECT * FROM disk_space_allocation_service WHERE user_id = $1;"
)