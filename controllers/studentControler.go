package controllers

import (
	"bai2/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetStudents())
}

func GetStudentById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lỗi 400, máy chủ wed không thể xử lý truy vấn"})
	}
	student := models.GetStudentById(id)
	if student == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lỗi 404, Lỗi truy vấn trình duyệt khi giao tiếp với máy chủ"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func AddStudent(c *gin.Context) {
	var newStudent models.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newStudent.Id = len(models.GetStudents()) + 1
	models.AddStudent(newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

func UpdateStudent(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)

	var updateStudent models.Student

	if err := c.ShouldBindJSON(&updateStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateStudent.Id = id
	if models.UpdateStudent(id, updateStudent) {
		c.JSON(http.StatusOK, updateStudent)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "lỗi 404, lỗi truy vấn trình duyệt khi giao tiếp với máy chủ"})
}

func DeleteStudent(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)
	if models.DeleteStudent(id) {
		c.JSON(http.StatusOK, gin.H{"mes": "xóa thành công"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "lỗi 404, lỗi truy vấn trình duyệt khi giao tiếp với máy chủ"})
}

// hàm kiểm tra xem chuỗi sa có chứa chuỗi substr không
func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

/*
func SearchStudent(c *gin.Context) {

		name := c.Query("name")
		agestr := c.Query("age")
		avgmaskstr := c.Query("minAvgmask")

		var resualt []models.Student

		age, _ := strconv.Atoi(agestr)
		avgmask, _ := strconv.ParseFloat(avgmaskstr, 64)

		for _, student := range models.GetStudents() {
			if (name == "") || contains(student.Name, name) &&
				(agestr == "") || (student.Age == age) &&
				(avgmaskstr == "") || (student.Avgmask == avgmask) {
				resualt = append(resualt, student)
			}
		}
		c.JSON(http.StatusOK, resualt)
	}
*/
func SearchStudent(c *gin.Context) {

	// Lấy các giá trị tìm kiếm từ query parameters
	name := c.Query("name")
	agestr := c.Query("age")
	avgmaskstr := c.Query("minAvgmask")

	// Khởi tạo slice chứa kết quả tìm kiếm
	var result []models.Student

	// Chuyển đổi age và avgmask từ chuỗi sang kiểu số
	var age int
	var avgmask float64
	var ageErr, avgmaskErr error

	if agestr != "" {
		age, ageErr = strconv.Atoi(agestr) // Chuyển đổi tuổi từ chuỗi sang số nguyên
	}

	if avgmaskstr != "" {
		avgmask, avgmaskErr = strconv.ParseFloat(avgmaskstr, 64) // Chuyển đổi điểm trung bình từ chuỗi sang số thực
	}

	// Lặp qua danh sách sinh viên và kiểm tra điều kiện
	for _, student := range models.GetStudents() {
		// Kiểm tra điều kiện tìm kiếm
		if (name == "" || contains(student.Name, name)) && // Tìm kiếm theo tên
			(agestr == "" || (ageErr == nil && student.Age == age)) && // Tìm kiếm theo tuổi nếu có
			(avgmaskstr == "" || (avgmaskErr == nil && student.Avgmask >= avgmask)) { // Tìm kiếm theo điểm trung bình
			result = append(result, student) // Thêm sinh viên vào kết quả nếu thỏa mãn các điều kiện
		}
	}

	// Trả về kết quả dưới dạng JSON
	c.JSON(http.StatusOK, result)
}
