package space

const EarthSeconds = 31557600

type Planet string

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case
		"Mercury":
		return seconds / (.2408467 * EarthSeconds)
	case "Venus":
		return seconds / (.61519726 * EarthSeconds)
	case "Earth":
		return seconds / (1 * EarthSeconds)
	case "Mars":
		return seconds / (1.8808158 * EarthSeconds)
	case "Jupiter":
		return seconds / (11.862615 * EarthSeconds)
	case "Saturn":
		return seconds / (29.447498 * EarthSeconds)
	case "Uranus":
		return seconds / (84.016846 * EarthSeconds)
	case "Neptune":
		return seconds / (164.79132 * EarthSeconds)
	}
	return seconds
}
