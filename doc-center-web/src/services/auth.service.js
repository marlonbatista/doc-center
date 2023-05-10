import axios from "axios"

const API_URL = "http://localhost:8080/"

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
  return axios
    .post(API_URL + "login", {
      email: email,
      password
    })
    .then((response) => {
      if (response.data.token) {
        console.log(response.data)
        const token = JSON.stringify(response.data.token)
        console.log("Token depois :", token)
        localStorage.setItem("token", token)
      }
      return response.data
    })
}

const logout = () => {
  localStorage.removeItem("user")
}

const getCurrentUser = () => {
  return JSON.parse(localStorage.getItem("user"))
}

const AuthService = {
  register,
  login,
  logout,
  getCurrentUser
}

export default AuthService
