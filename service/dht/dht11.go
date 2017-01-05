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

package dht

import (
	"net/http"

	"github.com/emicklei/go-restful"
	sensor "github.com/floreks/dht-server/sensor/dht"
	"github.com/floreks/dht-server/service"
	"strconv"
)

// TODO add doc
type DHT11Service struct {
	reader sensor.DHTReader
}

// TODO add doc
func (d DHT11Service) Handler() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/dht11").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(d.readFromSensor).
		Doc("Reads temperature and humidity from DHT11 sensor").
		Writes(sensor.DHTResponse{}))

	ws.Route(ws.POST("/").To(d.setGPIO).
		Doc("Sets GPIO pin number that should be used to read sensor data from"))

	ws.Route(ws.GET("/temperature").To(d.readTemperature).
		Doc("Returns only temperature read from DHT11 sensor").
		Writes(sensor.DHTTemperature{}))

	ws.Route(ws.GET("/humidity").To(d.readHumidity).
		Doc("Returns only humidity read from DHT11 sensor").
		Writes(sensor.DHTHumidity{}))

	ws.Route(ws.GET("/reset").To(d.resetGPIO).
		Doc("Resets GPIO pin number to default value"))

	return ws
}

// TODO add doc
func (d DHT11Service) readFromSensor(request *restful.Request, response *restful.Response) {
	result, err := d.reader.ReadFromSensor()
	if err != nil {
		service.HandleInternalServerError(response, err)
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// TODO add doc
func (d DHT11Service) readTemperature(request *restful.Request, response *restful.Response) {
	result, err := d.reader.ReadTemperature()
	if err != nil {
		service.HandleInternalServerError(response, err)
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// TODO add doc
func (d DHT11Service) readHumidity(request *restful.Request, response *restful.Response) {
	result, err := d.reader.ReadHumidity()
	if err != nil {
		service.HandleInternalServerError(response, err)
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// TODO add doc
func (d DHT11Service) setGPIO(request *restful.Request, response *restful.Response) {
	gpioPin, err := strconv.Atoi(request.QueryParameter("gpio"))
	if err != nil {
		service.HandleInternalServerError(response, err)
		return
	}

	d.reader.SetGPIO(gpioPin)
	response.WriteHeader(http.StatusAccepted)
}

// TODO add doc
func (d DHT11Service) resetGPIO(request *restful.Request, response *restful.Response) {
	d.reader.ResetGPIO()
	response.WriteHeader(http.StatusAccepted)
}

// TODO add doc
func NewDHT11Service() DHT11Service {
	return DHT11Service{reader: sensor.DHT11Reader{}}
}
