import React, { useState, useEffect } from 'react';
import logo from './logo.svg';
import './App.css';
import Saludar from "./components/Saludar"
import Popfunct from "./components/Popfunct"

function App() {

  const [stateCar, setStateCar] = useState(false);
  const[contar, setContar] = useState(0);

  const user = {
    nombre:"Raymundo Pulido Bejarano",
    edad:21,
    color:"Rojo"
  }
  useEffect(() => {
    console.log("total: {contar}")
  }, [contar])

  const SaludarFn = (name) => {
    console.log(`Hola mi nombre es ${name}, ten un gran dia`)
  }
    // Esto va dentro del div se saco para ejemplos de hook
    //<Popfunct saludar={SaludarFn} name = "Ray"/>
    //<Saludar userInfo = {user}></Saludar>
  const encenderApagar = () => {
    //setStateCar(!stateCar);
    setStateCar(prevValue => !prevValue);
    setContar(contar+1)
  }
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <h3>El coche esta {stateCar ? "Encendido" : "Apagado"}</h3>
        <h4>Clicks : {contar}</h4>
        <button onClick= {encenderApagar}>Encender/Apagar</button>
      </header>
    </div>
  );
}

export default App;
