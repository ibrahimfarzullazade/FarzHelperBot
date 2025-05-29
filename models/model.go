package models

type ExchangeRateResponse struct {
	Result string             `json:"result"`
	Rates  map[string]float64 `json:"rates"`
}
type WeatherResponse struct {
	Temp string `json:"temperature"`
	Wind string `json:"wind"`
	Desc string `json:"description"`
}
type WikiResponse struct {
	BatchComplete string `json:"batchcomplete"`
	Query         struct {
		Pages map[string]struct {
			PageID  int    `json:"pageid"`
			NS      int    `json:"ns"`
			Title   string `json:"title"`
			Extract string `json:"extract"`
		} `json:"pages"`
	} `json:"query"`
}
