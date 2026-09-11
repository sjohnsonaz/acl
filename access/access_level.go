package access

type AccessLevel string

const (
	AccessLevelNone  AccessLevel = ""
	AccessLevelRead  AccessLevel = "read"
	AccessLevelWrite AccessLevel = "write"
	AccessLevelOwner AccessLevel = "owner"
)

func (al AccessLevel) OwnerAccess() bool {
	switch al {
	case AccessLevelOwner:
		return true
	default:
		return false
	}
}

func (al AccessLevel) WriteAccess() bool {
	switch al {
	case AccessLevelOwner:
		fallthrough
	case AccessLevelWrite:
		return true
	default:
		return false
	}
}

func (al AccessLevel) ReadAccess() bool {
	switch al {
	case AccessLevelOwner:
		fallthrough
	case AccessLevelWrite:
		fallthrough
	case AccessLevelRead:
		return true
	default:
		return false
	}
}
