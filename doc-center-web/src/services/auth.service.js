import axios from "axios";

const API_URL = "http://localhost:8080/api/";

const register = (username,isPerson ,cpf ,borndate, email, password) => {
  return axios.post(API_URL + "signup", {
    username,
    isPerson,
    cpf,
    borndate,
    email,
    password,
  });
};

const login = (username, password) => {
  const axios = require('axios');
  const API_URL = "http://localhost:8080/api/";

  return axios
    .post(API_URL + "signin", {
      username,
      password,
    })
    .then((response) => {
      if (response.data.accessToken) {
        localStorage.setItem("user", JSON.stringify(response.data));
      }
      return response.data;
    });
};
const logout = () => {
  localStorage.removeItem("user");
};

const getCurrentUser = () => {
  return JSON.parse(localStorage.getItem("user"));
};

const AuthService = {
  register,
  login,
  logout,
  getCurrentUser,
};

export default AuthService;
