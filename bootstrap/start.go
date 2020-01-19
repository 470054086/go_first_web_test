package bootstrap

type providers struct {
	providers []fun
}
type fun func()

var Func = new(providers)

func ProviderNew() *providers {
	return &providers{
		providers: []fun{},
	};
}
func (p *providers) AddProviders(f fun) {
	p.providers = append(p.providers,f)
}
func (p *providers) Providers() []fun {
	return p.providers;
}
func (p *providers) Start()  {
	for _,val := range p.Providers() {
		val();
	}
}


