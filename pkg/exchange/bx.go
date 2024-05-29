package exchange

import (
	"btaw/pkg/timeutil"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

type Bx struct{}

type QuoteKlinesParams struct {
	symbol    string
	interval  string
	limit     int64
	startTime int64
	endTime   int64
}

func (bx Bx) Klines(symbol string, tf timeutil.Interval, startTime int64, duration timeutil.Interval) ([]byte, error) {
	// now_ms := time.Now().UnixNano() / 1e6
	// t0_ms := now_ms - (1000 * 60 * 60 * 24)

	queryString := url.Values{}

	durationMs, err0 := duration.ToMs()

	if err0 != nil {
		return nil, nil
	}

	intervalMs, err1 := tf.ToMs()

	if err1 != nil {
		return nil, nil
	}

	p := QuoteKlinesParams{
		symbol:    symbol,
		interval:  tf.String(),
		limit:     durationMs / intervalMs,
		startTime: startTime,
		endTime:   startTime + durationMs,
	}

	// [!] improve naming
	o := reflect.ValueOf(p)

	for i := 0; i < o.NumField(); i++ {
		var string_value string

		// [!] improve naming
		f := o.Type().Field(i)

		switch f.Type.Name() {
		case "string":
			string_value = o.Field(i).String()
		case "int64":
			string_value = fmt.Sprintf("%d", o.Field(i).Int())
		}
		queryString.Set(f.Name, string_value)

	}

	queryStr := queryString.Encode()

	fullURL := fmt.Sprintf("https://open-api.bingx.com/openApi/swap/v3/quote/klines?%s", queryStr)

	req, err0 := http.NewRequest("GET", fullURL, nil)
	if err0 != nil {
		fmt.Println("Error creating request:", err0)
		return nil, err0
	}

	req.Header.Add("X-BX-APIKEY", "")

	client := &http.Client{}

	resp, err0 := client.Do(req)
	if err0 != nil {
		fmt.Println("Error sending request:", err0)
		return nil, err0
	}
	defer resp.Body.Close()
	body, err0 := io.ReadAll(resp.Body)
	if err0 != nil {
		fmt.Println("Error reading response body:", err0)
		return nil, err0
	}

	return body, err0
}
