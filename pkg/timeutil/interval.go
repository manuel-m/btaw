package timeutil

import "fmt"

type Interval string

func (interval Interval) String() string {
	return string(interval)
}

type Converter interface {
	ToMs() (int64, error)
}

func (interval Interval) ToMs() (int64, error) {

	switch interval {

	case "1m":
		return 1000000, nil
	case "3m":
		return 1000000 * 3, nil
	case "5m":
		return 1000000 * 5, nil
	case "15m":
		return 1000000 * 15, nil
	case "30m":
		return 1000000 * 30, nil
	case "1h":
		return 1000000 * 60, nil
	case "2h":
		return 1000000 * 60 * 2, nil
	case "4h":
		return 1000000 * 60 * 4, nil
	case "8h":
		return 1000000 * 60 * 8, nil
	case "12h":
		return 1000000 * 60 * 12, nil
	case "1d":
		return 1000000 * 60 * 24, nil
	case "3d":
		return 1000000 * 60 * 24 * 3, nil
	case "1w":
		return 1000000 * 60 * 24 * 7, nil
	default:
		return 0, fmt.Errorf("unknown unit: %s", interval)

	}

}
