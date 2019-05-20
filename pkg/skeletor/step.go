package skeletor

type Step struct {
	Description string   `yaml:"description"`
	Outputs     []Output `yaml:"outputs"`
}

type Output struct {
	Name string `yaml:"name"`
	Key  string `yaml:"key"`
}
