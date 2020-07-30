package main

import (
	"fmt"
)

type person struct {
	Name    string
	Surname string
	Age     int
	Gender  string
}

type worker struct {
	City          string
	Education     string
	MeritalStatus bool
	PersonInfo    person
}

type doctor struct {
	Experience int `json:"experience"`
	WorkerInfo worker
	Salary     int    `json:"-"`
	City       string `json:"city"`
}

type driver struct {
	CarModel   string `json:"carModel"`
	WorkerInfo worker
}

type teacher struct {
	ChildrenGroup string `json:"carModel,omitempty"`
	WorkerInfo    worker
}

type military struct {
	Injuries   bool `json:"injuries"`
	WorkerInfo worker
}

type builder struct {
	Department int `json:"department"`
	WorkerInfo worker
}

func main() {
	var doctor = doctor{
		Experience: 2,
		WorkerInfo: worker{
			City:          "Moscov",
			Education:     "high",
			MeritalStatus: false,
			PersonInfo: person{
				Name:    "Ivan",
				Surname: "Kuznetsov",
				Age:     25,
				Gender:  "male",
			},
		},
		Salary: 2000,
		City:   "Moscov",
	}

	var driver = driver{
		CarModel: "BMW X5",
		WorkerInfo: worker{
			City:          "Kyiv",
			Education:     "medium",
			MeritalStatus: true,
			PersonInfo: person{
				Name:    "Bogdan",
				Surname: "Ivanov",
				Age:     30,
				Gender:  "male",
			},
		},
	}

	var teacher = teacher{
		ChildrenGroup: "junior",
		WorkerInfo: worker{
			City:          "Odessa",
			Education:     "high",
			MeritalStatus: true,
			PersonInfo: person{
				Name:    "Anna",
				Surname: "Davydova",
				Age:     29,
				Gender:  "female",
			},
		},
	}

	military := military{
		Injuries: false,
		WorkerInfo: worker{
			City:          "Odessa",
			Education:     "low",
			MeritalStatus: false,
			PersonInfo: person{
				Name:    "Michael",
				Surname: "Sydorov",
				Age:     22,
				Gender:  "male",
			},
		},
	}

	builder := builder{
		Department: 15,
		WorkerInfo: worker{
			City:          "Poltava",
			Education:     "low",
			MeritalStatus: false,
			PersonInfo: person{
				Name:    "Oleg",
				Surname: "Popov",
				Age:     24,
				Gender:  "male",
			},
		},
	}

	fmt.Printf("%+v\n %+v\n %+v\n %+v\n %+v\n", doctor, driver, teacher, military, builder)
}
