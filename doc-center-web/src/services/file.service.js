import axios from "axios"
import authHeader from "./auth-header"
import ENV from "../env"

const API_URL = ENV().URL_API

const getAllDocument = (idUser) => {
  return axios.get(`${API_URL}/file/?idUser=${idUser}`, {
    headers: authHeader()
  })
}

const getDocumentById = (idDocument) => {
  return axios.get(`${API_URL}/file/?id=${idDocument}`, {
    headers: authHeader()
  })
}
const getUserBoard = () => {
  return axios.get(API_URL + "user", { headers: authHeader() })
}

const getModeratorBoard = () => {
  return axios.get(API_URL + "mod", { headers: authHeader() })
}

const getAdminBoard = () => {
  return axios.get(API_URL + "admin", { headers: authHeader() })
}

const UserService = {
  getPublicContent,
  getUserBoard,
  getModeratorBoard,
  getAdminBoard
}

export default UserService
