package model

type Hello struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
}

func (hello *Hello) Query(id int, name string) (hellos []Hello, err error) {
	allHello := []Hello{
		{ 0, "vinli" },
		{ 1, "daisy" },
	}
	for _, v := range allHello {
		if id >= 0 && len(name) <= 0 {
			if v.Id == id {
				hellos = append(hellos, v)
			}
		}

		if id < 0 && len(name) > 0 {
			if v.Name == name {
				hellos = append(hellos, v)
			}
		}

		if id >= 0 && len(name) > 0 {
			if v.Name == name && v.Id == id {
				hellos = append(hellos, v)
			}
		}

		if id < 0  && len(name) <= 0 {
			hellos = append(hellos, v)
		}
	}
	return
}