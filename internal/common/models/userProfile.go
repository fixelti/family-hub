package models

type UserProfile struct {
	// Name string `json:"name"`
	Email string `json:"email"`
	DiskSpaceAllocationService []DiskSpaceAllocationService 
}