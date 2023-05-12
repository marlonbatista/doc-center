import CardFile from "../components/CardFile/CardFile"

const FILES = [
  {
    Id: 1,
    Description: "RG",
    Number: "29.717.366-2",
    DataOfIssue: "2002-10-03",
    File: "1k2oisk2k",
    UserId: 2
  },
  {
    Id: 2,
    Description: "CPF",
    Number: "339.859.110-88",
    DataOfIssue: "2000-11-15",
    File: "1k2oisk2k",
    UserId: 2
  }
]

const File = () => {
  const documentList = []
  return (
    <>
      <h1>Meus Documentos</h1>
      {FILES.map((file) => {
        return (
          <>
            <CardFile
              Description={file.Description}
              Number={file.Number}
              DataOfIssue={file.DataOfIssue}
            />
            <br />
          </>
        )
      })}
    </>
  )
}

const getAllDocuments = () => {
  return
}

export default File
