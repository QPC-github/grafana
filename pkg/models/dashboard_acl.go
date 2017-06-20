package models

import (
	"errors"
	"time"
)

type PermissionType int

const (
	PERMISSION_NONE                = 0
	PERMISSION_VIEW PermissionType = 1 << iota
	PERMISSION_EDIT
	PERMISSION_ADMIN
)

func (p PermissionType) String() string {
	names := map[int]string{
		int(PERMISSION_NONE):  "None",
		int(PERMISSION_VIEW):  "View",
		int(PERMISSION_EDIT):  "Edit",
		int(PERMISSION_ADMIN): "Admin",
	}
	return names[int(p)]
}

// Typed errors
var (
	ErrDashboardAclInfoMissing           = errors.New("User id and user group id cannot both be empty for a dashboard permission.")
	ErrDashboardPermissionDashboardEmpty = errors.New("Dashboard Id must be greater than zero for a dashboard permission.")
)

// Dashboard ACL model
type DashboardAcl struct {
	Id          int64
	OrgId       int64
	DashboardId int64

	UserId      int64
	UserGroupId int64
	Permissions PermissionType

	Created time.Time
	Updated time.Time
}

type DashboardAclInfoDTO struct {
	Id          int64 `json:"id"`
	OrgId       int64 `json:"-"`
	DashboardId int64 `json:"dashboardId"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`

	UserId         int64          `json:"userId"`
	UserLogin      string         `json:"userLogin"`
	UserEmail      string         `json:"userEmail"`
	UserGroupId    int64          `json:"userGroupId"`
	UserGroup      string         `json:"userGroup"`
	Role           RoleType       `json:"role"`
	Permissions    PermissionType `json:"permissions"`
	PermissionName string         `json:"permissionName"`
}

//
// COMMANDS
//

type SetDashboardAclCommand struct {
	DashboardId int64          `json:"-"`
	OrgId       int64          `json:"-"`
	UserId      int64          `json:"userId"`
	UserGroupId int64          `json:"userGroupId"`
	Permissions PermissionType `json:"permissions" binding:"Required"`

	Result DashboardAcl `json:"-"`
}

type RemoveDashboardAclCommand struct {
	AclId int64
	OrgId int64
}

//
// QUERIES
//
type GetDashboardAclInfoListQuery struct {
	DashboardId int64
	Result      []*DashboardAclInfoDTO
}

// Returns dashboard acl list items and parent folder items
type GetInheritedDashboardAclQuery struct {
	DashboardId int64
	OrgId       int64
	Result      []*DashboardAcl
}
