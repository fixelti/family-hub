package models

type DiskSpaceAllocationService struct {
	ID       uint          `json:"-"`
	UserID   uint          `json:"user_id"`
	Name     string        `json:"name"`
	DiskSize uint          `json:"disk_size"`
	Status   ServiceStatus `json:"status"`
}
