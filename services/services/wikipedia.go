package services

import (
	"Helper_Bot/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Wiki(title string) string {
	url := fmt.Sprintf("https://az.wikipedia.org/w/api.php?action=query&format=json&prop=extracts&exintro&explaintext&titles=%s", strings.ToLower(title))
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

	var result models.WikiResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "Xəta baş verdi: JSON formatı düzgün deyil."
	}

	for pageID, page := range result.Query.Pages {
		if pageID == "-1" {
			return fmt.Sprintf("Haqqında məlumat tapılmadı: %s", title)
		}
		if page.Extract == "" {
			return fmt.Sprintf("Haqqında məlumat tapılmadı: %s", title)
		}
		return page.Extract
	}

	return fmt.Sprintf("Heç bir məlumat tapılmadı: %s", title)
}
