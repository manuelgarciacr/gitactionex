package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	// Obtener los últimos 3 commits
	cmd := exec.Command("git", "log", "-n", "3", "--pretty=format:%h - %an, %ar : %s")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error ejecutando git log: %v\n", err)
		os.Exit(1)
	}

	// Crear directorio log/ en la raíz del repo (no en scripts/)
	logDir := filepath.Join("..", "log")  // 👈 Cambio clave
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, 0755)
		if err != nil {
			fmt.Printf("Error creando directorio %s: %v\n", logDir, err)
			os.Exit(1)
		}
	}

	// Generar nombre de archivo
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	logFile := filepath.Join(logDir, fmt.Sprintf("commits_%s.txt", currentTime))

	// Escribir archivo
	content := fmt.Sprintf("Últimos 3 commits del repositorio:\n\n%s", string(out))
	err = os.WriteFile(logFile, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error escribiendo en %s: %v\n", logFile, err)
		os.Exit(1)
	}

	fmt.Printf("Archivo de log creado en: %s\n", logFile)
}
