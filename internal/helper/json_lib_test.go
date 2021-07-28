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

package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/ngsilib"
	"github.com/stretchr/testify/assert"
)

type jsonLib struct {
}

func (j *jsonLib) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func (j *jsonLib) Encode(w io.Writer, v interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(v)
}

func (j *jsonLib) Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}

func (j *jsonLib) Valid(data []byte) bool {
	return json.Valid(data)
}

func TestJSONLibDecode(t *testing.T) {
	j := &MockJSONLib{Jsonlib: &jsonLib{}}

	v := make(map[string]interface{})
	r := bytes.NewReader([]byte(`{"name":"fiware"}`))

	err := j.Decode(r, &v)

	if assert.NoError(t, err) {
		assert.Equal(t, "fiware", v["name"].(string))
	}
}

func TestJSONLibDecodeErr(t *testing.T) {
	j := &MockJSONLib{DecodeErr: [5]error{errors.New("Decode error")}}

	v := make(map[string]interface{})
	r := bytes.NewReader([]byte(`{"name":"fiware"}`))

	err := j.Decode(r, &v)

	if assert.Error(t, err) {
		assert.Equal(t, "Decode error", err.Error())
	}
}

func TestJSONLibEncode(t *testing.T) {
	j := &MockJSONLib{Jsonlib: &jsonLib{}}

	v := make(map[string]interface{})
	v["name"] = "fiware"
	buf := &bytes.Buffer{}

	err := j.Encode(buf, &v)

	if assert.NoError(t, err) {
		assert.Equal(t, "{\"name\":\"fiware\"}\n", buf.String())
	}
}

func TestJSONLibEncodeError(t *testing.T) {
	j := &MockJSONLib{Jsonlib: &jsonLib{}, EncodeErr: [5]error{errors.New("Encode error")}}

	v := make(map[string]interface{})
	v["name"] = "fiware"
	buf := &bytes.Buffer{}

	err := j.Encode(buf, &v)

	if assert.Error(t, err) {
		assert.Equal(t, "Encode error", err.Error())
	}
}

func TestJSONLibSetJSONDecodeErr(t *testing.T) {
	j := &MockJSONLib{Jsonlib: &jsonLib{}}
	ngsi := &ngsilib.NGSI{JSONConverter: j}

	SetJSONDecodeErr(ngsi, 1)

	jj := (ngsi.JSONConverter).(*MockJSONLib)

	assert.Equal(t, "json error", jj.DecodeErr[1].Error())
}

func TestJSONLibSetJSONEncodeErr(t *testing.T) {
	j := &MockJSONLib{Jsonlib: &jsonLib{}}
	ngsi := &ngsilib.NGSI{JSONConverter: j}

	SetJSONEncodeErr(ngsi, 1)

	jj := (ngsi.JSONConverter).(*MockJSONLib)

	assert.Equal(t, "json error", jj.EncodeErr[1].Error())
}

func TestJSONLibIndent(t *testing.T) {
	j := &MockJSONLib{Jsonlib: &jsonLib{}}

	src := []byte(`{"orion":{"version":"2.6.1","uptime":"33d,1h,2m,18s","git_hash":"ec20c8bcfd883d6d7214a28818c8310fb179bbcf","compile_time":"TueMar3011:43:48UTC2021","compiled_by":"root","compiled_in":"7af2a988ed1d","release_date":"TueMar3011:43:48UTC2021","doc":"https://fiware-orion.rtfd.io/en/2.6.1/","libversions":{"boost":"1_53","libcurl":"libcurl/7.29.0NSS/3.53.1zlib/1.2.7libidn/1.28libssh2/1.8.0","libmicrohttpd":"0.9.70","openssl":"1.0.2k","rapidjson":"1.1.0","mongodriver":"legacy-1.1.2"}}}`)
	buf := &bytes.Buffer{}

	err := j.Indent(buf, src, "  ", "  ")

	if assert.NoError(t, err) {
		assert.Equal(t, "{\n    \"orion\": {\n      \"version\": \"2.6.1\",\n      \"uptime\": \"33d,1h,2m,18s\",\n      \"git_hash\": \"ec20c8bcfd883d6d7214a28818c8310fb179bbcf\",\n      \"compile_time\": \"TueMar3011:43:48UTC2021\",\n      \"compiled_by\": \"root\",\n      \"compiled_in\": \"7af2a988ed1d\",\n      \"release_date\": \"TueMar3011:43:48UTC2021\",\n      \"doc\": \"https://fiware-orion.rtfd.io/en/2.6.1/\",\n      \"libversions\": {\n        \"boost\": \"1_53\",\n        \"libcurl\": \"libcurl/7.29.0NSS/3.53.1zlib/1.2.7libidn/1.28libssh2/1.8.0\",\n        \"libmicrohttpd\": \"0.9.70\",\n        \"openssl\": \"1.0.2k\",\n        \"rapidjson\": \"1.1.0\",\n        \"mongodriver\": \"legacy-1.1.2\"\n      }\n    }\n  }", buf.String())
	}
}

func TestJSONLibIndentErr(t *testing.T) {
	j := &MockJSONLib{Jsonlib: &jsonLib{}, IndentErr: errors.New("Indent error")}

	src := []byte(`{"orion":{"version":"2.6.1","uptime":"33d,1h,2m,18s","git_hash":"ec20c8bcfd883d6d7214a28818c8310fb179bbcf","compile_time":"TueMar3011:43:48UTC2021","compiled_by":"root","compiled_in":"7af2a988ed1d","release_date":"TueMar3011:43:48UTC2021","doc":"https://fiware-orion.rtfd.io/en/2.6.1/","libversions":{"boost":"1_53","libcurl":"libcurl/7.29.0NSS/3.53.1zlib/1.2.7libidn/1.28libssh2/1.8.0","libmicrohttpd":"0.9.70","openssl":"1.0.2k","rapidjson":"1.1.0","mongodriver":"legacy-1.1.2"}}}`)
	buf := &bytes.Buffer{}

	err := j.Indent(buf, src, "  ", "  ")

	if assert.Error(t, err) {
		assert.Equal(t, "Indent error", err.Error())
	}
}

func TestJSONLibSetJSONIndentError(t *testing.T) {
	j := &MockJSONLib{Jsonlib: &jsonLib{}}
	ngsi := &ngsilib.NGSI{JSONConverter: j}

	SetJSONIndentError(ngsi)

	jj := (ngsi.JSONConverter).(*MockJSONLib)

	assert.Equal(t, "json error", jj.IndentErr.Error())
}

func TestJSONLibValid(t *testing.T) {
	j := &MockJSONLib{Jsonlib: &jsonLib{}}
	src := []byte(`{"orion":{"version":"2.6.1","uptime":"33d,1h,2m,18s","git_hash":"ec20c8bcfd883d6d7214a28818c8310fb179bbcf","compile_time":"TueMar3011:43:48UTC2021","compiled_by":"root","compiled_in":"7af2a988ed1d","release_date":"TueMar3011:43:48UTC2021","doc":"https://fiware-orion.rtfd.io/en/2.6.1/","libversions":{"boost":"1_53","libcurl":"libcurl/7.29.0NSS/3.53.1zlib/1.2.7libidn/1.28libssh2/1.8.0","libmicrohttpd":"0.9.70","openssl":"1.0.2k","rapidjson":"1.1.0","mongodriver":"legacy-1.1.2"}}}`)

	actual := j.Valid(src)

	assert.Equal(t, true, actual)
}

func TestJSONLibValidError(t *testing.T) {
	b := false
	j := &MockJSONLib{Jsonlib: &jsonLib{}, ValidErr: &b}

	actual := j.Valid(nil)

	assert.Equal(t, false, actual)
}
