package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

// Vector represents a 3D vector.
type Vector struct {
	X float32
	Y float32
	Z float32
}

// EntityStatePDU represents the DIS Entity State PDU.
type EntityStatePDU struct {
	Header                     PDUHeader
	EntityID                   EntityID
	ForceID                    uint8
	NumberOfArticulationParams uint8
	EntityType                 EntityType
	EntityLinearVelocity       Vector
	EntityLocation             Vector
	EntityOrientation          Vector
	EntityAppearance           uint32
	DeadReckoningParams        [44]byte // Placeholder for simplicity
	EntityMarking              [32]byte // Placeholder for simplicity
	Capabilities               uint32
}

func main() {

	// Load configuration from file
	config := Config{}
	data, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Failed to read config file:", err)
		return
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Failed to parse config file:", err)
		return
	}

	os.Setenv("FYNE_THEME", "dark")

	myApp := app.New()
	//myApp.Settings().SetTheme(theme.DarkTheme())

	myWindow := myApp.NewWindow("DIS Controller")

	// Example instantiation of an EntityStatePDU
	entity := EntityStatePDU{
		Header:               InitializePDUHeader(),
		EntityID:             InitializeEntityID(),
		ForceID:              1,
		EntityType:           InitializeEntityType(),
		EntityLinearVelocity: Vector{X: 0.0, Y: 0.0, Z: 0.0},
		EntityLocation:       Vector{X: 100.0, Y: 200.0, Z: 300.0},
		EntityOrientation:    Vector{X: 0.0, Y: 0.0, Z: 0.0},
		EntityAppearance:     0,
		Capabilities:         0,
	}

	pduHeaderTab := CreatePDUHeaderTab(&entity)   /* PDUHeader Tab */
	entityIDTab := CreateEntityIDTab(&entity)     /* EntityID Tab  */
	entityTypeTab := CreateEntityTypeTab(&entity) /* EntityType Tab */

	// Vector and Other Fields Tab
	linearVelocityXEntry := widget.NewEntry()
	linearVelocityXEntry.SetText(strconv.FormatFloat(float64(entity.EntityLinearVelocity.X), 'f', 2, 32))

	linearVelocityYEntry := widget.NewEntry()
	linearVelocityYEntry.SetText(strconv.FormatFloat(float64(entity.EntityLinearVelocity.Y), 'f', 2, 32))

	linearVelocityZEntry := widget.NewEntry()
	linearVelocityZEntry.SetText(strconv.FormatFloat(float64(entity.EntityLinearVelocity.Z), 'f', 2, 32))

	locationXEntry := widget.NewEntry()
	locationXEntry.SetText(strconv.FormatFloat(float64(entity.EntityLocation.X), 'f', 2, 32))

	locationYEntry := widget.NewEntry()
	locationYEntry.SetText(strconv.FormatFloat(float64(entity.EntityLocation.Y), 'f', 2, 32))

	locationZEntry := widget.NewEntry()
	locationZEntry.SetText(strconv.FormatFloat(float64(entity.EntityLocation.Z), 'f', 2, 32))

	orientationXEntry := widget.NewEntry()
	orientationXEntry.SetText(strconv.FormatFloat(float64(entity.EntityOrientation.X), 'f', 2, 32))

	orientationYEntry := widget.NewEntry()
	orientationYEntry.SetText(strconv.FormatFloat(float64(entity.EntityOrientation.Y), 'f', 2, 32))

	orientationZEntry := widget.NewEntry()
	orientationZEntry.SetText(strconv.FormatFloat(float64(entity.EntityOrientation.Z), 'f', 2, 32))

	// Vector Fields Tab
	vectorTab := container.NewVBox(
		widget.NewLabel("EntityLinearVelocity"),
		linearVelocityXEntry, linearVelocityYEntry, linearVelocityZEntry,
		widget.NewLabel("EntityLocation"),
		locationXEntry, locationYEntry, locationZEntry,
		widget.NewLabel("EntityOrientation"),
		orientationXEntry, orientationYEntry, orientationZEntry,
	)

	forceIDEntry := widget.NewEntry()
	forceIDEntry.SetText(strconv.Itoa(int(entity.ForceID)))

	numArticulationParamsEntry := widget.NewEntry()
	numArticulationParamsEntry.SetText(strconv.Itoa(int(entity.NumberOfArticulationParams)))

	entityAppearanceEntry := widget.NewEntry()
	entityAppearanceEntry.SetText(strconv.Itoa(int(entity.EntityAppearance)))

	capabilitiesEntry := widget.NewEntry()
	capabilitiesEntry.SetText(strconv.Itoa(int(entity.Capabilities)))

	// Other Fields Tab
	otherFieldsTab := container.NewVBox(
		forceIDEntry, numArticulationParamsEntry, entityAppearanceEntry, capabilitiesEntry,
	)

	// Button to update EntityStatePDU based on input fields
	updateButton := widget.NewButton("Send", func() {
		// ... [Update logic for each field here] ...
		// Example:
		// protocolVersion, _ := strconv.Atoi(pduHeader.protocolVersionEntry.Text)
		// entity.Header.ProtocolVersion = uint8(protocolVersion)
		// ... [Continue for other fields] ...

		// Serialize the Entity PDU
		buf := new(bytes.Buffer)
		err = binary.Write(buf, binary.LittleEndian, entity)
		if err != nil {
			fmt.Println("Failed to serialize Entity PDU:", err)
			return
		}

		// Send the serialized PDU over UDP
		conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", config.IP, config.Port))
		if err != nil {
			fmt.Println("Failed to create UDP connection:", err)
			return
		}
		defer conn.Close()

		_, err = conn.Write(buf.Bytes())
		if err != nil {
			fmt.Println("Failed to send PDU:", err)
			return
		}

		fmt.Println("Entity PDU sent successfully!")

		// Print updated EntityStatePDU (you can replace this with other logic)
		fmt.Println(entity)
	})

	tabs := container.NewAppTabs(
		container.NewTabItem("Header", pduHeaderTab),
		container.NewTabItem("Entity ID", entityIDTab),
		container.NewTabItem("Entity Type", entityTypeTab),
		container.NewTabItem("Vectors", vectorTab),
		container.NewTabItem("Other Fields", otherFieldsTab),
	)

	content := container.NewVBox(
		tabs,
		updateButton,
	)

	myWindow.Resize(fyne.NewSize(500, 420))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
