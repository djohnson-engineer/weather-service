package interfaces

// Defines the contract for characterizing temperatures
type ITemperatureCharacterizer interface {
	CharacterizeTemperature(temp int) string
}
