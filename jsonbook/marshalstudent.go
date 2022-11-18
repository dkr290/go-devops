package main

type StudentMarshal struct {
	StudentId     int             `json:"id"`
	LastName      string          `json:"lname"`
	MiddleInitial string          `json:"mname,omitempty"`
	FirstName     string          `json:"fname"`
	IsEnrolled    bool            `json:"enrolled,omitempty"`
	IsMarried     bool            `json:"-"`
	Courses       []CourseMarshal `json:"classes"`
}

type CourseMarshal struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

func NewStudent(stID int, lastName, Middlename, firstName string, isMarried, isEnrolled bool) StudentMarshal {
	s := StudentMarshal{
		StudentId:     stID,
		MiddleInitial: Middlename,
		FirstName:     firstName,
		LastName:      lastName,
		IsMarried:     isMarried,
		IsEnrolled:    isEnrolled,
	}

	return s
}
