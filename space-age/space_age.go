package space

type Planet string

const earthsecs = 31557600

// Age is a function that calculates how old would a person be on a different planet
func Age(seconds float64, planet Planet) float64 {

	var orbitalYears float64

	switch planet {
	case "Earth":
		orbitalYears = 1
	case "Mercury":
		orbitalYears = 0.2408467
	case "Venus":
		orbitalYears = 0.61519726
	case "Mars":
		orbitalYears = 1.8808158
	case "Jupiter":
		orbitalYears = 11.862615
	case "Saturn":
		orbitalYears = 29.447498
	case "Uranus":
		orbitalYears = 84.016846
	case "Neptune":
		orbitalYears = 164.79132
	}

	earthYears := seconds / earthsecs
	planetYears := earthYears / orbitalYears

	return planetYears

}
