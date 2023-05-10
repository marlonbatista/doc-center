import React, { useState } from 'react';
import { AppBar, Toolbar, Typography, Container, TextField, Button } from '@mui/material';
import { Box } from '@mui/system';

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [usernameError, setUsernameError] = useState('');
  const [passwordError, setPasswordError] = useState('');
 
  const handleUsernameChange = (event) => {
    setUsername(event.target.value);
    if (event.target.value.length < 3 || event.target.value.length > 50) {
        setUsernameError('Username must be between 3 and 50 characters');
      } else {
        setUsernameError('');
      }
  };

  
  const handlePasswordChange = (event) => {
    setPassword(event.target.value);
    if (event.target.value.length < 6 || event.target.value.length > 12) {
        setPasswordError('Password must be between 3 and 12 characters');
      } else {
        setPasswordError('');
      }
    }

  const handleLoginClick = () => {
    // Verificar se os dados estão válidos antes de enviar para o backend
    if (username.length >= 3 && username.length <= 50 && password.length >= 6) {
      fetch('/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          username,
          password
        })
      })
      .then(response => response.json())
      .then(data => console.log(data))
      .catch(error => console.error(error));
    }
  };

  return (
    <div style={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ position: 'absolute', top: '18px', left: '16px'}}>
            Doc center
          </Typography>
        </Toolbar>
      </AppBar>

      <Box sx={{ flex: 1, display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
        <Container maxWidth="sm">
          <Box
            component="form"
            sx={{
              '& > :not(style)': { mt: 2 },
            }}
            noValidate
            autoComplete="off"
          >
            <TextField
  fullWidth
  id="username"
  label="Username"
  variant="outlined"
  size="small"
  value={username}
  onChange={handleUsernameChange}
  error={Boolean(usernameError)}
  helperText={usernameError}
/>
<TextField
  fullWidth
  id="password"
  label="Password"
  variant="outlined"
  size="small"
  type="password"
  value={password}
  onChange={handlePasswordChange}
  error={Boolean(passwordError)}
  helperText={passwordError}
/>

            <Button variant="contained" size="large" onClick={handleLoginClick}>
              Login
            </Button>
          </Box>
        </Container>
      </Box>

      <footer style={{ marginTop: 'auto' }}>
        <Typography variant="body2" color="text.secondary" align="center">
          © {new Date().getFullYear()}, ADS Fatec
        </Typography>
      </footer>
      
    </div>
  );
};

export default Login;
