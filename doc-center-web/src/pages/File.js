import CardFile from "../components/CardFile";
import { connect } from "react-redux";
import AddFile from "../components/AddFile";

const File = ({documents})  => (
  <>
  {console.log("Documents", documents)}
    <h1>Meus Documentos</h1>
    {documents.map((file, index) => {
      return (
        <>
          <CardFile
            key={index}
            Description={file.Description}
            Number={file.Number}
            DataOfIssue={file.DataOfIssue}
            Id={file.Id} />
          <br />
        </>
      );
    })}
    <AddFile />
  </>
)

const mapStateToProps = state => ({ documents: state.documents});


export default connect(mapStateToProps)(File)
