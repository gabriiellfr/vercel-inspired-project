// app/handlers/build_handler.go

package handlers

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func BuildHandler(c *gin.Context) {
	repo := c.Query("repo")
	projectType := c.Query("type")

	if repo == "" || projectType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameters"})
		return
	}

	cmd := exec.Command("ansible-playbook", "-i", "localhost,", "../ansible/playbook.yml", "--extra-vars", fmt.Sprintf("repo=%s type=%s", repo, projectType))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Build failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Build completed successfully"})
}
