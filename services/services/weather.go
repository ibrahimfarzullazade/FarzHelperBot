package services

import (
	"Helper_Bot/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func translateWeatherCity(city string) string {
	translations := map[string]string{
		"Bakı": "Baku",
	}
	if val, ok := translations[city]; ok {
		return val
	}
	return city
}
func translateWeatherDesc(desc string) string {
	translations := map[string]string{
		"Sunny":         "Günəşli",
		"Cloudy":        "Buludlu",
		"Partly cloudy": "Qismən buludlu",
		"Rain":          "Yağışlı",
		"Thunderstorm":  "Şimşəkli",
		"Snow":          "Qarlı",
		"Clear":         "Açıq",
		"Fog":           "Dumanlı",
		"Windy":         "Küləkli",
	}

	if val, ok := translations[desc]; ok {
		return val
	}
	return desc // tapılmasa ingiliscə qalır
}

func Weather(city string) string {
	city = translateWeatherCity(city)
	url := fmt.Sprintf("https://goweather.herokuapp.com/weather/%s", strings.ToLower(city))
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

	var result models.WeatherResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "Xəta baş verdi: JSON formatı düzgün deyil."
	}

	if result.Temp == "" && result.Wind == "" && result.Desc == "" {
		return fmt.Sprintf("Hava məlumatı tapılmadı: %s", city)
	}

	return fmt.Sprintf(
		"%s üçün hava:\n🌡 Temperatur: %s\n💨 Külək: %s\n🌤 Təsvir: %s",
		strings.Title(city),
		translateWeatherDesc(result.Temp),
		translateWeatherDesc(result.Wind),
		translateWeatherDesc(result.Desc),
	)
}
