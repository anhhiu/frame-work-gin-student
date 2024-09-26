package models

type Student struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Avgmask float64 `json:"avgmask"`
}

var students = []Student{
	{Id: 1, Name: "Ha", Age: 23, Avgmask: 8.9},
	{Id: 2, Name: "Hao", Age: 231, Avgmask: 8.9},
	{Id: 3, Name: "Ha1", Age: 231, Avgmask: 8.9},
	{Id: 4, Name: "Ha2", Age: 232, Avgmask: 8.9},
	{Id: 5, Name: "Ha3", Age: 231, Avgmask: 8.9},
	{Id: 6, Name: "Ha4", Age: 233, Avgmask: 8.9},
	{Id: 7, Name: "Ha5", Age: 232, Avgmask: 8.9},
	{Id: 8, Name: "Ha6", Age: 235, Avgmask: 8.9},
}

func AddStudent(student Student) {
	students = append(students, student)
}

func GetStudents() []Student {
	return students
}

func GetStudentById(id int) *Student {
	for _, student := range students {
		if student.Id == id {
			return &student
		}
	}
	return nil
}

func UpdateStudent(id int, updateStudent Student) bool {
	for i, student := range students {
		if student.Id == id {
			students[i] = updateStudent
			return true
		}
	}
	return false
}

func DeleteStudent(id int) bool {
	for i, student := range students {
		if student.Id == id {
			students = append(students[:i], students[i+1:]...)
			return true
		}
	}
	return false
}
