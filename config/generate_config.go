package config

type GenerateConfig struct {
	Prefix        *NamingConfig `yaml:"prefix,omitempty"`
	Suffix        *NamingConfig `yaml:"suffix,omitempty"`
	UnamedPattern string        `yaml:"unamedPattern,omitempty"`
	// Deprecated: not working because it is generated by gqlgen
	Query *bool `yaml:"query,omitempty"`
	// Deprecated: not working because it is generated by gqlgen
	Mutation            *bool   `yaml:"mutation,omitempty"`
	Client              *bool   `yaml:"client,omitempty"`
	ClientInterfaceName *string `yaml:"clientInterfaceName,omitempty"`
	OmitEmptyTypes      *bool   `yaml:"omitEmptyTypes,omitempty"`
	// Deprecated: not working because v1 is deleted. Must use ClientV2
	// if true, used client v2 in generate code
	ClientV2                   bool  `yaml:"clientV2,omitempty"`
	StructFieldsAlwaysPointers *bool `yaml:"structFieldsAlwaysPointers,omitempty"`
	OnlyUsedModels             *bool `yaml:"onlyUsedModels,omitempty"`
}

func (c *GenerateConfig) ShouldGenerateClient() bool {
	if c == nil {
		return true
	}

	if c.Client != nil && !*c.Client {
		return false
	}

	return true
}

func (c *GenerateConfig) ShouldOmitEmptyTypes() bool {
	if c == nil {
		return false
	}

	if c.OmitEmptyTypes != nil && *c.OmitEmptyTypes {
		return true
	}

	return false
}

func (c *GenerateConfig) GetClientInterfaceName() *string {
	if c == nil {
		return nil
	}

	return c.ClientInterfaceName
}

type NamingConfig struct {
	Query    string `yaml:"query,omitempty"`
	Mutation string `yaml:"mutation,omitempty"`
}
