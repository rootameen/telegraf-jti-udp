# JTI Parser Plugin

This plugin parses Juniper Networks implementation of UDP data from sensors sending data on the specified port using Junos Telemetry Interface and converts it into an easy to store format. 

### Configuration:

```toml
[[inputs.socket_listener]]
   ## URL to listen on
   # service_address = "tcp://:8094"
#   # service_address = "tcp://127.0.0.1:http"
#   # service_address = "tcp4://:8094"
#   # service_address = "tcp6://:8094"
#   # service_address = "tcp6://[2001:db8::1]:8094"
#   # Specify the host port on which you are receiving the juniper UDP sensor data
    service_address = "udp://:29000"
#   # service_address = "udp4://:8094"
#   # service_address = "udp6://:8094"
#   # service_address = "unix:///tmp/telegraf.sock"
#   # service_address = "unixgram:///tmp/telegraf.sock"
#
#   ## Maximum number of concurrent connections.
#   ## Only applies to stream sockets (e.g. TCP).
#   ## 0 (default) is unlimited.
#   # max_connections = 1024
#
#   ## Read timeout.
#   ## Only applies to stream sockets (e.g. TCP).
#   ## 0 (default) is unlimited.
#   # read_timeout = "30s"
#
#   ## Maximum socket buffer size in bytes.
#   ## For stream sockets, once the buffer fills up, the sender will start backing up.
#   ## For datagram sockets, once the buffer fills up, metrics will start dropping.
#   ## Defaults to the OS default.
#   # read_buffer_size = 65535
#
#   ## Period between keep alive probes.
#   ## Only applies to TCP sockets.
#   ## 0 disables keep alive probes.
#   ## Defaults to the OS configuration.
#   # keep_alive_period = "5m"
#
#   ## Data format to consume.
#   ## Each data format has its own unique set of configuration options, read
#   ## more about them here:
#   ## https://github.com/influxdata/telegraf/blob/master/docs/DATA_FORMATS_INPUT.md
#   ## Specify "juniperUDP" data format to use this juniper specific parser plugin
    data_format = "juniperUDP"

```

### Tags:

- All measurements are tagged appropriately using the identifier information
  in incoming data
