/*
Package v1 plugins are ways to add extended functionality to the Secretless Broker

WARNING: Given the speed of development, there will likely be cases of outdated documentation so please use this document
as a reference point and use the source code in this folder as the true representation of the API state!

Supported plugin types:
  -Listeners
  -Handlers
  -Providers
  -Configuration managers
  -Connection managers

There is also an additional EventNotifier class used to bubble up events from listeners and handlers up
to the plugin manager but this class may be removed as we move more of the abstract functionality to the plugin manager
itself.

All plugins are currently loaded in the following manner:
  - Directory in `/usr/local/lib/secretless` is listed and any `*.so` files are iterated over. Sub-directory traversal is not supported at this time.
  - Each shared library plugin is searched for these variables:
    - PluginAPIVersion
    - PluginInfo
    - GetHandlers
    - GetListeners
    - GetProviders
    - GetConfigurationManagers
    - GetConnectionManagers

  - All plugin factories are enumerated:
    - Handler plugins are added to handler factory map.
    - Listener plugins are added to listener factory map.
    - Provider plugins are added to provider factory map.
    - Connection manager plugins are added to connection manager factory map.
    - Configuration manager plugins are added to configuration manager factory map.

  - Connection manager plugins are instantiated
  - Chosen configuration manager plugin is instantiated
  - Program waits for a valid configuration to be provided

  - After the configuration is provided and loaded:
    - Providers are instantiated.
    - Listeners/handlers are instantiated as needed


PluginAPIVersion

`PluginAPIVersion` (returns string) indicates the target API version of the Secretless Broker and must match the
supported version found at https://github.com/cyberark/secretless-broker/blob/master/internal/pkg/plugin/manager.go#L108 list in the
main daemon.

PluginInfo

This `string->string` map (returns `map[string]string`) has information about the plugin that the daemon might use for logging, prioritization, and masking.
While extraneous keys in the map are ignored, the map _must_ contain the following keys:

  - `version`
  Indicates the plugin version
  - `id`
  A computer-friendly id of the plugin. Naming should be constrained to short, spaceless ASCII lowercase alphanumeric set with a limited set of special characters (`-`, `_`, and `/`).
  - `name`
  User-friendly name of the plugin. This name will be used in most user-facing messages about the plugin and should be constrained in length to <30 chars.
  - `description`
  A longer description of the plugin though it should not exceed 100 characters.

GetListeners

Returns a map of provided listener ids to their factory methods (`map[string]func(v1.ListenerOptions) v1.Listener`) that
accept v1.ListenerOptions when invoked and return a new v1.Listener.

GetHandlers

Returns a map of provided handler ids to their factory methods (`map[string]func(v1.HandlerOptions) v1.Handler`) that
accept v1.HandlerOptions when invoked and return a new v1.Handler.

GetProviders

Returns a map of provided provider ids to their factory methods (`map[string]func(v1.ProviderOptions) v1.Provider`) that
accept v1.ProviderOptions when invoked and return a new v1.Provider.

GetConnectionManagers

Returns a map of provided connection manager ids to their factory methods (`map[string]func() v1.ConnectionManager`) that
return a new v1.ConnectionManager connection manager when invoked.

Note: There is a high likelihood that this method will also have `v1.ConnectionManagerOptions` as the
factory parameter like the rest of the factory maps in the future releases

GetConfigurationManagers

Returns a map of provided configuration manager ids to their factory methods (`map[string]func() v1.ConfigurationManager`) that
return a new v1.ConfigurationManager manager when invoked.

Example plugin

The following shows a sample plugin that conforms to the expected API:

  package main

  import (
  	plugin_v1 "github.com/cyberark/secretless-broker/pkg/secretless/plugin/v1"
  	"github.com/cyberark/secretless-broker/test/plugin/example"
  )

  // PluginAPIVersion is the API version being used
  var PluginAPIVersion = "0.0.8"

  // PluginInfo describes the plugin
  var PluginInfo = map[string]string{
  	"version":     "0.0.8",
  	"id":          "example-plugin",
  	"name":        "Example Plugin",
  	"description": "Example plugin to demonstrate plugin functionality",
  }

  // GetListeners returns the echo listener
  func GetListeners() map[string]func(plugin_v1.ListenerOptions) plugin_v1.Listener {
  	return map[string]func(plugin_v1.ListenerOptions) plugin_v1.Listener{
  		"echo": example.ListenerFactory,
  	}
  }

  // GetHandlers returns the example handler
  func GetHandlers() map[string]func(plugin_v1.HandlerOptions) plugin_v1.Handler {
  	return map[string]func(plugin_v1.HandlerOptions) plugin_v1.Handler{
  		"example-handler": example.HandlerFactory,
  	}
  }

  // GetProviders returns the example provider
  func GetProviders() map[string]func(plugin_v1.ProviderOptions) plugin_v1.Provider {
  	return map[string]func(plugin_v1.ProviderOptions) plugin_v1.Provider{
  		"example-provider": example.ProviderFactory,
  	}
  }

  // GetConfigurationManagers returns the example configuration manager
  func GetConfigurationManagers() map[string]func(plugin_v1.ConfigurationManagerOptions) plugin_v1.ConfigurationManager {
  	return map[string]func(plugin_v1.ConfigurationManagerOptions) plugin_v1.ConfigurationManager{
  		"example-plugin-config-manager": example.ConfigManagerFactory,
  	}
  }

  // GetConnectionManagers returns the example connection manager
  func GetConnectionManagers() map[string]func() plugin_v1.ConnectionManager {
  	return map[string]func() plugin_v1.ConnectionManager{
  		"example-plugin-connection-manager": example.ManagerFactory,
  	}
  }

*/
package v1
