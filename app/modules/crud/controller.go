package crud

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/illinois1204/WOLWEB-switch/app/constants"
	"github.com/illinois1204/WOLWEB-switch/app/middleware"
	"github.com/illinois1204/WOLWEB-switch/app/repository"
	"github.com/illinois1204/WOLWEB-switch/app/service"
)

func Add(c *fiber.Ctx) error {
	_name := c.FormValue("name")
	_mac := strings.ReplaceAll(c.FormValue("mac"), ":", "-")
	_port := c.FormValue("port")
	port, err := strconv.Atoi(_port)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Oops, something went wrong")
	}

	object := repository.Device{Name: _name, Mac: _mac, Port: uint16(port)}
	index, err := repository.Write(object)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Oops, something went wrong")
	}

	repository.DeviceStorage.Add(index, object)

	c.Cookie(middleware.MakeCookie())
	return c.Status(201).Render("render/table", fiber.Map{"devices": repository.DeviceStorage.ToArray()})
}

func Update(c *fiber.Ctx) error {
	var d repository.Device
	if err := c.BodyParser(&d); err != nil {
		return c.Status(500).SendString("Oops, something went wrong")
	}

	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Oops, something went wrong")
	}

	if err = repository.DeviceStorage.Update(uint(id), d); err != nil {
		return c.Status(500).SendString("Oops, something went wrong")
	}

	c.Cookie(middleware.MakeCookie())
	return c.Status(200).Render("render/table", fiber.Map{"devices": repository.DeviceStorage.ToArray()})
}

func Remove(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Oops, something went wrong")
	}

	repository.DeviceStorage.Remove(uint(id))

	c.Cookie(middleware.MakeCookie())
	return c.Status(200).Render("render/table", fiber.Map{"devices": repository.DeviceStorage.ToArray()})
}

func Export(c *fiber.Ctx) error {
	list, _, _ := service.ListStoreFiles()
	if len(list) == 0 {
		// TODO: make browser alert or better view response
		return c.Status(204).SendString("No files to export")
	}

	var buf bytes.Buffer
	tarWriter := tar.NewWriter(&buf)
	defer tarWriter.Close()

	for _, fname := range list {
		f, err := os.Open(fmt.Sprintf("%s/%s", constants.StoreDir, fname))
		if err != nil {
			fmt.Println(err)
			return c.Status(500).SendString("Oops, something went wrong")
		}

		handledGoneExit := func() error {
			f.Close()
			fmt.Println(err)
			return c.Status(500).SendString("Oops, something went wrong")
		}

		finfo, err := f.Stat()
		if err != nil {
			return handledGoneExit()
		}

		headers := &tar.Header{
			Name:    fname,
			Mode:    int64(finfo.Mode().Perm()),
			Size:    finfo.Size(),
			ModTime: finfo.ModTime(),
		}

		if err := tarWriter.WriteHeader(headers); err != nil {
			return handledGoneExit()
		}

		if _, err := io.Copy(tarWriter, f); err != nil {
			return handledGoneExit()
		}
		
		f.Close()
	}

	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", "attachment; filename=\"devices.tar\"")
	c.Set("Content-Length", strconv.Itoa(buf.Len()))

	tarWriter.Close()
	return c.SendStream(bytes.NewReader(buf.Bytes()))
}
