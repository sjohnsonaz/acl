package key

import "uuid"

type OrganizationID string

func NewOrganizationID() OrganizationID {
	return OrganizationID(uuid.New().String())
}

func (o OrganizationID) String() string {
	return string(o)
}
