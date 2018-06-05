package certifier

type MemStoreProvider struct {
	byHeight	Seeds
	byHash 		map[string]Seed
}

func NewMemStoreProvider() *MemStoreProvider {
	return &MemStoreProvider {
		byHeight: Seeds{},
		byHash: map[string]Seed{},
	}
}


func (p *MemStoreProvider) StoreSeed(seed Seed) error {
	return nil
}
func (p *MemStoreProvider) GetByHeight(h int) (Seed, error) {
	return Seed{}, nil
}
func (p *MemStoreProvider) GetByHash(hash []byte) (Seed, error) {
	return Seed{}, nil
}



