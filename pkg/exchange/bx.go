package exchange

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

type Bx struct{}

type QuoteKlinesParams struct {
	symbol    string
	interval  string
	limit     string
	startTime int64
	timestamp int64
}

func (bx Bx) Klines(symbol string, interval string) ([]byte, error) {
	now_ms := time.Now().UnixNano() / 1e6
	t0_ms := now_ms - (1000 * 60 * 60 * 24)

	queryString := url.Values{}

	p := QuoteKlinesParams{
		symbol:    symbol,
		interval:  interval,
		limit:     "100",
		startTime: t0_ms,
		timestamp: now_ms,
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

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Add("X-BX-APIKEY", "")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	return body, err
}
