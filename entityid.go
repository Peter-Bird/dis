package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// EntityID represents a unique identifier for an entity.
type EntityID struct {
	SiteID        uint16
	ApplicationID uint16
	EntityID      uint16
}

// InitializeEntityID initializes an EntityID with default values.
func InitializeEntityID() EntityID {
	return EntityID{
		SiteID:        1,
		ApplicationID: 1,
		EntityID:      1001,
	}
}

// CreateEntityIDTab creates the Fyne UI components for EntityID.
func CreateEntityIDTab(entity *EntityStatePDU) *fyne.Container {
	siteIDEntry := widget.NewEntry()
	siteIDEntry.SetText(strconv.Itoa(int(entity.EntityID.SiteID)))

	applicationIDEntry := widget.NewEntry()
	applicationIDEntry.SetText(strconv.Itoa(int(entity.EntityID.ApplicationID)))

	entityIDEntry := widget.NewEntry()
	entityIDEntry.SetText(strconv.Itoa(int(entity.EntityID.EntityID)))

	// Create a form with the EntityID fields
	entityIDForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Site ID", Widget: siteIDEntry},
			{Text: "Application ID", Widget: applicationIDEntry},
			{Text: "Entity ID", Widget: entityIDEntry},
		},
	}

	entityIDTab := container.NewVBox(
		entityIDForm,
	)

	return entityIDTab
}
