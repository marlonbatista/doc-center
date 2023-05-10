import React, { useState, useRef } from "react";
import {
TextField,
Button,
Radio,
RadioGroup,
FormControlLabel,
FormHelperText
} from '@material-ui/core';
import AuthService from "../services/auth.service";

const Register = () => {
const form = useRef();
const checkBtn = useRef();

const [username, setUsername] = useState("");
const [cpf, setCpf] = useState("");
const [email, setEmail] = useState("");
const [password, setPassword] = useState("");
const [isPerson, setIsPerson] = useState(true);
const [successful, setSuccessful] = useState(false);
const [message, setMessage] = useState("");

const required = (value) => {
if (!value) {
return 'Este campo é obrigatório!';
}
};

const validEmail = (value) => {
if (!/^[^\s@]+@[^\s@]+.[^\s@]+$/.test(value)) {
return 'Insira um e-mail válido!';
}
};

const vusername = (value) => {
if (value.length < 3 || value.length > 20) {
return 'O nome de usuário deve ter entre 3 e 20 caracteres.';
}
};

const vpassword = (value) => {
if (value.length < 6 || value.length > 40) {
return 'A senha deve ter entre 6 e 40 caracteres.';
}
};

const onChangeUsername = (e) => {
const username = e.target.value;
setUsername(username);
};

const onChangeCpf = (e) => {
const cpf = e.target.value;
setCpf(cpf);
};

const onChangeEmail = (e) => {
const email = e.target.value;
setEmail(email);
};

const onChangePassword = (e) => {
const password = e.target.value;
setPassword(password);
};

const handleRegister = (e) => {
e.preventDefault();

setMessage("");
setSuccessful(false);

form.current.reportValidity();

if (checkBtn.current.context._errors.length === 0) {
  AuthService.register(username, cpf, email, password, isPerson).then(
    (response) => {
      setMessage(response.data.message);
      setSuccessful(true);
    },
    (error) => {
      const resMessage =
        (error.response &&
          error.response.data &&
          error.response.data.message) ||
        error.message ||
        error.toString();

      setMessage(resMessage);
      setSuccessful(false);
    }
  );
}
};

return (
  <div className="col-md-12">
    <div className="card card-container">
      <img
        src="//ssl.gstatic.com/accounts/ui/avatar_2x.png"
        alt="profile-img"
        className="profile-img-card"
      />
      <Form onSubmit={handleRegister} ref={form}>
        {!successful && (
          <div>
            <div className="form-group">
              <TextField
                fullWidth
                label="Username"
                name="username"
                value={username}
                onChange={onChangeUsername}
                variant="outlined"
                required
                onBlur={(e) => form.current.reportValidity()}
                error={form.current?.elements?.username?.validationMessage}
                helperText={
                  form.current?.elements?.username?.validationMessage
                }
                inputProps={{ maxLength: 20 }}
              />
            </div>

            <div className="form-group">
              <TextField
                fullWidth
                label="CPF"
                name="cpf"
                value={cpf}
                onChange={onChangeCpf}
                variant="outlined"
                required
                onBlur={(e) => form.current.reportValidity()}
                error={form.current?.elements?.cpf?.validationMessage}
                helperText={form.current?.elements?.cpf?.validationMessage}
              />
            </div>

            <div className="form-group">
              <Button variant="contained" color="primary" type="submit">
                Register
              </Button>
            </div>
          </div>
        )}
      </Form>
    </div>
  </div>
);
}