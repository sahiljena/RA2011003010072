package main

var (
	trainNumbers = []string{
		"2340",
		"2341",
		"2342",
		"2343",
		"2344",
		"2345",
		"2346",
		"2347",
		"2348",
		"2349",
		"2249",
		"2248",
	}
	schedules = map[string]Train{
		"2340": {
			Name:   "Yatri Exp",
			Number: "2340",
			DepartureTime: CustomTime{
				Hours:   6,
				Minutes: 10,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 10,
				AC:      3,
			},
			Price: Price{
				Sleeper: 400,
				AC:      533,
			},
			DelayedBy: 0,
		},
		"2341": {
			Name:   "Funny Exp",
			Number: "2341",
			DepartureTime: CustomTime{
				Hours:   7,
				Minutes: 15,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 15,
				AC:      5,
			},
			Price: Price{
				Sleeper: 532,
				AC:      765,
			},
			DelayedBy: 0,
		},
		"2342": {
			Name:   "Late Latif Exp",
			Number: "2342",
			DepartureTime: CustomTime{
				Hours:   8,
				Minutes: 30,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 18,
				AC:      7,
			},
			Price: Price{
				Sleeper: 1132,
				AC:      1763,
			},
			DelayedBy: 0,
		},
		"2343": {
			Name:   "Delhi Door Hai Exp",
			Number: "2343",
			DepartureTime: CustomTime{
				Hours:   9,
				Minutes: 45,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 32,
				AC:      1,
			},
			Price: Price{
				Sleeper: 555,
				AC:      1533,
			},
			DelayedBy: 0,
		},
		"2344": {
			Name:   "Chai Wala Exp",
			Number: "2344",
			DepartureTime: CustomTime{
				Hours:   11,
				Minutes: 00,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 33,
				AC:      13,
			},
			Price: Price{
				Sleeper: 543,
				AC:      654,
			},
			DelayedBy: 10,
		},
		"2345": {
			Name:   "Pappu Exp",
			Number: "2345",
			DepartureTime: CustomTime{
				Hours:   11,
				Minutes: 23,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 4,
				AC:      4,
			},
			Price: Price{
				Sleeper: 756,
				AC:      1532,
			},
			DelayedBy: 0,
		},
		"2346": {
			Name:   "Love Exp",
			Number: "2346",
			DepartureTime: CustomTime{
				Hours:   12,
				Minutes: 03,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 10,
				AC:      3,
			},
			Price: Price{
				Sleeper: 400,
				AC:      533,
			},
			DelayedBy: 0,
		},
		"2347": {
			Name:   "Sundar Exp",
			Number: "2347",
			DepartureTime: CustomTime{
				Hours:   13,
				Minutes: 32,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 55,
				AC:      0,
			},
			Price: Price{
				Sleeper: 343,
				AC:      533,
			},
			DelayedBy: 0,
		},
		"2348": {
			Name:   "Taxi Wala Exp",
			Number: "2347",
			DepartureTime: CustomTime{
				Hours:   13,
				Minutes: 32,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 2,
				AC:      2,
			},
			Price: Price{
				Sleeper: 400,
				AC:      533,
			},
			DelayedBy: 50,
		},
		"2349": {
			Name:   "Nannu Exp",
			Number: "2349",
			DepartureTime: CustomTime{
				Hours:   13,
				Minutes: 32,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 2,
				AC:      2,
			},
			Price: Price{
				Sleeper: 440,
				AC:      587,
			},
			DelayedBy: 0,
		},
		"2249": {
			Name:   "Hawai Exp",
			Number: "2349",
			DepartureTime: CustomTime{
				Hours:   14,
				Minutes: 55,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 1,
				AC:      0,
			},
			Price: Price{
				Sleeper: 867,
				AC:      954,
			},
			DelayedBy: 0,
		},
		"2248": {
			Name:   "Test Exp",
			Number: "2348",
			DepartureTime: CustomTime{
				Hours:   15,
				Minutes: 55,
				Seconds: 0,
			},
			SeatsAvailable: Seat{
				Sleeper: 1,
				AC:      0,
			},
			Price: Price{
				Sleeper: 767,
				AC:      1054,
			},
			DelayedBy: 0,
		},
	}
)
