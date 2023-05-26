import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Stack from '@mui/material/Stack';

const CardFile = (props) => {
        return (
            <Card sx={{ maxWidth: '95%' }}>
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  <b>{props.Description}</b>
                </Typography>
                <Stack direction="row" spacing={5}>
                <Typography variant="body2" color="text.secondary">
                 Número: {props.Number}
                </Typography>  
                <Typography variant="body2" color="text.secondary">
                 Data de Emissão: {props.DataOfIssue}
                </Typography>
                </Stack>
                
              </CardContent>
              <CardActions>
              <Stack direction="row" spacing={5}>
                <Button variant="contained" size="small" color="warning">Editar</Button>
                <Button variant="contained" size="small" color="primary">Anexos</Button>
                <Button variant="contained" size="small" color="error">Apagar</Button>
                <Button variant="contained" size="small" color="success">Imprimir</Button>
                </Stack>
              </CardActions>
            </Card>
          );
}

export default CardFile;