package translation

type DefaultTemperatureCategorizer struct{}

func (c *DefaultTemperatureCategorizer) CharacterizeTemperature(temp int) string {
	if temp <= 40 {
		return "cold"
	} else if temp >= 80 {
		return "hot"
	}
	return "moderate"
}
