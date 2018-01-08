package routes

import (
	"github.com/eirsyl/statuspage/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
 * Services API
 */

func ServiceList(c *gin.Context) {
	services := c.Keys["services"].(pkg.Services)

	s, err := services.GetServices()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, s)
}

func ServicePost(c *gin.Context) {
	var service pkg.Service
	err := c.BindJSON(&service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services := c.Keys["services"].(pkg.Services)

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

	services := c.Keys["services"].(pkg.Services)

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

	var service pkg.Service
	err = c.BindJSON(&service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services := c.Keys["services"].(pkg.Services)

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

	services := c.Keys["services"].(pkg.Services)

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
	incidents := c.Keys["incidents"].(pkg.Incidents)

	i, err := incidents.GetLatestIncidents()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, i)
}

func IncidentPost(c *gin.Context) {
	var incident pkg.Incident
	err := c.BindJSON(&incident)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	incident.Updates = nil

	incidents := c.Keys["incidents"].(pkg.Incidents)

	err = incidents.InsertIncident(&incident)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, incident)
}

func IncidentGet(c *gin.Context) {
	incidentId := c.Param("id")
	id, err := strconv.Atoi(incidentId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	incidents := c.Keys["incidents"].(pkg.Incidents)

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

	incidents := c.Keys["incidents"].(pkg.Incidents)

	err = incidents.DeleteIncident(int64(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func IncidentUpdateList(c *gin.Context) {
	incidentId := c.Param("id")
	id, err := strconv.Atoi(incidentId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	incidents := c.Keys["incidents"].(pkg.Incidents)

	incident, err := incidents.GetIncident(int64(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, incident.Updates)
}

func IncidentUpdatePost(c *gin.Context) {
	incidentId := c.Param("id")
	id, err := strconv.Atoi(incidentId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var incidentUpdate pkg.IncidentUpdate
	err = c.BindJSON(&incidentUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	incidents := c.Keys["incidents"].(pkg.Incidents)

	err = incidents.InsertIncidentUpdate(int64(id), &incidentUpdate)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, incidentUpdate)
}

func IncidentUpdateGet(c *gin.Context) {
	incidentId := c.Param("updateId")
	id, err := strconv.Atoi(incidentId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	incidents := c.Keys["incidents"].(pkg.Incidents)

	incidentUpdate, err := incidents.GetIncidentUpdate(int64(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, incidentUpdate)
}

func IncidentUpdateDelete(c *gin.Context) {
	incidentId := c.Param("updateId")
	id, err := strconv.Atoi(incidentId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	incidents := c.Keys["incidents"].(pkg.Incidents)

	err = incidents.DeleteIncidentUpdate(int64(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
