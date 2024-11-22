package model

import (
	"errors"
	"os"
)

type GuiaDocente struct{
	Anio int
	Asignatura string
	
	BibliografiaBasica []Libro	    // Lista de libros básicos
	BibliografiaComplementaria []Libro  // Lista de libros complementarios
}


// Crea una nueva guía docente con datos válidos
func NuevaGuiaDocente(anio int, asignatura string, basica []Libro, complementaria []Libro) (*GuiaDocente, error){
	anioValido := EsPositivo(anio)

	// Ambos datos inválidos
	if !anioValido {
		return nil, errors.New("El año debe ser positivo")
	}

	// Datos válidos: se crea la guía docente sin devolver errores
	return &GuiaDocente{anio, asignatura, basica, complementaria}, nil
}

// NuevaGuiaDocente crea una nueva guía docente a partir de un handler de fichero.
// El handler debe estar previamente abierto para lectura y contendrá los datos necesarios 
// para inicializar la estructura. El contenido del fichero será procesado en el futuro 
// para extraer y validar la información requerida.
//
// Parámetros:
// - file: Handler del fichero que contiene la información de la guía docente.
//
// Retorno:
// - *GuiaDocente: Una nueva instancia inicializada de GuiaDocente.
// - error: Cualquier error que ocurra durante el procesamiento del fichero.
func NuevaGuiaDocente(file *os.File) (*GuiaDocente, error) {
	// Por implementar: procesar el contenido del fichero para crear la guía docente.
	return nil, nil
}
