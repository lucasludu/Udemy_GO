package main

import "fmt"

func getMiniproyecto() {

	// --- Ticket de compras con impuestos --- //
	const tax = 0.21 // Impuesto del 21%

	fmt.Println("=== Ticket de compras ===")

	// Nombre del cliente
	var nombreCliente string
	fmt.Printf("Ingrese el nombre del cliente: ")
	fmt.Scanln(&nombreCliente)

	// Email del cliente
	var emailCliente string
	fmt.Printf("Ingrese el email del cliente: ")
	fmt.Scanln(&emailCliente)

	// Subtotal de la compra
	var subtotal float64
	fmt.Printf("Ingrese el subtotal de la compra: ")
	fmt.Scanln(&subtotal)

	// Respuesta:
	fmt.Println("##########################################")

	fmt.Printf("Subtotal: %.2f\n", subtotal)
	fmt.Printf("Impuestos (IVA): %.2f\n", subtotal*tax)
	fmt.Printf("Total: %.2f\n", (1+tax)*subtotal)

	fmt.Println("##########################################")

	// Subtotal
	// Impuestos (IVA)
	// Total (Con impuestos incluidos)

}
