package placeholder

import "platform/logging"

type WeatherHandler struct {
	logging.Logger
}

func (h WeatherHandler) GetWeather() string {
	return "Sunny"
}
