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

package ds18b20

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/floreks/dht-server/sensor/ds18b20"
	"github.com/floreks/dht-server/service"
)

// TODO add doc
type DS18B20Service struct {
	reader sensor.DS18B20Reader
}

// TODO add doc
func (d DS18B20Service) Handler() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/ds18b20").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(d.readFromSensor).
		Doc("Reads temperature from DS18B20 sensor").
		Writes(sensor.DS18B20Response{}))

	return ws
}

// TODO add doc
func (d DS18B20Service) readFromSensor(request *restful.Request, response *restful.Response) {
	result, err := d.reader.ReadFromSensor()
	if err != nil {
		service.HandleInternalServerError(response, err)
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// TODO add doc
func NewDS18B20Service() DS18B20Service {
	return DS18B20Service{reader: sensor.DS18B20Reader{}}
}
