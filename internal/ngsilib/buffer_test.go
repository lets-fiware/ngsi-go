/*
MIT License

Copyright (c) 2020-2022 Kazuhito Suda

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

package ngsilib

import (
	"bytes"
	"os"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
)

func TestBufferOpen(t *testing.T) {
	jsonBuf := NewJsonBuffer()
	jsonBuf.BufferOpen(os.Stdout, false, false)

	if !assert.Equal(t, "[", jsonBuf.delimiter) {
		t.FailNow()
	}
}

func TestBufferWrite1(t *testing.T) {
	buffer := &bytes.Buffer{}

	jsonBuf := NewJsonBuffer()
	jsonBuf.BufferOpen(buffer, false, false)

	jsonBuf.BufferWrite([]byte("[abc]"))

	if !assert.Equal(t, []byte("abc"), jsonBuf.buf) {
		t.FailNow()
	}
}

func TestBufferWrite2(t *testing.T) {
	buf := &bytes.Buffer{}

	jsonBuf := NewJsonBuffer()
	jsonBuf.BufferOpen(buf, false, false)

	jsonBuf.BufferWrite([]byte("[abc]"))
	jsonBuf.BufferWrite([]byte("[xyz]"))

	if assert.Equal(t, []byte("xyz"), jsonBuf.buf) {
		jsonBuf.BufferClose()
		actual := buf.String()
		expected := "[abc,xyz]"
		assert.Equal(t, expected, actual)
	}
}

func TestBufferWrite3(t *testing.T) {
	buf := &bytes.Buffer{}

	jsonBuf := NewJsonBuffer()
	jsonBuf.BufferOpen(buf, false, false)

	jsonBuf.BufferWrite([]byte("[abc]"))
	jsonBuf.BufferWrite([]byte("[xyz]"))
	jsonBuf.BufferWrite(nil)

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		jsonBuf.BufferClose()
		actual := buf.String()
		expected := "[abc,xyz]"
		assert.Equal(t, expected, actual)
	}
}

func TestBufferGeoJSON1(t *testing.T) {
	buf := &bytes.Buffer{}

	geoJSON := true
	pretty := false

	jsonBuf := NewJsonBuffer()
	jsonBuf.BufferOpen(buf, geoJSON, pretty)

	jsonBuf.BufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.BufferWrite(nil)

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		jsonBuf.BufferClose()
		actual := buf.String()
		expected := "{\"type\":\"FeatureCollection\",\"features\":[{\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestBufferGeoJSON2(t *testing.T) {
	buf := &bytes.Buffer{}

	geoJSON := true
	pretty := false

	jsonBuf := NewJsonBuffer()
	jsonBuf.BufferOpen(buf, geoJSON, pretty)

	jsonBuf.BufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.BufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.BufferWrite(nil)

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		jsonBuf.BufferClose()
		actual := buf.String()
		expected := "{\"type\":\"FeatureCollection\",\"features\":[{\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}},{\"id\":\"urn:ngsi-ld:TemperatureSensor:002\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}}]}"
		assert.Equal(t, expected, actual)
	}
}

func TestBufferGeoJSON1Pretty(t *testing.T) {
	buf := &bytes.Buffer{}

	geoJSON := true
	pretty := true

	jsonBuf := NewJsonBuffer()
	jsonBuf.BufferOpen(buf, geoJSON, pretty)

	jsonBuf.BufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.BufferWrite(nil)

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		jsonBuf.BufferClose()
		actual := buf.String()
		expected := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": [{\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}}]\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestBufferGeoJSON2Pretty(t *testing.T) {
	buf := &bytes.Buffer{}

	geoJSON := true
	pretty := true

	jsonBuf := NewJsonBuffer()
	jsonBuf.BufferOpen(buf, geoJSON, pretty)

	jsonBuf.BufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:001","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.BufferWrite([]byte(`[{"id":"urn:ngsi-ld:TemperatureSensor:002","type":"Feature","properties":{"type":"TemperatureSensor","temperature":{"type":"Property","value":25,"unitCode":"CEL"},"location":{"type":"GeoProperty","value":{"type":"Point","coordinates":[139.76,35.68]}}}}]`))
	jsonBuf.BufferWrite(nil)

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		jsonBuf.BufferClose()
		actual := buf.String()
		expected := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": [{\"id\":\"urn:ngsi-ld:TemperatureSensor:001\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}},{\"id\":\"urn:ngsi-ld:TemperatureSensor:002\",\"type\":\"Feature\",\"properties\":{\"type\":\"TemperatureSensor\",\"temperature\":{\"type\":\"Property\",\"value\":25,\"unitCode\":\"CEL\"},\"location\":{\"type\":\"GeoProperty\",\"value\":{\"type\":\"Point\",\"coordinates\":[139.76,35.68]}}}}]\n}"
		assert.Equal(t, expected, actual)
	}
}

func TestBufferClose(t *testing.T) {
	buf := &bytes.Buffer{}

	jsonBuf := NewJsonBuffer()
	jsonBuf.BufferOpen(buf, false, false)

	jsonBuf.BufferWrite([]byte("[abc]"))
	jsonBuf.BufferWrite([]byte("[xyz]"))
	jsonBuf.BufferClose()

	if assert.Equal(t, []uint8([]byte(nil)), jsonBuf.buf) {
		actual := buf.String()
		expected := "[abc,xyz]"
		assert.Equal(t, expected, actual)
	}
}
