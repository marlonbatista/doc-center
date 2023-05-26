import * as React from "react"
import Button from "@mui/material/Button"
import Dialog from "@mui/material/Dialog"
import DialogActions from "@mui/material/DialogActions"
import DialogContent from "@mui/material/DialogContent"
import DialogContentText from "@mui/material/DialogContentText"
import DialogTitle from "@mui/material/DialogTitle"
import useMediaQuery from "@mui/material/useMediaQuery"
import TextField from "@mui/material/TextField"
import Box from "@mui/material/Box"
import Fab from "@mui/material/Fab"
import AddIcon from "@mui/icons-material/Add"
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs"
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider"
import { DatePicker } from "@mui/x-date-pickers/DatePicker"
import { useTheme } from "@mui/material/styles"
import { addDOC } from "../reducers/actions/actions"
import { useDispatch, useSelector } from "react-redux"

const AddFile = () => {
  const [open, setOpen] = React.useState(false)
  const [description, setDescription] = React.useState("")
  const [number, setNumber] = React.useState("")
  const [dataOfIssue, setDataOfIssue] = React.useState("")
  const [userId, setUserId] = React.useState("")
  const [file, setfile] = React.useState("")

  let dispatch = useDispatch()
  const docsAmount = useSelector((state) => state.documents.length)

  const theme = useTheme()
  const fullScreen = useMediaQuery(theme.breakpoints.down("md"))

  const handleDescription = (e) => setDescription(e.target.value)

  const handleNumber = (e) => setNumber(e.target.value)

  const handleDataOfIssue = (e) => {
    setDataOfIssue(e.format());
  }

  const handleFile = (e) => setfile(e.target.value)

  const handleUserId = () => {
    const userId = localStorage.getItem("UserId")
    setUserId(userId)
  }

  const handleClickOpen = () => {
    setOpen(true)
  }

  const handleClose = () => {
    setOpen(false)
  }

  const handleAdd = () => {
    handleUserId()
    if (description && number && dataOfIssue && userId) {
      dispatch(
        addDOC({
          Id: docsAmount + 1,
          Description: description,
          Number: number,
          DataOfIssue: dataOfIssue,
          File: file,
          UserId: userId
        })
      )
    }
    dispatch(addDOC(file))
    setOpen(false)
  }

  return (
    <div>
      <Box sx={{ "& > :not(style)": { m: 1 } }}>
        <Fab
          sx={{ float: "right", marginRight: "45px" }}
          color="primary"
          aria-label="add"
          onClick={handleClickOpen}
        >
          <AddIcon />
        </Fab>
      </Box>
      <Dialog
        fullScreen={fullScreen}
        open={open}
        onClose={handleClose}
        aria-labelledby="responsive-dialog-title"
      >
        <DialogTitle id="responsive-dialog-title">Novo Arquivo</DialogTitle>
        <br />
        <DialogContent>
          <DialogContentText>
            <TextField
              required
              id="outlined-required"
              label="Descrição do arquivo"
              defaultValue="Ex: RG"
              onChange={handleDescription}
            />
            <br />
            <br />
            <TextField
              required
              id="outlined-required"
              label="Númeração do arquivo"
              defaultValue="Ex: 99.999-99"
              onChange={handleNumber}
            />
            <br />
            <br />
            <label>Data de Registo</label>
            <br />
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <DatePicker onChange={handleDataOfIssue} />
            </LocalizationProvider>
            <br />
            <br />
            <Button variant="contained" component="label">
              Upload File
              <input type="file" hidden />
            </Button>
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button
            variant="contained"
            size="small"
            color="error"
            autoFocus
            onClick={handleClose}
          >
            Cancelar
          </Button>
          <Button
            variant="contained"
            size="small"
            color="success"
            onClick={handleAdd}
            disabled={(description !== "" && number !== "" && dataOfIssue !== "") ? false: true }
            autoFocus
          >
            Confirmar
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  )
}

export default AddFile

// const mapDispatchToProps = dispatch => dispatch(deleteDOC(doc));

// export default connect(mapDispatchToProps)(DeleteFile)
