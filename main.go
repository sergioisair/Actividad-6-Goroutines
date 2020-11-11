package main

import (
	"fmt"
	"time"
)

func main()  {
	var opc int
  bandera := false
	cont := 0 // contador de procesos
	var canal chan bool = make(chan bool)
	procesos := []Proceso {Proceso{0,-1,make(chan bool)}}
	go procesos[0].start(canal)
	canal <- false
	for opc != 4 {
		if(!bandera){
      fmt.Println("\n\nBienvenido al Administrador de Procesos")
      fmt.Println("1- Agregar Proceso")
      fmt.Println("2- Mostrar Proceso")
      fmt.Println("3- Terminar Proceso")
      fmt.Println("4- Salir")
		}
		fmt.Scan(&opc)
    switch opc {
      case 1:
        p := Proceso{0, cont, make(chan bool)}
        cont++
        go p.start(canal)
        procesos = append(procesos, p)
        fmt.Printf("AGREGADO CORRECTAMENTE \n")
        break
      case 2:
        fmt.Print("MOSTRANDO PROCESOS (ingrese 2 para No Mostrar)\n")
        mostrar := <- canal
        if mostrar {
          canal <- false
          bandera = false
        }else {
          canal <- true
          bandera = true
        }
        break
      case 3:
        fmt.Print("INGRESE ID A ELIMINAR: ")
        var idel int
        fmt.Scan(&idel)
        i := idel
        procesos[i+1].stop()
        break
      case 4:
        break
      default:
        fmt.Print("OPCION INCORRECTA")
        break
    }
	}
}

type Proceso struct  {
	i int64
	id int
	bandEnd chan bool
}

func (p *Proceso) start(canal chan bool)  {
	continuar := true
	for (continuar) {
		select {
		case <- p.bandEnd:
			continuar = false
		default:
		}
		if( <- canal ){
			if(p.id != -1){
				fmt.Printf("ID %d: %d \n", p.id, p.i)
			}	
			canal <- true
		}else{
			canal <- false
		}
		if p.id != -1 {
			p.i = p.i + 1
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Proceso) stop(){
	p.bandEnd <- true
}
