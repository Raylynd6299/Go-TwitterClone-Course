import React from "react";

export default function Popfunct(props){
    const {name = "Anonimo",saludar}  = props
    console.log(props)

    return (
        <div>
            <button onClick={() => saludar(name)}>Saludar</button>
        </div>
    )
}