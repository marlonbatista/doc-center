import CardFile from "../components/CardFile";
import { connect } from "react-redux";

const File = ({documents})  => (
  <>
  {console.log("Documents", documents)}
    <h1>Meus Documentos</h1>
    {documents.map((file) => {
      return (
        <>
          <CardFile
            Description={file.Description}
            Number={file.Number}
            DataOfIssue={file.DataOfIssue}
            IdElement={file.Id} />
          <br />
        </>
      );
    })}
  </>
)

const mapStateToProps = state => ({ documents: state.documents});


export default connect(mapStateToProps)(File)
