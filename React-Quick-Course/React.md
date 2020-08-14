# React-Js

React es un frame work que nos ayuda a hacer mas agil nuestra app web, debido a su virtual DOM que nos ayuda a cuando se recarga un componente de la app, no sea necesario recargar todo, sino solo los componentes hijos del que sufrio la modificacion.

Tambien gracias a JSX, que nos aporta facilidad en manejo de etiquetas HTML en la escritura del JS, para posteriormente Renderizarlo a la Pagina dinamicamente

## Crear Proyecto con React JS ##
Para crear un proyecto con react solo necesitamos colocarnos por terminal en la carpeta que alojara nuestro proyecto y escribir el siguiente codigo
```Bash
    npx create-react-app <nombreDeLaApp>
```
Esto Descarga todas las dependencias y prepara todo para trabajar

Para Iniciar la aplicacion como tal es necesario usar el siguiente comando
```Bash 
    cd <nombreDeLaApp>
# Si usas npm
    npm start
# si usas yarn 
    yarn start
```
Este comando levanta la aplicacion en el puerto 3000 en adelante

### Componentes
Todo componente debe de estar dentro del directorio src y todos los componentes desde el archivo ***deben*** hasta la funcion den empezar con mayusculas

#### Exportar Componentes

+ Al final del archivo
    ```JS
    // Opcion 1
        function <Nombre> {
            return ...
        }

        export default <Nombre>;
    ```
+ al principio de la funcion
    ```JS
    // Opcion 2
        export default function <Nombre>{
            return ...
        }
    ```
Aunque Solo puede haber un ***default*** por archivo a los demas solo se elimina la palabra default

### Importar Componentes
Para importar el default estan facil comomlo siguiente:
```JS
        import <nombre de la etiqueta> from "Path al archivo sin .js"
```
Para los demas exports
```JS
        import <nombre de la etiqueta>,{<segundo nombre>} from "Path al archivo sin .js"
        // Ejemplo
        import HolaMundo, { Adios } from "./components/HolaMundo";
```

## Props
Es la forma para pasar informacion de un componente a otro
Ya que todo surge del componente padre ***App.js***, y para pasar informacion entre componente y estructuraar la informacion usamos ***Props***

Se pueden pasar cualquier tipo de Datos, desde cadenas o enteros, hasta objetos o funciones, esto se recibe en los componentes con un parametro denominado ***props***

### Destructing
Esto es lo mismo a reasignar las propiedades pasadas por props a una constante o variable, y de esta forma no estar trabajando con los objetos y sub obvjetos
```JS
    export default function <NombreFuncion>(props){
        const {<NombrePropiendad1>,<NombrePropiendad2>}  = props
    }
```

## Hook
Esto nos permite obtener una variable estado donde podemos controlar el comportamiento u estado de nuestros componentes, por esto es muy usado para hacer mas reactiva nuestras aplicaciones
```JS
    //Ejemplo de importacion y uso
    import React, { useState } from "react";
    ...

    const [count,setCount] = useState(0)

    //sintaxis generalizada

    const [<variableConElValor>, <FuncionParaModificarElEstado>] = useSatet(<ValorInicial>)

```
