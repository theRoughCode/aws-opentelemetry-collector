/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package main

import (
	"aws-observability.io/collector/pkg/defaultcomponents"
	"aws-observability.io/collector/pkg/logger"
	"aws-observability.io/collector/tools/version"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/service"
	"go.opentelemetry.io/collector/service/builder"
	"go.uber.org/zap/zapcore"
	"log"
)

func main() {
	logger.SetupErrorLogger()
	handleErr := func(message string, err error) {
		if err != nil {
			log.Fatalf(
				"%s: %v", message, err)
		}
	}

	factories, err := defaultcomponents.Components()
	handleErr("Failed to build components", err)

	// configuration factory
	cfgFactory := func(otelViper *viper.Viper, f config.Factories) (*configmodels.Config, error) {
		// use the default config
		if len(builder.GetConfigFile()) == 0 {
			handleErr("configuration file is not provided", nil)
			// TODO - load default config?
		}
		// use OTel yaml config from input
		otelCfg, err := service.FileLoaderConfigFactory(otelViper, f)
		if err != nil {
			return nil, err
		}
		return otelCfg, nil
	}

	lumberHook := logger.GetLumberHook()
	info := service.ApplicationStartInfo{
		ExeName:  "aws-opentelemetry-collector",
		LongName: "AWS OpenTelemetry Collector",
		Version:  version.Version,
		GitHash:  version.GitHash,
	}

	if err := run(service.Parameters{
		Factories:            factories,
		ApplicationStartInfo: info,
		ConfigFactory:        cfgFactory,
		LoggingHooks:         []func(entry zapcore.Entry) error{lumberHook}}); err != nil {
		log.Fatal(err)
	}

}

func runInteractive(params service.Parameters) error {
	app, err := service.New(params)
	if err != nil {
		return errors.Wrap(err, "failed to construct the application")
	}

	err = app.Start()
	if err != nil {
		return errors.Wrap(err, "application run finished with error: %v")
	}

	return nil
}
