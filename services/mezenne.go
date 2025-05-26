package services

import (
	"Helper_Bot/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Mezenne(currency string, currency2 string) string {
	url := fmt.Sprintf("https://open.er-api.com/v6/latest/%s", strings.ToUpper(currency))
	resp, err := http.Get(url)
	if err != nil {
		return "Xəta baş verdi: API-yə qoşulmaq mümkün olmadı."
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("Xəta baş verdi: API status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Xəta baş verdi: cavab oxuna bilmədi."
	}

	var result models.ExchangeRateResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "Xəta baş verdi: JSON formatı düzgün deyil."
	}

	rate, ok := result.Rates[strings.ToUpper(currency2)]
	if !ok {
		return fmt.Sprintf("Valyuta tapılmadı: %s", currency2)
	}

	return fmt.Sprintf("1 %s = %.2f %s", strings.ToUpper(currency), rate, strings.ToUpper(currency2))
}
