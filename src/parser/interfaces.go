package parser

type ArgumentsParser interface {
	Take2Numbers(args []string) (float64, float64, error)
}
