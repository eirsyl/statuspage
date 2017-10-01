package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/eirsyl/statuspage/src"
	"net/http"
	"strconv"
)

/*
 * Services API
 */

func ServiceList(c *gin.Context) {
	services, _ := c.Keys["services"].(src.Services)

	s, err := services.GetServices()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, s)
}

func ServicePost(c *gin.Context) {
	var service src.Service
	err := c.BindJSON(&service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services, _ := c.Keys["services"].(src.Services)

	err = services.InsertService(&service)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, service)
}

func ServiceGet(c *gin.Context) {
	serviceId := c.Param("id")
	id, err := strconv.Atoi(serviceId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	services, _ := c.Keys["services"].(src.Services)

	s, err := services.GetService(int64(id))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, s)
}

func ServicePatch(c *gin.Context) {
	serviceId := c.Param("id")
	id, err := strconv.Atoi(serviceId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var service src.Service
	err = c.BindJSON(&service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services, _ := c.Keys["services"].(src.Services)

	err = services.UpdateService(int64(id), &service)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, service)
}

func ServiceDelete(c *gin.Context) {
	serviceId := c.Param("id")
	id, err := strconv.Atoi(serviceId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	services, _ := c.Keys["services"].(src.Services)

	err = services.DeleteService(int64(id))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

/*
 * Incident API
 */

func IncidentList(c *gin.Context) {
	incidents, _ := c.Keys["incidents"].(src.Incidents)

	i, err := incidents.GetLatestIncidents()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, i)
}

func IncidentPost(c *gin.Context) {
	var incident src.Incident
	err := c.BindJSON(&incident)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	incidents, _ := c.Keys["incidents"].(src.Incidents)

	err = incidents.InsertIncident(&incident)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	incident.Updates = nil

	c.JSON(http.StatusCreated, incident)
}

func IncidentGet(c *gin.Context) {
	incidentId := c.Param("id")
	id, err := strconv.Atoi(incidentId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	incidents, _ := c.Keys["incidents"].(src.Incidents)

	i, err := incidents.GetIncident(int64(id))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, i)
}

func IncidentDelete(c *gin.Context) {
	incidentId := c.Param("id")
	id, err := strconv.Atoi(incidentId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	incidents, _ := c.Keys["incidents"].(src.Incidents)

	err = incidents.DeleteIncident(int64(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}