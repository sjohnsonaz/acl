package access

import "acl/key"

type (
	AccessControlList struct {
		ID     AccessControlListID
		Levels map[key.UserID]AccessLevel
	}

	AccessControlListID string
)

func (acl *AccessControlList) initLevels() {
	if acl.Levels == nil {
		acl.Levels = make(map[key.UserID]AccessLevel)
	}
}

func (acl *AccessControlList) Grant(userID key.UserID, level AccessLevel) {
	switch level {
	case AccessLevelOwner:
		fallthrough
	case AccessLevelRead:
		fallthrough
	case AccessLevelWrite:
		acl.initLevels()
		acl.Levels[userID] = level
	default:
		acl.Revoke(userID)
	}
}

func (acl *AccessControlList) GrantOwner(userID key.UserID) {
	acl.initLevels()
	acl.Levels[userID] = AccessLevelOwner
}

func (acl *AccessControlList) GrantWrite(userID key.UserID) {
	acl.initLevels()
	acl.Levels[userID] = AccessLevelWrite
}

func (acl *AccessControlList) GrantRead(userID key.UserID) {
	acl.initLevels()
	acl.Levels[userID] = AccessLevelRead
}

func (acl *AccessControlList) Revoke(userID key.UserID) {
	delete(acl.Levels, userID)
}

func (acl *AccessControlList) Access(userID key.UserID) AccessLevel {
	return acl.Levels[userID]
}

func (acl *AccessControlList) OwnerAccess(userID key.UserID) bool {
	return acl.Levels[userID].OwnerAccess()
}

func (acl *AccessControlList) WriteAccess(userID key.UserID) bool {
	return acl.Levels[userID].WriteAccess()
}

func (acl *AccessControlList) ReadAccess(userID key.UserID) bool {
	return acl.Levels[userID].ReadAccess()
}
