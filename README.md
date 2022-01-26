
# telegraf-jti-udp

this repo is a _wip_ copy of [https://github.com/Juniper/telegraf-jti-plugins](https://github.com/Juniper/telegraf-jti-plugins), based on Telegraf 1.22

![screenshot](https://i.imgur.com/LAzrRwA.png)

## Minimum Requirements

Telegraf shares the same [minimum requirements][] as Go:

- Linux kernel version 2.6.23 or later
- Windows 7 or later
- FreeBSD 11.2 or later
- MacOS 10.11 El Capitan or later

[minimum requirements]: https://github.com/golang/go/wiki/MinimumRequirements#minimum-requirements

### Build From Source

Telegraf requires Go version 1.17 or newer, the Makefile requires GNU make.

1. [Install Go](https://golang.org/doc/install) >=1.17 (1.17.2 recommended)
2. Clone the repository:

   ```shell
   git clone https://github.com/rootameen/telegraf-jti-udp.git
   ```

3. Run `make` from the source directory

   ```shell
   cd telegraf-jti-udp
   make
   ```

or build for a specific OS

 ```shell
   cd telegraf-jti-udp
   env GOOS=linux GOARCH=amd64 go build -v ./cmd/telegraf
   ```

## Getting Started

Here's a sample config file `telegraf.conf` with the plugin enabled as an input:

```ini
[[inputs.socket_listener]]

 service_address = "udp://:29000"
 data_format = "juniperUDP"


[[outputs.influxdb_v2]]

 urls = ["http://${INFLUX_HOST}:8086"]
 token = "${INFLUX_TOKEN}"
 organization = "${INFLUX_ORG}"
 bucket = "${INFLUX_BUCKET}"
```

More details in the [plugin directory](https://github.com/rootameen/telegraf-jti-udp/tree/master/plugins/parsers/juniperUDP)

start the service:

```shell
./telegraf --config telegraf.conf
```

## Known Issues

Currently, it appears that the parsing of ingest data is not working properly. It is unknown if this state was the same in the original repo or not. The following list of issues is observed:

- if multiple devices are sending data, plugin is not tagging devices as source filter properly
- metrics received are being segmented into separate data points (i.e. Metric Name, Sequence Number, Value, Timestamp, Device are all seperate instead of being fields of the same data point)
  - this basically means that for interface traffic, for instance, the intefrace name and the metric value are two separate measurements, instead of being fields of the same measurement / data point
