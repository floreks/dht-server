// The MIT License
//
// Copyright (c) 2016 Sebastian Florek
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package sensor

import (
	"fmt"
	"strings"

	"github.com/d2r2/go-dht"
	"github.com/emicklei/go-restful/log"
)

const (
	GPIO_PIN        = 17
	RETRIES         = 5
	BOOST_PERF_FLAG = false
)

var (
	gpioPinOverride = GPIO_PIN
)

// TODO add doc
type DHT11Reader struct{}

// TODO add doc
func (DHT11Reader) ReadFromSensor() (*DHTResponse, error) {
	log.Printf("Reading from sensor %s on GPIO pin %d", dht.DHT11, gpioPinOverride)
	temp, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT11, gpioPinOverride, BOOST_PERF_FLAG,
		RETRIES)
	if err != nil {
		if strings.Contains(err.Error(), "C.dial_DHTxx_and_read") {
			return nil, fmt.Errorf("Could not read from sensor.")
		}

		return nil, err
	}

	return &DHTResponse{DHTTemperature{Temperature: temp}, DHTHumidity{Humidity: humidity}}, nil
}

// TODO add doc
func (d DHT11Reader) ReadTemperature() (*DHTTemperature, error) {
	response, err := d.ReadFromSensor()
	if err != nil {
		return nil, err
	}

	return &DHTTemperature{Temperature: response.Temperature}, nil
}

// TODO add doc
func (d DHT11Reader) ReadHumidity() (*DHTHumidity, error) {
	response, err := d.ReadFromSensor()
	if err != nil {
		return nil, err
	}

	return &DHTHumidity{Humidity: response.Humidity}, nil
}

// TODO add doc
func (d DHT11Reader) SetGPIO(pinNumber int) {
	log.Printf("Overriding GPIO pin number. Current: %d, New: %d", gpioPinOverride, pinNumber)
	gpioPinOverride = pinNumber
}

func (d DHT11Reader) ResetGPIO() {
	log.Printf("Resetting GPIO pin number to default. Current: %d, New: %d", gpioPinOverride, GPIO_PIN)
	gpioPinOverride = GPIO_PIN
}
