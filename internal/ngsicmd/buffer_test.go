/*
MIT License

Copyright (c) 2020-2021 Kazuhito Suda

This file is part of NGSI Go

https://github.com/lets-fiware/ngsi-go

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package ngsicmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufferOpen(t *testing.T) {
	jsonBuf := jsonBuffer{}
	jsonBuf.bufferOpen(os.Stdout, false, false)

	if !assert.Equal(t, "[", jsonBuf.delimiter) {
		t.FailNow()
	}
}

func TestBufferWrite1(t *testing.T) {
	buffer := &bytes.Buffer{}

	jsonBuf := jsonBuffer{}
	jsonBuf.bufferOpen(buffer, false, false)

	jsonBuf.bufferWrite([]byte("[abc]"))

	if !assert.Equal(t, []byte("abc"), jsonBuf.buf) {
		t.FailNow()
	}
}

func TestBufferWrite2(t *testing.T) {
	buf := &bytes.Buffer{}

	jsonBuf := jsonBuffer{}
	jsonBuf.bufferOpen(buf, false, false)

	jsonBuf.bufferWrite([]byte("[abc]"))
	jsonBuf.bufferWrite([]byte("[xyz]"))

	if assert.Equal(t, []byte("xyz"), jsonBuf.buf) {
		jsonBuf.bufferClose()
		actual := buf.String()
		expected := "[abc,xyz]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBufferWrite3(t *testing.T) {
	buf := &bytes.Buffer{}

	jsonBuf := jsonBuffer{}
	jsonBuf.bufferOpen(buf, false, false)

	jsonBuf.bufferWrite([]byte("[abc]"))
	jsonBuf.bufferWrite([]byte("[xyz]"))
	jsonBuf.bufferWrite(nil)

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		jsonBuf.bufferClose()
		actual := buf.String()
		expected := "[abc,xyz]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBufferGeoJSON1(t *testing.T) {
	buf := &bytes.Buffer{}

	geoJSON := true
	pretty := false

	jsonBuf := jsonBuffer{}
	jsonBuf.bufferOpen(buf, geoJSON, pretty)

	jsonBuf.bufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.bufferWrite(nil)

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		jsonBuf.bufferClose()
		actual := buf.String()
		expected := "{\"type\":\"FeatureCollection\",\"features\":[{\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}}]}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBufferGeoJSON2(t *testing.T) {
	buf := &bytes.Buffer{}

	geoJSON := true
	pretty := false

	jsonBuf := jsonBuffer{}
	jsonBuf.bufferOpen(buf, geoJSON, pretty)

	jsonBuf.bufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.bufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.bufferWrite(nil)

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		jsonBuf.bufferClose()
		actual := buf.String()
		expected := "{\"type\":\"FeatureCollection\",\"features\":[{\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}},{\"id\":\"urn:ngsi-ld:TemperatureSensor:002\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}}]}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBufferGeoJSON1Pretty(t *testing.T) {
	buf := &bytes.Buffer{}

	geoJSON := true
	pretty := true

	jsonBuf := jsonBuffer{}
	jsonBuf.bufferOpen(buf, geoJSON, pretty)

	jsonBuf.bufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.bufferWrite(nil)

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		jsonBuf.bufferClose()
		actual := buf.String()
		expected := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": [{\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}}]\n}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBufferGeoJSON2Pretty(t *testing.T) {
	buf := &bytes.Buffer{}

	geoJSON := true
	pretty := true

	jsonBuf := jsonBuffer{}
	jsonBuf.bufferOpen(buf, geoJSON, pretty)

	jsonBuf.bufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.bufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.bufferWrite(nil)

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		jsonBuf.bufferClose()
		actual := buf.String()
		expected := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": [{\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}},{\"id\":\"urn:ngsi-ld:TemperatureSensor:002\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}}]\n}"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}

func TestBufferClose(t *testing.T) {
	buf := &bytes.Buffer{}

	jsonBuf := jsonBuffer{}
	jsonBuf.bufferOpen(buf, false, false)

	jsonBuf.bufferWrite([]byte("[abc]"))
	jsonBuf.bufferWrite([]byte("[xyz]"))
	jsonBuf.bufferClose()

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		actual := buf.String()
		expected := "[abc,xyz]"
		assert.Equal(t, expected, actual)
	} else {
		t.FailNow()
	}
}
