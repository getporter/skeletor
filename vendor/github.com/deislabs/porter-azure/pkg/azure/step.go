package azure

type Step struct {
	Description string        `yaml:"description"`
	Outputs     []AzureOutput `yaml:"outputs"`
}

type AzureOutput struct {
	Name string `yaml:"name"`
	Key  string `yaml:"key"`
}
