package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	// Read the source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Make a destination for images
	filename := fmt.Sprintf("temp_images/%s.jpg", name)
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Insert images to destination
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}

func main() {
	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.POST("/singleUpload", singleUpload)
	e.POST("/multipleUpload", multipleUpload)
	e.GET("/getPhotoByName", GetPhotoByName)
	// e.File("/Farras Timorremboko.jpg", "temp_images/Farras Timorremboko.jpg")

	e.Logger.Fatal(e.Start(":8080"))
}

func singleUpload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	// Read the source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Make a destination for images
	filename := fmt.Sprintf("temp_images/%s.jpg", name)
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Insert images to destination
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}

func multipleUpload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	// Read the source
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["files"]

	for i, file := range files {
		//source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Make a destination for images
		filename := fmt.Sprintf("temp_images/%s%d.jpg", name, i+1)
		dst, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Insert images to destination
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	}
	return c.HTML(http.StatusOK, fmt.Sprintf("<p>%d files uploaded successfully with fields name=%s and email=%s.</p>", len(files), name, email))
}

func GetPhotoByName(c echo.Context) error {
	fileName := c.FormValue("file_name")
	name := fmt.Sprintf("temp_images/%s.jpg", fileName)
	return c.JSON(http.StatusOK, c.File(name))
}
