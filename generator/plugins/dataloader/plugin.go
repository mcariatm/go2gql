package dataloader

import (
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"

	"github.com/EGT-Ukraine/go2gql/generator"
	"github.com/EGT-Ukraine/go2gql/generator/plugins/graphql"
)

const (
	PluginName            = "dataloader"
	DataLoadersConfigsKey = "data_loaders"
)

type Plugin struct {
	gqlPlugin         *graphql.Plugin
	generateCfg       *generator.GenerateConfig
	dataLoader        *DataLoader
	dataLoaderConfigs *DataLoadersConfig

	loaders map[string]LoaderModel
}

func (p *Plugin) AddLoader(loader LoaderModel) {
	p.loaders[loader.Name] = loader
}

func (p *Plugin) Prepare() error {
	return nil
}

func (p *Plugin) Init(config *generator.GenerateConfig, plugins []generator.Plugin) error {
	p.generateCfg = config

	for _, plugin := range plugins {
		if g, ok := plugin.(*graphql.Plugin); ok {
			p.gqlPlugin = g

			break
		}
	}

	if p.gqlPlugin == nil {
		return errors.New("graphql plugin was not found")
	}

	var dataLoadersConfig DataLoadersConfig

	if config.PluginsConfigs[DataLoadersConfigsKey] != nil {
		if err := mapstructure.Decode(config.PluginsConfigs[DataLoadersConfigsKey], &dataLoadersConfig); err != nil {
			return errors.Wrap(err, "failed to decode dataloaders config")
		}

		outPath, err := filepath.Abs(dataLoadersConfig.OutputPath)

		if err != nil {
			return errors.Wrapf(err, "Failed to normalize path")
		}

		dataLoadersConfig.OutputPath = outPath

		p.dataLoaderConfigs = &dataLoadersConfig
	}

	p.loaders = make(map[string]LoaderModel)

	return nil
}

func (p Plugin) Name() string {
	return PluginName
}

func (p *Plugin) validateOutputObjects(gqlFiles map[string]*graphql.TypesFile) error {
	for _, gqlFile := range gqlFiles {
		for _, outputObject := range gqlFile.OutputObjects {
			for _, dataLoaderField := range outputObject.DataLoaderFields {
				dataLoader, ok := p.loaders[dataLoaderField.DataLoaderName]

				if !ok {
					return errors.Errorf(
						"Failed to found dataloader with name %s in object %s",
						dataLoaderField.DataLoaderName,
						outputObject.GraphQLName,
					)
				}

				outputArgument := outputObject.FindFieldByName(dataLoaderField.ParentKeyFieldName)

				if outputArgument == nil {
					return errors.Errorf(
						"Field `%s` not found in `%s`",
						dataLoaderField.ParentKeyFieldName,
						outputObject.GraphQLName,
					)
				}

				if !outputArgument.GoType.Scalar {
					return errors.Errorf(
						"Field `%s` in `%s` must be scalar",
						dataLoaderField.ParentKeyFieldName,
						outputObject.GraphQLName,
					)
				}

				if dataLoader.InputGoType.ElemType.Kind != outputArgument.GoType.Kind {
					// TODO: use type casting if possible.
					return errors.New("input argument must be same type as output")
				}
			}
		}
	}

	return nil
}

func (p *Plugin) PrintInfo(info generator.Infos) {
}

func (p *Plugin) Infos() map[string]string {
	return nil
}

func (p *Plugin) Generate() error {
	if p.dataLoaderConfigs == nil {
		return nil
	}

	dataLoader, err := p.createDataLoader(p.dataLoaderConfigs, p.generateCfg.VendorPath)

	if err != nil {
		return errors.Wrap(err, "failed to process dataloader config")
	}

	p.dataLoader = dataLoader

	p.gqlPlugin.AddOutputObjectFieldRenderer(&fieldsRenderer{
		dataLoader: p.dataLoader,
	})

	if err := p.validateOutputObjects(p.gqlPlugin.Types()); err != nil {
		return errors.Wrap(err, "failed to validate graphql files")
	}

	loaderGen := NewLoaderGenerator(dataLoader)

	if err := loaderGen.GenerateDataLoaders(); err != nil {
		return errors.Wrap(err, "failed to generate data loader files")
	}

	return nil
}
