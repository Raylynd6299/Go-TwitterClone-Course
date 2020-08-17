import React, { useState } from 'react';
import SignInSingUP from "./page/SignInSingUP";

export default function App() {

  const [user,setUser] = useState({});


  return (
    <div>
      {user ? (
        <div>
          <SignInSingUP/>
        </div>
      ) : (
        <h1>No estas Logeado</h1>
      ) }
    </div>
  )

 
}
