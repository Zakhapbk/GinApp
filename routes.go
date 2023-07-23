package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func MainPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"title": "Main page",
			"Flag": true},
	)
}
func List(c *gin.Context) {
	DataFDb()
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{"title": "List",
			"content": EmpList,
		},
	)
}
func Employee(c *gin.Context) {
	err, person := InfById(c.Param("name"))
	if err == nil {
		c.HTML(
			http.StatusOK,
			"employee_info.html",
			gin.H{"title": person.Name,
				"content": person.About,
				"Id":      person.Id,
			},
		)
	} else {
		c.AbortWithError(http.StatusNotFound, err)

	}
}
func AddEmployeeShow(c *gin.Context) {
	c.HTML(http.StatusOK,
		"add_form.html",
		gin.H{
			"title": "New",
		})
}
func AddEmployeePost(c *gin.Context) {
	name, pos, about := c.PostForm("Name"), c.PostForm("Position"), c.PostForm("About")
	salary, _ := strconv.Atoi(c.PostForm("Salary"))
	exp, _ := strconv.Atoi(c.PostForm("Exp"))
	DataTDb(name, pos, about, salary, exp)
	c.HTML(http.StatusOK,
		"success.html",
		gin.H{
			"title": "Success",
		},
	)
}
func DeleteEmployee(c *gin.Context) {
	err := DeleteById(c.Param("name"))
	if err == nil {
		List(c)
	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}

}
