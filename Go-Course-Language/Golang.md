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
+ ***Enteros***
    - Entero 
    ```Go 
        int
    ```
    - Entero 16 bits 
    ```Go
     int16
     ```
    - Entero 32 bits 
    ```Go
     int32
     ```
    - Entero 64 bits 
    ```Go
     int64
     
+ ***Flotantes***
    - Flotante 64 bits 
    ```Go 
    float64
    ```
    - Flotante 32 bits
    ```Go 
    float32
    ```

+ ***Cadenas***
    - Cadena de texto 
    ```Go 
    string
    ```

+ ***Booleans***
    - Verdadero 
    ```Go 
    true
    ```
    - Falso 
    ```Go 
    false
    ```

+ ***Estructuras***

    Para más información [aqui](#Estrucuras)
    ```Go
        type nameStructure struct{
            variables int
            variabbles string
            others float64
        }
    ```

+ ***Interfaces***
    ```Go
            type nameInterface interface{
                nameFunc(Arguments) ReturnValues
            }
    ```

+ ***Funciones***

    Puede Cheacar un ejemplo [aqui](./Ejemplos/function1.go)
    Y puede checar más informacion [aqui](#functions)
    ```Go
        func (r receptor) FunctionName(Argument) (Returns values{

        }
    ```

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
## Ciclos
En go unicamente existe la directiva for,pero a cambio nos ofrece una gran flexibilidad de esta directiva
### For like C
```Go
    for( i:= 0; i < 15 ; i++ ){
        // Code Block
    }
``` 
### For infinito
```Go
    for{

    }
```
### For Each
Puede Cheacar un ejemplo [aqui](./Ejemplos/function1.go)
```Go
    for ind,val := range <list or iterable>{

    }
```
### For like while
```Go
    for condition{

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
    
    Nota : Puede Checar un ejemplo [aqui](./Ejemplos/function1.go)

<a name="functions"></a>
## Funciones 
Una funcion es un segmento de codigo que tiene una tarea determinada a realizar, como podemos ver [aqui](./Ejemplos/function1.go), pero realmente en ciertos momentos del desarrollo donde esa dicha fucion solo sera realizada una unica vez o es necesario pasarla como argumento a alguna funcion, no existe razon para otorgarle un nombre, por esto se desarrollo las ***funciones anonimas***

### Funciones anonimas
***Nota importante***: las funciones en golang son ciudadanos importantes, es decir son un tipo de dato, con todo lo que esto implica

Como ya lo mencione antes esto se refiere a que estas no tienen un identificador o nombre

Ya que esto es algo complicado de explicar con palabras, acuda al [ejemplo](./Ejemplos/function2.go) para una mejor comprencion

```GO
//Sintaxis sin ejecucion
    func()(Valores de retorno){

    }
//Sintaxis con ejecucion
    func()(Valores de retorno){

    }()
```
### Closures
Estos nacieron de la existencia de funciones anonimas, esto refiere a funciones que son devuletas por otras, pero las primeras mantienen todo el contexto de las segundas, y pueden realizar operaciones con esto, independientes de todas las que puedan ser resultado de las llamadas independientes a la misma fucion original

Ya que esto es algo complicado de explicar con palabras, acuda al [ejemplo](./Ejemplos/function2.go) para una mejor comprencion


## Arrays
Vectores estaticos en tamaño que solo pueden almacenar un tipo de dato especifico

```Go
//Sintaxis
    var Nombre [Tam]TipodeDatoUnitario
// Ejemplo
    var arr [10]int
    var calif [5]int{9,10,10,9,10}
```
longitud de un arreglo
```Go
    var calif [5]int{9,10,10,9,10}

    len(calif)
```
## Slice
Su tamaño puede ser apliable si es necesario

```Go
//Sintaxis
    var Nombre []TipodeDatoUnitario
// Ejemplo
    var calif []int{9,10,10,9,10}
    calif[1:]
//sintaxis (make)
//La capacidad es el tamaño con el que creara el slice para evitar el tiempo de apliarlo cada vez y len es el espacio que tenemo disponible para control mientras no solicitemos mas
    sslice := make([]TipoDeDatoUnitario,Longitud,Capacidad)
//Ejemplo 
    elementos := make([]int,5,20)
```
### Agregar elementos al final
```Go
    var calif []int{9,10,10,9,10}

    calif = append(calif,5)
```

## Map o Diccionarios
Son objetos iterables muy parecidos a los slice solo que los indices no son numeriocos (0,1,2,3 ...), son de un tipo designado cualquiera inmutable,

Se dicen que un diccionario es un conjunto de elementos clave:valor

[Aqui](./Ejemplos/maps.go) un ejemplo de Maps

```Go
// Creacion de un map
    // Version 1
    mappa := make(map[<TipoParaLaLlave>]<TipoParaValor>)

    // Version 2
    mapppa := map[<TipoParaLaLlave>]<TipoParaElValor>

// Cargar datos
    //Directamente
    mapppa := map[<TipoParaLaLlave>]<TipoParaElValor>{
        Llave:Valor,
        Llave:Valor,
    }
//Eliminar Datos
    delete(VariableMap,<Clave aa eliminar>)
```
<a name = "Estrucuras"></a>
## Estructuras
Una estructura es una forma de ordenar valores que tendran una relacion estrecha entre si, como pueden ser las caracteristicas de un objeto como tal, por lo que en ***Go*** se dice que a pesar de no contar con un manejo de POO como tal, que este si cuenta con el, esto es gracias a el uso de Estructuras.

En el [ejemplo](./Ejemplos/Estructuras.go) podemos ver una muestra de esto los atributos metodos y herencia, como polimorfismo
    
```Go
    type nameStructure struct{
        variables int
        variabbles string
        others float64
    }
```



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