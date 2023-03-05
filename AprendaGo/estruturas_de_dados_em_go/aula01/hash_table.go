package main

import "sync"

type HashTabl struct {
	items map[int][]Pessoa
	lock  sync.RWMutex
}

func hash(nome string) (key int) {
	for _, letra := range nome {
		key = 31*key + int(letra)
	}
}

func (ht *HashTable) Put(pessoa Pessoa) int {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	key := hash(pessoa.Nome)
	ht.items[key] = append(ht.items[key], pessoa)
	return key
}
