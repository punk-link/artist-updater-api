package controllers

import "main/models/organizations"

func getCurrentManager() (organizations.Manager, error) {
	return organizations.Manager{
		Id:             1,
		Name:           "Kirill Taran",
		OrganizationId: 1,
	}, nil
}
