# Golang

### Package
El package se debe llamar igual que el archivo, o el nombre del paquete en conjunto

### Imports
Esta palabra reservada nos ayuda a importar los packetes necesarios

```Go
    // Un solo paquetes
    import "packagename"

    // Dos o mas paquetes
    import (
        "packagename1"
        "packagename2"
    )
```
### Main
Como en la mayoria de lenguajes,necesitamos indicar al compilador donde esta el cuerpo del flujo del programa,
para esto es necesario la funcion main en el cuerpo del programa, a menos que el archivo sea parte de un paquete sin ser este anteriormente mensionado el main.go

```Go
    func main(){

    }
```
## Variables
```Go
// Declaracion formal de variable

//Sintaxis
    var NombredeVarible <TipoDeDato>
//Ejemplo
    var Soyentero int

// Declaracion por interpretacion del compilador
    NombredeVarible := <Valor_a_asignar>
```
##### Tipos de Datos
+ ** Enteros **
    - Entero int
    - Entero 16 bits int16
    - Entero 32 bits int32
    - Entero 64 bits int64

+ ** Flotantes **
    - Flotante 64 bits float64
    - Flotante 32 bits float32

+ ** Cadenas **
    - Cadena de texto string

+ **Booleans
    - Verdadero true
    - Falso false

+ ** Estructuras **

+ ** Interfaces **

+ ** Funciones **

### Conversion de Tipos
Esto es tan sencillo como lo siguiente

```Go
// Sintaxis
    <TipoDeDato>()
// Ejemplo
    int(5.22 )
```
### Condicionales

#### If-Else
```Go
// Sintaxis
    if condicion {

    }else{

    }
// Ejemplo
    var edad int = 18

    if edad >= 18{
        fmt.Print("Eres mayor de Edad")
    }else{
        fmt.Print("Aun eres joven")
    }
```
#### Switch
```Go
// Sintaxis
    switch variable_a_evaluar{
        case valor1:
            accion1
        case valor2:
            accion2
        default:
            acciondefault
    }
// Ejemplo
    var edad int = 18

    switch edad{
        case 16:
            fmt.Print(edad)
        case 17:
            fmt.Print(edad)
        case 18:
            fmt.Print(edad, "Ya eres mayor de edad")
        default:
            fmt.print("No se que onda contigo")
    }
```
## Entrada y Salida de Datos

Para esto Compruebe la parte del [Paquete fmt](#fmt) para comprobar las estructira de las funciones y para verlo en practica cheque el [ejemplo](./Ejemplos/in_out.go)

## Operadores Aritmetico y logicos
+ AND -> ``` && ```
+ OR  -> ``` || ```
+ NOT -> ``` !  ```
+ Suma -> ``` + ```
+ Resta -> ``` - ```
+ Multiplicacion -> ```  * ```
+ Division -> ``` / ```
+ Modulo -> ``` % ```
+ Suma y Resta Unitaria -> ``` ++ -- ```
+ Channel in -> ``` varchan<- ```
+ Channel out -> ``` <-varchan ```
+ Tres puntos -> ``` ...var o ...var //Depende ```

<a name="fmt"></a>
## Package fmt
El paquete fmt o format, es el que nos permitira realizar operaciones con los busses de entrada y salida de la computadora

#### Salida a Pantalla
Las funciones para escribir en pantalla con las que cuenta fmt son:

- Print: ```func Print(a ...interface{}) (n int, err error)```

- Println: ```func Println(a ...interface{}) (n int, err error)```

- Printf: ```func Printf(format string, a ...interface{}) (n int, err error)```

- Sprint: ```func Sprint(a ...interface{}) string```

- Sprintf: ```func Sprintf(format string, a ...interface{}) string```

#### Entrada por teclado
Las funciones para leer por teclado son las siguientes:

- Scan: ```func Scan(a ...interface{}) (n int, err error)```

- Scanf: ```func Scanf(format string, a ...interface{}) (n int, err error)```

- Scanln: ```func Scanln(a ...interface{}) (n int, err error)```

- Sscan: ```func Sscan(str string, a ...interface{}) (n int, err error)```