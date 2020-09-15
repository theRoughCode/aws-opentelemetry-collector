package config

import (
	"aws-observability.io/collector/pkg/consts"
	"aws-observability.io/collector/pkg/defaultcomponents"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/config"
	"os"
	"reflect"
	"testing"
)

func TestGetCfgFactory(t *testing.T) {
	cfgFunc := GetCfgFactory()
	assert.True(t, reflect.TypeOf(cfgFunc).Kind() == reflect.Func)
}

func TestGetCfgFactoryContainer(t *testing.T) {
	os.Setenv(consts.RUN_IN_CONTAINER, "True")
	os.Setenv(consts.AOC_CONFIG_CONTENT, "extensions:\n  health_check:\n  pprof:\n    endpoint: 0.0.0.0:1777\nreceivers:\n  otlp:\n    protocols:\n      grpc:\n        endpoint: 0.0.0.0:55680\nprocessors:\n  batch:\n  queued_retry:\nexporters:\n  logging:\n    loglevel: debug\n  awsxray:\n    local_mode: true\n    region: 'us-west-2'\n  awsemf:\n    local_mode: true\n    region: 'us-west-2'\nservice:\n  pipelines:\n    traces:\n      receivers: [prometheusreceiver]\n      exporters: [logging,awsxray]\n    metrics:\n      receivers: [prometheusreceiver]\n      exporters: [awsemf]\n  extensions: [pprof]")
	v := config.NewViper()
	factories, _ := defaultcomponents.Components()
	cfgFunc := GetCfgFactory()
	cfgModel, _ := cfgFunc(v, factories)
	assert.True(t, cfgModel.Receivers["otlp"] != nil)
	assert.True(t, cfgModel.Receivers["prometheus"] == nil)
	assert.True(t, cfgModel.Exporters["awsemf"] != nil)
	assert.True(t, cfgModel.Processors["queued_retry"] != nil)
	assert.True(t, cfgModel.Extensions["pprof"] != nil)
}
