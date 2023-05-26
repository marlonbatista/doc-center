import * as React from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useTheme } from '@mui/material/styles';
import { deleteDOC } from '../reducers/actions/actions';
// import { connect } from "react-redux";
import {useDispatch} from 'react-redux';


const DeleteFile = (props,{ DeleteDOC  }) => {
  const [open, setOpen] = React.useState(false);
  const [file, setFile] = React.useState({
    Id: 1,
    Description: "",
    Number: "",
    DataOfIssue: "",
    File: "",
    UserId: 0
});
  let dispatch = useDispatch();
  const theme = useTheme();
  const fullScreen = useMediaQuery(theme.breakpoints.down('md'));

  const handleClickOpen = () => {
    console.log('Props ',props);
    setFile(props.file);
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleDelete = () => {
    dispatch(deleteDOC(file));
    setOpen(false);
  };

  return (
    
    <div>
      <Button variant="contained" size="small" color="error" onClick={handleClickOpen}>
        Apagar
      </Button>
      <Dialog
        fullScreen={fullScreen}
        open={open}
        onClose={handleClose}
        aria-labelledby="responsive-dialog-title"
      >
        <DialogTitle id="responsive-dialog-title">
          Apagar este documento?
        </DialogTitle>
        <DialogContent>
          <DialogContentText>
            
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button variant="contained" size="small" color="error" autoFocus onClick={handleClose}>
            Cancelar
          </Button>
          <Button variant="contained" size="small" color="success" onClick={handleDelete} autoFocus>
            Confirmar
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}

export default DeleteFile;

// const mapDispatchToProps = dispatch => dispatch(deleteDOC(doc));

// export default connect(mapDispatchToProps)(DeleteFile)