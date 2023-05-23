import axios from "axios"
import Env from "../env"
// const API_URL = "http://localhost:8080/"
const  API_URL = Env().URL_API

const register = (username, isPerson, cpf, borndate, email, password) => {
  return axios.post(API_URL + "signup", {
    username,
    isPerson,
    cpf,
    borndate,
    email,
    password
  })
}

const login = (email, password) => {
  console.log("Rota ", API_URL)
  return axios
    .post(API_URL + "login", {
      email: email,
      password
    })
    .then((response) => {
      if (response.data.token) {
        console.log(response.data)
        const data = response.data
        const {token, idUser } = data;
        localStorage.setItem("token", token)
        localStorage.setItem("IdUser", idUser)
      }
      return response.data
    })
}

const logout = () => {
  localStorage.removeItem("IdUser")
}

const getCurrentUser = () => {
  return JSON.parse(localStorage.getItem("IdUser"))
}

const AuthService = {
  register,
  login,
  logout,
  getCurrentUser
}

export default AuthService