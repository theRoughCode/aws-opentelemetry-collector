[![codecov](https://codecov.io/gh/mxiamxia/aws-opentelemetry-collector/branch/master/graph/badge.svg)](https://codecov.io/gh/mxiamxia/aws-opentelemetry-collector)
![CI](https://github.com/mxiamxia/aws-opentelemetry-collector/workflows/CI/badge.svg)
![CD](https://github.com/mxiamxia/aws-opentelemetry-collector/workflows/CD/badge.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/mxiamxia/aws-opentelemetry-collector)


### Overview

AWS Observability Collector is a certified Amazon distribution of OpenTelemetry Collector. It will fully support AWS CloudWatch Metrics, Traces and Logs with correlations and export your data from AWS to the other monitoring parterns backend services.

### Getting Help

Use the following community resources for getting help with AWS Observability Collector. We use the GitHub issues for tracking bugs and feature requests.

* Ask a question in [AWS CloudWatch Forum](https://forums.aws.amazon.com/forum.jspa?forumID=138).
* Open a support ticket with [AWS Support](http://docs.aws.amazon.com/awssupport/latest/user/getting-started.html).
* If you think you may have found a bug, open an [issue](https://github.com/mxiamxia/aws-opentelemetry-collector/issues/new).
* For contributing guidelines refer [CONTRIBUTING.md](https://github.com/mxiamxia/aws-opentelemetry-collector/blob/master/CONTRIBUTING.md).

### Get Started

#### AOC Components
* [OpenTelemetry Collector-v0.7.0](https://github.com/open-telemetry/opentelemetry-collector/)
* [Trace X-Ray Exporter-v0.7.0](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/master/exporter/awsxrayexporter)
* Metrics EMF Exporter
* More coming

#### Try out AOC Beta
* [Run it in Docker](docs/developers/docker-demo.md)
* [Run it on AWS Linux EC2](docs/developers/linux-rpm-demo.md)
* [Run it on AWS Windows EC2](docs/developers/windows-other-demo.md)
* [Run it on AWS Debian](docs/developers/debian-deb-demo.md)

#### Build Your Own Executables
* [Build RPM/Deb/MSI](docs/developers/build-aoc.md)
* [Build Docker Image](docs/developers/build-aoc.md)
* more

### Release Process
* [Release new version](docs/developers/release-new-version.md)

### Benchmark

aws-observability-collector is based on open-telemetry-collector. Here is the benchmark of AWSXray trace exporter and AWSEMF metrics exporter.

This table shows the performance of AWSEMF exporter  against 1kData/sec,5kData/sec and 10kData/sec metrics

| Test                | Result | Duration | CPU Avg% | CPU Max% | RAM Avg MiB | RAM Max MiB | Sent Items | Received Items |
|---------------------|--------|----------|----------|----------|-------------|-------------|------------|----------------|
| Metric1kDPS/AWSEmf  | PASS   | 17s      |       34 |     36.3 |          59 |          82 |     105000 |         105000 |
| Metric5kDPS/AWSEmf  | PASS   | 43s      |     58.5 |    101.5 |         505 |         678 |     508200 |         508200 |
| Metric10kDPS/AWSEmf | PASS   | 72s      |       63 |    145.7 |         971 |        1178 |     849100 |         849100 |

This table shows the performance of AWSXray  exporter against 1kData/sec,5kData/sec and 10kData/sec spans(traces).

| Test                | Result | Duration | CPU Avg% | CPU Max% | RAM Avg MiB | RAM Max MiB | Sent Items | Received Items |
|---------------------|--------|----------|----------|----------|-------------|-------------|------------|----------------|
| Trace1kSPS/AwsXray  | PASS   | 15s      |      8.5 |     11.6 |          32 |          36 |      15000 |          15000 |
| Trace5kSPS/AwsXray  | PASS   | 15s      |    26.12 |     27.8 |          33 |          38 |      74400 |          74400 |
| Trace10kSPS/AwsXray | PASS   | 15s      |     43.8 |     45.3 |          37 |          43 |     132500 |         132500 |


### License
aws-observability-collector is under Apache 2.0 license.
