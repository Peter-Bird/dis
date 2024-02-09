package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// PDUHeader represents the header for a DIS PDU.
type PDUHeader struct {
	ProtocolVersion uint8
	ExerciseID      uint8
	PDUType         uint8
	ProtocolFamily  uint8
	Timestamp       uint32
	Length          uint16
}

// InitializePDUHeader initializes an example PDUHeader.
func InitializePDUHeader() PDUHeader {
	return PDUHeader{
		ProtocolVersion: 7,
		ExerciseID:      1,
		PDUType:         1,
		ProtocolFamily:  2,
		Timestamp:       12345678,
		Length:          144, // Example length
	}
}

// CreatePDUHeaderTab creates the PDUHeader tab for the GUI using a form.
func CreatePDUHeaderTab(entity *EntityStatePDU) *fyne.Container {
	protocolVersionEntry := widget.NewEntry()
	protocolVersionEntry.SetText(strconv.Itoa(int(entity.Header.ProtocolVersion)))

	exerciseIDEntry := widget.NewEntry()
	exerciseIDEntry.SetText(strconv.Itoa(int(entity.Header.ExerciseID)))

	// PDU Type Dropdown
	pduTypeOptions := []string{
		"0 Other",
		"1 Entity State",
		"10 Repair Response",
		"11 Create Entity",
		"12 Remove Entity",
		"129 Announce Object",
		"13 Start/Resume",
		"130 Delete Object",
		"131 Describe Application",
		"132 Describe Event",
		"133 Describe Object",
		"134 Request Event",
		"135 Request Object",
		"14 Stop/Freeze",
		"140 Time Space Position Indicator - FI",
		"141 Appearance-FI",
		"142 Articulated Parts - FI",
		"143 Fire - FI",
		"144 Detonation - FI",
		"15 Acknowledge",
		"150 Point Object State",
		"151 Linear Object State",
		"152 Areal Object State",
		"153 Environment",
		"155 Transfer Control Request",
		"156 Transfer Control",
		"157 Transfer Control Acknowledge",
		"16 Action Request",
		"160 Intercom Control",
		"161 Intercom Signal",
		"17 Action Response",
		"170 Aggregate",
		"18 Data Query",
		"19 Set Data",
		"2 Fire",
		"20 Data",
		"21 Event Report",
		"22 Comment",
		"23 Electromagnetic Emission",
		"24 Designator",
		"25 Transmitter",
		"26 Signal",
		"27 Receiver",
		"3 Detonation",
		"4 Collision",
		"5 Service Request",
		"6 Resupply Offer",
		"7 Resupply Received",
		"8 Resupply Cancel",
		"9 Repair Complete",
	}
	pduTypeDropdown := widget.NewSelect(pduTypeOptions, nil)
	pduTypeDropdown.SetSelected("1 Entity State")

	// Protocol Family Dropdown
	protocolFamilyOptions := []string{
		"0 Other",
		"1 Entity Information/Interaction",
		"129 Experimental - CGF",
		"130 Experimental - Entity Interaction/Information - Field Instrumentation",
		"131 Experimental - Warfare Field Instrumentation",
		"132 Experimental - Environment Object Information/Interaction",
		"133 Experimental - Entity Management",
		"2 Warfare",
		"3 Logistics",
		"4 Radio Communication",
		"5 Simulation Management",
		"6 Distributed Emission Regeneration",
	}
	protocolFamilyDropdown := widget.NewSelect(protocolFamilyOptions, nil)
	protocolFamilyDropdown.SetSelected("2 Warfare")

	timestampEntry := widget.NewEntry()
	timestampEntry.SetText(strconv.Itoa(int(entity.Header.Timestamp)))

	lengthEntry := widget.NewEntry()
	lengthEntry.SetText(strconv.Itoa(int(entity.Header.Length)))

	pduHeaderForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Protocol Version", Widget: protocolVersionEntry},
			{Text: "Exercise ID", Widget: exerciseIDEntry},
			{Text: "PDU Type", Widget: pduTypeDropdown},
			{Text: "Protocol Family", Widget: protocolFamilyDropdown},
			{Text: "Timestamp", Widget: timestampEntry},
			{Text: "Length", Widget: lengthEntry},
		},
	}

	pduHeaderTab := container.NewVBox(
		pduHeaderForm,
	)

	return pduHeaderTab
}
