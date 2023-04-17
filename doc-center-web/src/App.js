import './styles.css'
import { useState } from 'react';

function App() {
  
  return (
   <div className="container">
    <div className="container-login">
      <div className="wrap-login">
        <form className="login-form">
          <samp className="login-form-title">Doc center</samp>
        <div className="wrap-input">
          <input
           className={email !== "" ? 'has-val input' : 'input'}
           type="email"
           value={email}
           onChange={(e) => setEmail(e.target.value)}
          />
          <span className="focus-input" data-placeholder="Email"></span>

        </div>
        <div className="wrap-input">
          <input 
          className={password !== "" ? 'has-val input' : "input"}
           type="password"
           value={password}
           onChange={(e) => setPassword(e.target.value)}
          
          />
          <span className="focus-input" data-placeholder="Password"></span>
        </div>

        <div className="container-form-btn">
          <button className="login-form-btn">Login</button>
        </div>
        <div className="text-center">
          <span className="txt1">Não possui conta?</span>
          <a className="txt2" href="#">Criar conta.</a>

        </div>
        </form>

      </div>
    </div>
   </div>
  );
}

export default App;
