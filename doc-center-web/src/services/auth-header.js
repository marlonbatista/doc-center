export default function authHeader() {
  let header = new Headers();
  header.append('Access-Control-Allow-Origin', '*');
  header.append('Content-Type', 'application/json');
  const token = localStorage.getItem('token');
  if (token !== '' && token !== null) {
    header = header.append('Authorization',`Bearer ${token}`)
  }
  return header;
  // if (user && user.a) {
  //   return { Authorization: 'Bearer ' + user.accessToken }; // for Golang or Springbott
  //   // return { 'x-access-token': user.accessToken };       // for Node.js Express back-end
  // } else {
  //   return {};
  // }
}