### Run AOC on Debian and Windows hosts,

### Run AOC on AWS Windows Ec2 Host

To run AOC on AWS windows ec2 host, you can choose to install AOC MSI on your host by the following steps.

**Steps,**
1. Login on AWS Windows EC2 host and download aws-observability-collector MSI with the following command.
```
wget https://aws-opentelemetry-collector-release.s3.amazonaws.com/windows/amd64/latest/aws-opentelemetry-collector.msi
```
2. Install aws-observability-collector MSI by running the following command on the host
```
msiexec /i aws-opentelemetry-collector.msi
```
3. Once MSI is installed, it will create AOC in directory C:\Program Files\Amazon\AwsOpentelemetryCollector
