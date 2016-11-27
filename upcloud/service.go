package upcloud

import (
	log "github.com/Sirupsen/logrus"

	upcloud "github.com/Jalle19/upcloud-go-sdk/upcloud"
	upcloud_client "github.com/Jalle19/upcloud-go-sdk/upcloud/client"
	upcloud_service "github.com/Jalle19/upcloud-go-sdk/upcloud/service"
)

/**
 * UpCloud SDK service wrapper
 */

// Constructor for UpcloudServiceSettings
func New_UpcloudServiceSettings(client upcloud_client.Client) *UpcloudServiceSettings {
	return &UpcloudServiceSettings{
		client: client,
	}
}

// Settings for the
type UpcloudServiceSettings struct {
	client upcloud_client.Client
}

// Get an Upcloud service from these settings
func (serviceSettings UpcloudServiceSettings) Service() *upcloud_service.Service {
	return New_UpcloudServiceFromClient(serviceSettings.client)
}

// Get an Upcloud service from these settings
func (serviceSettings UpcloudServiceSettings) ServiceWrapper() *UpcloudServiceWrapper {
	service := serviceSettings.Service()
	return New_UpcloudServiceWrapper(*service)
}

// Constructor for upcloud Service from a client
func New_UpcloudServiceFromClient(client upcloud_client.Client) *upcloud_service.Service {
	service := upcloud_service.New(&client)
	return service
}

// Constructor for UpcloudServiceWrapper
func New_UpcloudServiceWrapper(service upcloud_service.Service) *UpcloudServiceWrapper {
	return &UpcloudServiceWrapper{
		Service: service,
	}
}

// Define some values that can be used by the ServiceWrapper to limit and configure it
type UpcloudBuilderSettings struct {
	Tags     []string `yml:"Tags"`
	Hosts    []string `yml:"Hosts"`
	Zones    []string `yml:"Zones"`
	Storages []string `yml:"Storages"`
}

// Merge settings
func (settings *UpcloudBuilderSettings) Merge(merge UpcloudBuilderSettings) {
	// merge hosts
	for _, host := range merge.Hosts {
		exists := false
		for _, existing := range settings.Hosts {
			if existing == host {
				exists = true
				break
			}
		}
		if !exists {
			settings.Hosts = append(settings.Hosts, host)
		}
	}

	log.WithFields(log.Fields{"settings": settings}).Debug("Merged UpCloud settings")
}

// It doesn't want to automatically marshal, so do it manually @TODO why isn't it unmarshalling automatically?
func (settings *UpcloudBuilderSettings) UnmarshalYAML(unmarshal func(interface{}) error) error {
	placeholder := map[string][]string{}
	if err := unmarshal(&placeholder); err != nil {
		return err
	}

	if hosts, defined := placeholder["Hosts"]; defined {
		for _, host := range hosts {
			exists := false
			for _, existing := range settings.Hosts {
				if existing == host {
					exists = true
					break
				}
			}
			if !exists {
				settings.Hosts = append(settings.Hosts, host)
			}
		}
	}
	if tags, defined := placeholder["Tags"]; defined {
		for _, tag := range tags {
			exists := false
			for _, existing := range settings.Tags {
				if existing == tag {
					exists = true
					break
				}
			}
			if !exists {
				settings.Tags = append(settings.Tags, tag)
			}
		}
	}
	if zones, defined := placeholder["Zones"]; defined {
		for _, zone := range zones {
			exists := false
			for _, existing := range settings.Zones {
				if existing == zone {
					exists = true
					break
				}
			}
			if !exists {
				settings.Zones = append(settings.Zones, zone)
			}
		}
	}
	return nil
}

// Does this server match settings from the BuilderSettings (is it in this project)
func (settings *UpcloudBuilderSettings) ServerUUIDAllowed(uuid string) bool {
	// simple host UUID match
	for _, match := range settings.Hosts {
		if match == uuid {
			return true
		}
	}
	return false
}

// Does this storage match settings from the BuilderSettings (is it in this project)
func (settings *UpcloudBuilderSettings) StorageUUIDAllowed(uuid string) bool {
	// simple host UUID match
	for _, match := range settings.Storages {
		if match == uuid {
			return true
		}
	}
	return false
}

// Does this server match settings from the BuilderSettings (is it in this project)
func (settings *UpcloudBuilderSettings) ZoneAllowed(zone upcloud.Zone) bool {
	// simple host UUID match
	for _, match := range settings.Zones {
		if match == zone.Id {
			return true
		}
	}
	return false
}

// Wrapper for the upcloud service, so that we can limit operations
type UpcloudServiceWrapper struct {
	upcloud_service.Service
}
