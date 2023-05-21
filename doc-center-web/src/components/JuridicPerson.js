import React, { useState } from 'react';
import { Typography, Container, TextField, Button, Avatar } from '@mui/material';
import { Box } from '@mui/system';
import PersonIcon from '@mui/icons-material/Person';

const JuridicPerson = () => {
  const [formData, setFormData] = useState({
    companyName: '',
    cnpj: '',
    address: '',
    city: '',
    state: '',
    // adicione outros campos relevantes aqui
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Faça algo com os dados do formulário, como enviar para o backend
    console.log(formData);
  };

  return (
    <div style={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
      <Box sx={{ flex: 1, display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
        <Container maxWidth="sm">
          <Box
            component="form"
            sx={{
              '& > :not(style)': { mt: 2 },
            }}
            onSubmit={handleSubmit}
            noValidate
            autoComplete="off"
          >
            <Box sx={{ display: 'flex', justifyContent: 'center' }}>
              <Avatar sx={{ bgcolor: 'primary.main', m: 1, width: 96, height: 96 }}>
                <PersonIcon sx={{ fontSize: 64 }} />
              </Avatar>
            </Box>
            <TextField
              fullWidth
              label="Nome da Empresa"
              name="companyName"
              variant="outlined"
              size="small"
              value={formData.companyName}
              onChange={handleChange}
              required
            />
            <TextField
              fullWidth
              label="CNPJ"
              name="cnpj"
              variant="outlined"
              size="small"
              value={formData.cnpj}
              onChange={handleChange}
              required
            />
            <TextField
              fullWidth
              label="Endereço"
              name="address"
              variant="outlined"
              size="small"
              value={formData.address}
              onChange={handleChange}
            />
            <TextField
              fullWidth
              label="Cidade"
              name="city"
              variant="outlined"
              size="small"
              value={formData.city}
              onChange={handleChange}
            />
            <TextField
              fullWidth
              label="Estado"
              name="state"
              variant="outlined"
              size="small"
              value={formData.state}
              onChange={handleChange}
            />
            {/* Adicione outros campos relevantes aqui */}
            <Box sx={{ display: 'flex', justifyContent: 'center', marginTop: 2 }}>
              <Button type="submit" variant="contained" size="large">
                Enviar
              </Button>
            </Box>
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

export default JuridicPerson;