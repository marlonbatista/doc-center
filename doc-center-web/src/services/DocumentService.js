import http from "../common/http-common";

const getAll = () => {
  return http.get("/documents");
};

const get = (id) => {
  return http.get(`/documents/${id}`);
};

const create = (data) => {
  return http.post("/documents", data);
};

const update = (id, data) => {
  return http.put(`/documents/${id}`, data);
};

const remove = (id) => {
  return http.delete(`/documents/${id}`);
};

const removeAll = () => {
  return http.delete(`/documents`);
};

const findByTitle = (title) => {
  return http.get(`/documents?title=${title}`);
};

const DocumentService = {
  getAll,
  get,
  create,
  update,
  remove,
  removeAll,
  findByTitle,
};

export default DocumentService;
