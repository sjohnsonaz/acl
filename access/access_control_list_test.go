package access

import (
	"testing"

	"acl/key"
)

func TestAccessControlList(t *testing.T) {
	userID := key.UserID("userID")

	t.Run("should not allow default access", func(t *testing.T) {
		acl := AccessControlList{}

		if acl.OwnerAccess(userID) {
			t.Error("should not allow owner access")
		}
		if acl.WriteAccess(userID) {
			t.Error("should not allow write access")
		}
		if acl.ReadAccess(userID) {
			t.Error("should not allow read access")
		}
	})

	t.Run("should not allow revoked access", func(t *testing.T) {
		acl := AccessControlList{}

		acl.GrantOwner(userID)
		acl.Revoke(userID)

		if acl.OwnerAccess(userID) {
			t.Error("should not allow owner access")
		}
		if acl.WriteAccess(userID) {
			t.Error("should not allow write access")
		}
		if acl.ReadAccess(userID) {
			t.Error("should not allow read access")
		}
	})

	t.Run("should allow owner access", func(t *testing.T) {
		acl := AccessControlList{}

		acl.GrantOwner(userID)

		if !acl.OwnerAccess(userID) {
			t.Error("should allow owner access")
		}
		if !acl.WriteAccess(userID) {
			t.Error("should allow write access")
		}
		if !acl.ReadAccess(userID) {
			t.Error("should allow read access")
		}
	})

	t.Run("should allow write access", func(t *testing.T) {
		acl := AccessControlList{}

		acl.GrantWrite(userID)

		if acl.OwnerAccess(userID) {
			t.Error("should not allow owner access")
		}
		if !acl.WriteAccess(userID) {
			t.Error("should allow write access")
		}
		if !acl.ReadAccess(userID) {
			t.Error("should allow read access")
		}
	})

	t.Run("should allow read access", func(t *testing.T) {
		acl := AccessControlList{}

		acl.GrantRead(userID)

		if acl.OwnerAccess(userID) {
			t.Error("should not allow owner access")
		}
		if acl.WriteAccess(userID) {
			t.Error("should not allow write access")
		}
		if !acl.ReadAccess(userID) {
			t.Error("should allow read access")
		}
	})
}
