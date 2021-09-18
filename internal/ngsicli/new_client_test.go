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

package ngsicli

import (
	"testing"

	"github.com/lets-fiware/ngsi-go/internal/assert"
	"github.com/lets-fiware/ngsi-go/internal/ngsierr"
)

func TestNewClient(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = NewClient(c.Ngsi, c, false, []string{"brokerv2"})
	assert.NoError(t, err)
}

func TestNewClientSkipGetToken(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = NewClientSkipGetToken(c.Ngsi, c, false, []string{"brokerv2"})
	assert.NoError(t, err)
}

func TestNewClientSub(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, false, false)

	assert.NoError(t, err)
}

func TestNewClientError(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	link := linkFlag.Copy(true)
	err = link.SetValue("fiware")
	assert.NoError(t, err)

	c.Flags = []Flag{host, link}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, false, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestNewClientErrorNewClient(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("fiware")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, false, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: fiware", ngsiErr.Message)
	}
}

func TestNewClientErrorServerTypeNGSIv2(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion-ld")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, false, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	}
}

func TestNewClientErrorServerTypeNGSILD(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerld"}, false, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSI-LD", ngsiErr.Message)
	}
}

func TestNewClientErrorServerType1(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"keyrock"}, false, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	}
}

func TestNewClientErrorServerType2(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("keyrock")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, false, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by keyrock", ngsiErr.Message)
	}
}

func TestNewClientSkipGetTokenSub(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, true, false)

	assert.NoError(t, err)
}

func TestNewClientSkipGetTokenErrorParseFlags(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	link := linkFlag.Copy(true)
	err = link.SetValue("fiware")
	assert.NoError(t, err)

	c.Flags = []Flag{host, link}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, true, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestNewClientSkipGetTokenErrorNewClient(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("fiware")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, true, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: fiware", ngsiErr.Message)
	}
}

func TestNewClientSkipGetTokenErrorServerTypeNGSIV2(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion-ld")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, true, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	}
}

func TestNewClientSkipGetTokenErrorServerTypeNGSILD(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerld"}, true, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSI-LD", ngsiErr.Message)
	}
}

func TestNewClientSkipGetTokenErrorServerType1(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"keyrock"}, true, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	}
}

func TestNewClientSkipGetTokenErrorServerType(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("keyrock")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Host = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, true, false)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by keyrock", ngsiErr.Message)
	}
}

func TestNewClientDest(t *testing.T) {
	c := setupTestInitCmd()

	host := destinationFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Destination = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, false, true)

	assert.NoError(t, err)
}

func TestNewClientDestError(t *testing.T) {
	c := setupTestInitCmd()

	host := destinationFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	link := link2Flag.Copy(true)
	err = link.SetValue("fiware")
	assert.NoError(t, err)

	c.Flags = []Flag{host, link}

	c.Ngsi.Destination = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, false, true)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 1, ngsiErr.ErrNo)
		assert.Equal(t, "fiware not found", ngsiErr.Message)
	}
}

func TestNewClientDestErrorNewClient(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("fiware")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Destination = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, false, true)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 2, ngsiErr.ErrNo)
		assert.Equal(t, "error host: fiware (destination)", ngsiErr.Message)
	}
}

func TestNewClientDestErrorServerTypeNGSIv2(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion-ld")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Destination = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, false, true)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 3, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSIv2", ngsiErr.Message)
	}
}

func TestNewClientDestErrorServerTypeNGSILD(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Destination = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerld"}, false, true)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 4, ngsiErr.ErrNo)
		assert.Equal(t, "only available on NGSI-LD", ngsiErr.Message)
	}
}

func TestNewClientDestErrorServerType1(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("orion")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Destination = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"wirecloud"}, false, true)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by broker", ngsiErr.Message)
	}
}

func TestNewClientDestErrorServerType2(t *testing.T) {
	c := setupTestInitCmd()

	host := hostFlag.Copy(true)
	err := host.SetValue("keyrock")
	assert.NoError(t, err)

	c.Flags = []Flag{host}

	c.Ngsi.Destination = host.(*StringFlag).Value

	_, err = newClient(c.Ngsi, c, false, []string{"brokerv2"}, false, true)

	if assert.Error(t, err) {
		ngsiErr := err.(*ngsierr.NgsiError)
		assert.Equal(t, 5, ngsiErr.ErrNo)
		assert.Equal(t, "not supported by keyrock", ngsiErr.Message)
	}
}
