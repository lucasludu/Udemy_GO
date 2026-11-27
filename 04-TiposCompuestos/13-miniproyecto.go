package main

import (
	"fmt"
)

// Crear dos structs

// Item -> SKU, Name, Price, Qty
// Order -> ID, Customer, Items, Meta (City, Source)

type Item struct {
	SKU   string
	Name  string
	Price int
	Qty   int
}

type Order struct {
	ID       string
	Customer string
	Items    []Item
	Meta     map[string]string
}

/*
	1. Asignar valores al order
	2. Imprimir
		- Order ID
		- Cliente
		- Acceso al primer elemento del item
		- Acceso directo a un valor en el map de Meta
		- Modificar el cupon a HOLACARADEBOLA
		- Mostrar el cupon modificado
*/

func getMiniProyectoTipoCompuesto() {
	order := Order{
		ID:       "ORDER-1001",
		Customer: "Carlitos Robles",
		Items: []Item{
			{SKU: "TechCom453", Name: "Teclado", Price: 35, Qty: 1},
			{SKU: "TechCom123", Name: "Monitor", Price: 150, Qty: 2},
		},
		Meta: map[string]string{
			"city":   "Mar del Plata",
			"source": "Facebook Ads",
		},
	}

	fmt.Println("Order ID: ", order.ID)
	fmt.Println("Customer: ", order.Customer)
	fmt.Println("Primer Item: ", order.Items[0].Name, "SKU: ", order.Items[0].SKU)
	fmt.Println("Ciudad: ", order.Meta["city"])

	order.Meta["coupon"] = "HOLACARADEBOLA"
	fmt.Println("Cupón: ", order.Meta["coupon"])
	fmt.Println(order)
}
