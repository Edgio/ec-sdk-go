package enums

type Platform int

const (
	HttpLarge Platform = 3
	HttpSmall          = 8
	Adn                = 14
)

func (p Platform) String() string {
	switch p {
	case HttpLarge:
		return "httplarge"
	case HttpSmall:
		return "httpsmall"
	case Adn:
		return "adn"
	}
	return "unknown"
}
