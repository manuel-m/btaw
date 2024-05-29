package exchange

type Exchange interface {
	Klines(symbol string, interval string) ([]byte, error)
}
