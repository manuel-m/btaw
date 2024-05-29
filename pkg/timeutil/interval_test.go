package timeutil

import (
	"testing"
)

func TestTimeutil(t *testing.T) {

	t.Run("interval 1m", func(t *testing.T) {
		const wanted int64 = 1000000
		interval := Interval("1m")
		got, err := interval.ToMs()

		if err != nil {
			t.Errorf("%s|%v\n", t.Name(), err)
		}

		if got != wanted {
			t.Errorf("%s|%d != %d \n", t.Name(), got, wanted)
		}
	})

	t.Run("unknown interval", func(t *testing.T) {
		const wanted int64 = 0
		interval := Interval("!!UNKNOWN_INTERVAL!!")
		got, err := interval.ToMs()

		if err == nil {
			t.Errorf("%s|error expected\n", t.Name())
		}

		if got != wanted {
			t.Errorf("%s|%d != %d \n", t.Name(), got, wanted)
		}
	})

	t.Run("all intervals", func(t *testing.T) {

		intervals := []Interval{"1m", "3m", "5m", "15m", "30m", "1h", "2h", "4h", "8h", "12h", "1d", "3d", "1w"}

		for _, interval := range intervals {
			_, err := interval.ToMs()
			if err != nil {
				t.Errorf("%s|%s: %v\n", t.Name(), interval, err)
			}
		}
	})

}
