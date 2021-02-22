# admin loggers - Convenience command

This command allows you to manage loggers for Cygnus.

-   [List loggers](#list-loggers)
-   [Get a logger](#get-a-logger)
-   [Create a logger](#create-a-logger)
-   [Update a logger](#update-a-logger)
-   [Delete a logger](#delete-a-logger)

## Common Options

| Options                   | Description                |
| ------------------------- | -------------------------- |
| --host value, -h value    | specify host or alias      |
| --token value             | specify oauth token        |
| --help                    | show help (default: false) |

<a name="list-loggers"></a>

## List loggers

This command lists loggers for Cygnus

```console
ngsi admin [command options] loggers list [options]
```

### Options

| Options         | Description                                                                   |
| --------------- | ----------------------------------------------------------------------------- |
| --transient, -t | true, retrieving from memory, or false, retrieving from file (default: false) |
| --pretty, -P    | pretty format (default: false)                                                |
| --help          | show help (default: false)                                                    |

### Example

#### Request:

```console
ngsi admin --host cygnus loggers list --pretty
```

```json
{
  "success": "true",
  "loggers": [
    {
      "name": "com.telefonica.iot.cygnus.management.ManagementInterface",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.jmx.MBeanContainer",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.servlet.BaseHolder",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.jmx.ObjectMBean",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.http.MimeTypes",
      "level": null
    },
    {
      "name": "org.apache.flume.lifecycle.LifecycleSupervisor",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.server.SecureRequestCustomizer",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.io.ManagedSelector",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.security.SecurityHandler",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.util.component.AbstractLifeCycle",
      "level": null
    },
    {
      "name": "com.telefonica.iot.cygnus.handlers.NGSIRestHandler",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.util.thread.strategy.EatWhatYouKill",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.server.session",
      "level": null
    },
    {
      "name": "com.telefonica.iot.cygnus.utils.NGSIUtils",
      "level": null
    },
    {
      "name": "org.apache.flume.source.http.HTTPSource",
      "level": null
    },
    {
      "name": "org.apache.flume.channel.MemoryChannel",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.http.HttpFields",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.util.URIUtil",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.server.Server",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.util.thread.QueuedThreadPool",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.server.handler.jmx.AbstractHandlerMBean",
      "level": null
    },
    {
      "name": "com.telefonica.iot.cygnus.nodes.CygnusApplication",
      "level": null
    },
    {
      "name": "org.apache.avro.ipc.NettyTransceiver",
      "level": "WARN"
    },
    {
      "name": "com.telefonica.iot.cygnus.interceptors.NGSINameMappingsInterceptor",
      "level": null
    },
    {
      "name": "org.apache.flume.util.SSLUtil",
      "level": null
    },
    {
      "name": "org.apache.flume.source.DefaultSourceFactory",
      "level": null
    },
    {
      "name": "com.telefonica.iot.cygnus.containers.NameMappings",
      "level": null
    },
    {
      "name": "org.apache.flume.SinkRunner",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.io.SelectorManager",
      "level": null
    },
    {
      "name": "org.apache.flume.conf.FlumeConfiguration",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.util.thread.Invocable$InvocableExecutor",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.server.handler.AllowSymLinkAliasChecker",
      "level": null
    },
    {
      "name": "com.telefonica.iot.cygnus.interceptors.NGSINameMappingsInterceptor$PeriodicalNameMappingsReader",
      "level": null
    },
    {
      "name": "org.mongodb",
      "level": "WARN"
    },
    {
      "name": "org.apache.http",
      "level": "WARN"
    },
    {
      "name": "org.mortbay.log",
      "level": null
    },
    {
      "name": "com.telefonica.iot.cygnus.http.JettyServer",
      "level": null
    },
    {
      "name": "org.apache.flume.node.Application",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.http.pathmap.PathMappings",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.server.handler.ContextHandler.ROOT",
      "level": null
    },
    {
      "name": "org.apache.flume.node.AbstractConfigurationProvider",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.server.handler.AbstractHandler",
      "level": null
    },
    {
      "name": "org.apache.flume.lifecycle",
      "level": "WARN"
    },
    {
      "name": "com.telefonica.iot.cygnus.backends.mongo.MongoBackendImpl",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.servlet.Holder",
      "level": null
    },
    {
      "name": "org.apache.flume.channel.ChannelSelectorFactory",
      "level": null
    },
    {
      "name": "org.I0Itec",
      "level": "WARN"
    },
    {
      "name": "org.eclipse.jetty.server.AbstractConnector",
      "level": null
    },
    {
      "name": "org.apache.kafka",
      "level": "WARN"
    },
    {
      "name": "org.eclipse.jetty.util.StringUtil",
      "level": null
    },
    {
      "name": "com.telefonica.iot.cygnus.utils.CommonUtils",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.util.log",
      "level": null
    },
    {
      "name": "org.apache.flume.instrumentation.MonitoredCounterGroup",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.servlet.ServletHolder",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.util.DeprecationWarning",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.http.HttpGenerator",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.util.component.ContainerLifeCycle",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.server.handler.ContextHandler",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.servlet.ServletContextHandler",
      "level": null
    },
    {
      "name": "org.apache.hadoop",
      "level": "WARN"
    },
    {
      "name": "org.eclipse.jetty.servlet.ServletHandler",
      "level": null
    },
    {
      "name": "org.mortbay",
      "level": "WARN"
    },
    {
      "name": "org.jboss",
      "level": "WARN"
    },
    {
      "name": "org.apache.flume.channel.ChannelProcessor",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.http.PreEncodedHttpField",
      "level": null
    },
    {
      "name": "com.amazonaws",
      "level": "WARN"
    },
    {
      "name": "org.apache.flume.node.PropertiesFileConfigurationProvider",
      "level": null
    },
    {
      "name": "org.apache.flume.sink.DefaultSinkFactory",
      "level": null
    },
    {
      "name": "com.telefonica.iot.cygnus.sinks.NGSIMongoBaseSink",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.server.handler.ErrorHandler",
      "level": null
    },
    {
      "name": "org.apache.flume.channel.DefaultChannelFactory",
      "level": null
    },
    {
      "name": "org.apache.flume.node.PollingPropertiesFileConfigurationProvider",
      "level": null
    },
    {
      "name": "org.eclipse.jetty.util.TypeUtil",
      "level": null
    },
    {
      "name": "com.telefonica.iot.cygnus.management.LogHandlers",
      "level": null
    },
    {
      "name": "com.telefonica.iot.cygnus.channels.CygnusMemoryChannel",
      "level": null
    },
    {
      "name": "org.apache.zookeeper",
      "level": "WARN"
    },
    {
      "name": "org.eclipse.jetty.util.DecoratedObjectFactory",
      "level": null
    },
    {
      "name": "com.telefonica.iot.cygnus.sinks.NGSISink",
      "level": null
    }
  ]
}
```

<a name="get-a-logger"></a>

## Get a logger

This command gets a logger for Cygnus

```console
ngsi admin [command options] loggers get [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --name value, -n value | logger name                                                                   |
| --transient, -t        | true, retrieving from memory, or false, retrieving from file (default: false) |
| --pretty, -P           | pretty format (default: false)                                                |
| --help                 | show help (default: false)                                                    |

### Example

#### Request:

```console
ngsi admin loggers get --name org.mongodb
```

```json
{"success":"true","logger":"[{"name":"org.mongodb","level":"WARN"}]"}
```

<a name="create-a-logger"></a>

## Create a logger

This command creates a logger for Cygnus

```console
ngsi admin [command options] loggers create [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --name value, -n value | logger name                                                                   |
| --data value, -d value | logger information                                                            |
| --transient, -t        | true, retrieving from memory, or false, retrieving from file (default: false) |
| --pretty, -P           | pretty format (default: false)                                                |
| --help                 | show help (default: false)                                                    |

### Example

#### Request:

```console
ngsi admin --host cygnus loggers create --data '{"logger":{"name":"org.mongodb","level":"WARN"}}'
```

```json
{"success":"true","result":"Logger 'org.mongodb' put"}
```

<a name="update-a-logger"></a>

## Update a logger

This command updates a logger for Cygnus

```console
ngsi admin [command options] loggers update [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --name value, -n value | logger name                                                                   |
| --data value, -d value | logger information                                                            |
| --transient, -t        | true, retrieving from memory, or false, retrieving from file (default: false) |
| --pretty, -P           | pretty format (default: false)                                                |
| --help                 | show help (default: false)                                                    |

### Example

#### Request:

```console
ngsi admin --host cygnus loggers update --data '{"logger":{"name":"org.mongodb","level":"INFO"}}'
```

```json
{"success":"true","result":"Logger 'org.mongodb' updated succesfully"}
```

<a name="delete-a-logger"></a>

## Delete a logger

This command deletes a logger for Cygnus

```console
ngsi admin [command options] loggers delete [options]
```

### Options

| Options                | Description                                                                   |
| ---------------------- | ----------------------------------------------------------------------------- |
| --name value, -n value | logger name i                                                                 |
| --transient, -t        | true, retrieving from memory, or false, retrieving from file (default: false) |
| --pretty, -P           | pretty format (default: false)                                                |
| --help                 | show help (default: false)                                                    |

### Example

#### Request:

```console
ngsi admin loggers delete --name org.mongodb
```

```json
{"success":"true","result":" Logger 'org.mongodb' removed succesfully"}
```
