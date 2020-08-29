import React from 'react'
import {Link} from "react-router-dom"
import Error404_I from "../../assets/png/error-404.png"
import Logo from "../../assets/png/logo.png"

import "./Error404.scss"

export default function Error404() {
    return (
        <div className="error404">
            <img src={Logo} alt="Twitor"/>
            <img src={Error404_I} alt="Error 404"/>
            <Link to="/">Inicio</Link>
        </div>
    )
}
