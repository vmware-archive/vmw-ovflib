package ovf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Environment
     xmlns="http://schemas.dmtf.org/ovf/environment/1"
     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
     xmlns:oe="http://schemas.dmtf.org/ovf/environment/1"
     xmlns:ve="http://www.vmware.com/schema/ovfenv"
     oe:id=""
     ve:vCenterId="vm-12345">
   <PlatformSection>
      <Kind>VMware ESXi</Kind>
      <Version>5.5.0</Version>
      <Vendor>VMware, Inc.</Vendor>
      <Locale>en</Locale>
   </PlatformSection>
   <PropertySection>
         <Property oe:key="foo" oe:value="42"/>
         <Property oe:key="bar" oe:value="0"/>
   </PropertySection>
   <ve:EthernetAdapterSection>
      <ve:Adapter ve:mac="00:00:00:00:00:00" ve:network="foo" ve:unitNumber="7"/>
   </ve:EthernetAdapterSection>
</Environment>`)

func TestOvfEnvProperties(t *testing.T) {
	env := ReadEnvironment(data)
	props := env.Properties

	var val string
	var ok bool

	val, ok = props["foo"]
	assert.True(t, ok)
	assert.Equal(t, val, "42")

	val, ok = props["bar"]
	assert.True(t, ok)
	assert.Equal(t, val, "0")
}

func TestOvfEnvPlatform(t *testing.T) {
	env := ReadEnvironment(data)
	platform := env.Platform

	assert.Equal(t, platform.Kind, "VMware ESXi")
	assert.Equal(t, platform.Version, "5.5.0")
	assert.Equal(t, platform.Vendor, "VMware, Inc.")
	assert.Equal(t, platform.Locale, "en")
}
