import './styles.css'
import { useState } from 'react';

function testLogin() {
const[email, setEmail] = useState("")
  const[password, setPassword] = useState("")

  return (
   <div className="container">
    <div className="container-login">
      <div className="wrap-login">
        <form className="login-form">
          <samp className="login-form-title">Bem vindo!</samp>
        <div className="wrap-input">
          <imput
           className= {email !== "" ? 'has-val input' : 'input'}
           type="email"
          value={"email"}
          onChange={e => setEmail(e.target.value)}
          
        />
          
          <span className="focus-input" data-placeholder="Email"></span>

        </div>
        <div className="wrap-input">
          <imput 
          className= {password !== "" ? 'has-val input' : 'input'}
           type="password"
          value={"password"}
          onChange={e => setPassword(e.target.value)}
          
          />
          <span className="focus-input" data-placeholder="Password"></span>
        </div>

        <div className="container-form-btn">
          <button className="login-form-btn">Login</button>
        </div>
        <div className="text-center">
          <span className="txt1">Não possui conta?</span>
          <a className="txt2" href="">Criar conta.</a>

        </div>
        </form>

      </div>
    </div>
   </div>
  );
}

const App = () => {
  const [posts, setPosts] = useState([]);
  const Login = { 
    email: email,
    password: password
  }

  useEffect(() => {
     fetch('localhost:8080/login')
        .then((res) => res.json())
        .then((Login) => {
           console.log(Login);
           setPosts(Login);
          
        })
        .catch((err) => {
           console.log(err.message);
        });
  }, []);

  return (NULL);

};

export default testLogin