package main

import (
	"encoding/json"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// EntityType represents the type of an entity.
type EntityType struct {
	Kind        uint8
	Domain      uint8
	Country     uint16
	Category    uint8
	Subcategory uint8
	Specific    uint8
	Extra       uint8
}

type Field struct {
	Name    string `json:"name"`
	Default int    `json:"default"`
}

type EtConfig struct {
	Fields []Field `json:"fields"`
}

// LoadConfig loads the entity configuration from a JSON file.
func LoadConfig(filename string) (EtConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return EtConfig{}, err
	}
	defer file.Close()

	var config EtConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}

// CreateEntityTypeTab creates a Fyne form for EntityType fields.
func fCreateEntityTypeTab(config EtConfig, entity *EntityType) *fyne.Container {
	form := &widget.Form{}
	for _, field := range config.Fields {
		entry := widget.NewEntry()
		entry.SetText(strconv.Itoa(field.Default))
		form.Append(field.Name, entry)
	}
	return container.NewVBox(form)
}

// InitializeEntityType initializes an EntityType with default values.
func InitializeEntityType() EntityType {
	return EntityType{
		Kind:        1,
		Domain:      2,
		Country:     225, // Example for USA
		Category:    1,
		Subcategory: 1,
		Specific:    1,
		Extra:       0,
	}
}

// CreateEntityTypeTab creates a Fyne form for EntityType fields.
func CreateEntityTypeTab(entity *EntityStatePDU) *fyne.Container {
	kindEntry := widget.NewEntry()
	kindEntry.SetText(strconv.Itoa(int(entity.EntityType.Kind)))

	domainEntry := widget.NewEntry()
	domainEntry.SetText(strconv.Itoa(int(entity.EntityType.Domain)))

	countryEntry := widget.NewEntry()
	countryEntry.SetText(strconv.Itoa(int(entity.EntityType.Country)))

	categoryEntry := widget.NewEntry()
	categoryEntry.SetText(strconv.Itoa(int(entity.EntityType.Category)))

	subcategoryEntry := widget.NewEntry()
	subcategoryEntry.SetText(strconv.Itoa(int(entity.EntityType.Subcategory)))

	specificEntry := widget.NewEntry()
	specificEntry.SetText(strconv.Itoa(int(entity.EntityType.Specific)))

	extraEntry := widget.NewEntry()
	extraEntry.SetText(strconv.Itoa(int(entity.EntityType.Extra)))

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Kind", Widget: kindEntry},
			{Text: "Domain", Widget: domainEntry},
			{Text: "Country", Widget: countryEntry},
			{Text: "Category", Widget: categoryEntry},
			{Text: "Subcategory", Widget: subcategoryEntry},
			{Text: "Specific", Widget: specificEntry},
			{Text: "Extra", Widget: extraEntry},
		},
	}

	return container.NewVBox(form)
}
