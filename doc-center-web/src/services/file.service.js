import axios from "axios"
import authHeader from "./auth-header"
import ENV from "../env"

const API_URL = ENV().URL_API

const HEADER = authHeader();

const getAllDocument = (idUser) => {
  return axios.get(`${API_URL}/file/?userId=${idUser}`, {
    headers: HEADER
  })
}

const getDocumentById = (idDocument) => {
  return axios.get(`${API_URL}/file/?id=${idDocument}`, {
    headers: HEADER
  })
}
const getUserBoard = () => {
  return axios.get(API_URL + "user", { headers: HEADER })
}

const getModeratorBoard = () => {
  return axios.get(API_URL + "mod", { headers: HEADER })
}

const getAdminBoard = () => {
  return axios.get(API_URL + "admin", { headers: HEADER })
}

const postFile = (name, selectFile, ) => {
  const formData = new FormData();
    formData.append("name", name);
    formData.append("file", selectFile);
  
    axios
      .post(API_URL, formData)
      .then((res) => {
        alert("File Upload success");
      })
      .catch((err) => alert("File Upload Error"));
  // return axios.post(API_URL, file, { header: HEADER});
}

const FileService = {
  getAllDocument,
  getDocumentById,
  getUserBoard,
  getModeratorBoard,
  getAdminBoard,
  postFile
}

export default FileService;