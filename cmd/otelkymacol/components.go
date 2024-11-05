// Code generated by "go.opentelemetry.io/collector/cmd/builder". DO NOT EDIT.

package main

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/exporter"
	debugexporter "go.opentelemetry.io/collector/exporter/debugexporter"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	batchprocessor "go.opentelemetry.io/collector/processor/batchprocessor"
	"go.opentelemetry.io/collector/receiver"

	leaderelector "github.com/kyma-project/opentelemetry-collector-components/extension/leaderelector"
	dummyreceiver "github.com/kyma-project/opentelemetry-collector-components/receiver/dummyreceiver"
	kymastatsreceiver "github.com/kyma-project/opentelemetry-collector-components/receiver/kymastatsreceiver"
	singletonreceivercreator "github.com/kyma-project/opentelemetry-collector-components/receiver/singletonreceivercreator"
)

func components() (otelcol.Factories, error) {
	var err error
	factories := otelcol.Factories{}

	factories.Extensions, err = extension.MakeFactoryMap(
		leaderelector.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.ExtensionModules = make(map[component.Type]string, len(factories.Extensions))
	factories.ExtensionModules[leaderelector.NewFactory().Type()] = "github.com/kyma-project/opentelemetry-collector-components/extension/leaderelector v0.0.1"

	factories.Receivers, err = receiver.MakeFactoryMap(
		dummyreceiver.NewFactory(),
		kymastatsreceiver.NewFactory(),
		singletonreceivercreator.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.ReceiverModules = make(map[component.Type]string, len(factories.Receivers))
	factories.ReceiverModules[dummyreceiver.NewFactory().Type()] = "github.com/kyma-project/opentelemetry-collector-components/receiver/dummyreceiver v0.0.1"
	factories.ReceiverModules[kymastatsreceiver.NewFactory().Type()] = "github.com/kyma-project/opentelemetry-collector-components/receiver/kymastatsreceiver v0.0.1"
	factories.ReceiverModules[singletonreceivercreator.NewFactory().Type()] = "github.com/kyma-project/opentelemetry-collector-components/receiver/singletonreceivercreator v0.0.1"

	factories.Exporters, err = exporter.MakeFactoryMap(
		debugexporter.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.ExporterModules = make(map[component.Type]string, len(factories.Exporters))
	factories.ExporterModules[debugexporter.NewFactory().Type()] = "go.opentelemetry.io/collector/exporter/debugexporter v0.111.0"

	factories.Processors, err = processor.MakeFactoryMap(
		batchprocessor.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.ProcessorModules = make(map[component.Type]string, len(factories.Processors))
	factories.ProcessorModules[batchprocessor.NewFactory().Type()] = "go.opentelemetry.io/collector/processor/batchprocessor v0.111.0"

	factories.Connectors, err = connector.MakeFactoryMap()
	if err != nil {
		return otelcol.Factories{}, err
	}
	factories.ConnectorModules = make(map[component.Type]string, len(factories.Connectors))

	return factories, nil
}
