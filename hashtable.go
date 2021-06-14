package hashtable

type Hashtable struct {
	data *data
}

func New() *Hashtable {
	table := &Hashtable{
		data: &data{},
	}

	return table
}

func (ht *Hashtable) hash(key string) int {
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

func (ht *Hashtable) Put(key string, content interface{}) {
	bind := ht.hash(key)
	ht.data.bindData(bind, key, content)
}

func (ht *Hashtable) Delete(key string) {
	bind := ht.hash(key)
	ht.data.deleteData(bind, key)
}

func (ht *Hashtable) Get(key string) interface{} {
	bind := ht.hash(key)
	return ht.data.getData(bind, key)
}

func (ht *Hashtable) GetBind(key string) int {
	return ht.hash(key)
}

type data map[int]map[string]interface{}

func (d *data) bindData(bind int, key string, content interface{}) {
	if (*d)[bind] == nil {
		(*d)[bind] = make(map[string]interface{})
	}
	(*d)[bind][key] = content
}

func (d *data) deleteData(bind int, key string) {
	if (*d)[bind] == nil {
		return
	}
	if (*d)[bind][key] != nil {
		delete((*d)[bind], key)
	}
	if len((*d)[bind]) < 1 {
		delete((*d), bind)
	}
}

func (d *data) getData(bind int, key string) interface{} {
	if (*d)[bind] == nil {
		return nil
	}
	if (*d)[bind][key] != nil {
		return (*d)[bind][key]
	}
	return nil
}
