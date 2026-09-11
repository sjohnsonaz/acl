package access

type (
	OrganizationRoleAccess struct {
		Payment  AccessLevel
		Role     AccessLevel
		Schedule AccessLevel
	}

	OrganizationRole string
)

const (
	OrganizationRoleNone      OrganizationRole = ""
	OrganizationRoleManager   OrganizationRole = "manager"
	OrganizationRoleOwner     OrganizationRole = "owner"
	OrganizationRoleScheduler OrganizationRole = "scheduler"
)

var (
	None = OrganizationRoleAccess{
		Payment:  AccessLevelNone,
		Role:     AccessLevelNone,
		Schedule: AccessLevelNone,
	}
	Manager = OrganizationRoleAccess{
		Payment:  AccessLevelWrite,
		Role:     AccessLevelWrite,
		Schedule: AccessLevelWrite,
	}
	Owner = OrganizationRoleAccess{
		Payment:  AccessLevelOwner,
		Role:     AccessLevelOwner,
		Schedule: AccessLevelOwner,
	}
	Scheduler = OrganizationRoleAccess{
		Payment:  AccessLevelNone,
		Role:     AccessLevelNone,
		Schedule: AccessLevelWrite,
	}
)

func (or OrganizationRole) Access() OrganizationRoleAccess {
	switch or {
	case OrganizationRoleManager:
		return Manager
	case OrganizationRoleOwner:
		return Owner
	case OrganizationRoleScheduler:
		return Scheduler
	default:
		return None
	}
}
