import React from "react";

export default function Saludar(props){
    console.log(props)
    console.log(props.userInfo.nombre)

    return (
        <div>
            <h2>Hola {props.userInfo.nombre} tiene {props.userInfo.edad} años</h2>
        </div>
    )
}