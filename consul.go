package main

type ConsulProvider struct {
	Watch      bool
	Endpoint   string
	Prefix     string
	Filename   string
	KvProvider *KvProvider
}

func (provider *ConsulProvider) Provide(configurationChan chan<- configMessage) error {
	provider.KvProvider = NewConsulProvider(provider)
	return provider.KvProvider.provide(configurationChan)
}
