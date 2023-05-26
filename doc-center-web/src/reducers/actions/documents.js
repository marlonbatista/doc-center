import { createAsyncThunk, createSlice } from "@reduxjs/toolkit"
import FileService from "../../services/file.service"

export const fetchDocuments = createAsyncThunk("edifFile", async () => {
  const userId = localStorage.getItem("UserId")
  const responde = FileService.getAllDocument(userId)
  const documents = Response.json()
  return documents
});

const initialState = [
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
];

const documents = (state = initialState, action) => {
  switch (action.type) {
    case "ADD_DOC":
      return [
        ...state,
        {
          id: state.reduce((maxId, todo) => Math.max(todo.id, maxId), -1) + 1,
          description: action.description,
          number: action.number,
          dataOfIssue: action.fataOfIssue,
          file: action.file,
          userId: action.userId
        }
      ];

    case "DELETE_DOC":
        return state.filter(doc => doc.id !== action.id);

    case "UPD_DOC":
      return state.map(doc => doc.id === action.id ? 
        {   ...doc, 
            description: action.description,
            number: action.number,
            dataOfIssue: action.fataOfIssue,
            file: action.file
        }: doc );
    default:
        return state;
  }
}

export default documents;
// const documentsSlice = createSlice({
//     name: "documents",
//     initialState:{
//         entities: [],
//         loading: false,
//     },
//     reducers: {
//         documentAdded(state, action) {
//             state.entities.push(action.payload);
//         },
//         documentUpated(state, action) {
//             const { Id, Description, Number, DataOfIssue, File, UserId } = action.payload;
//             const existingDocument = state.entities.find((doc) => doc.Id = Id);
//             if (existingDocument) {
//                 existingDocument.Description = Description;
//                 existingDocument.Number = Number;
//                 existingDocument.DataOfIssue = DataOfIssue;
//                 existingDocument.File = File;
//                 existingDocument.UserId = UserId;
//             }
//         },
//         documentDeleted(state, action) {
//             const { id }= action.payload;
//             const existingDocument = state.entities.find((doc) => doc.Id === id);
//             if (existingDocument) {
//                 state.entities = state.entities.filter((doc) => doc.Id === id);
//             }
//         },
//     },
//     extraReducer: {
//         [fetchDocuments.pending]: (state, action) => {
//           state.loading = true;
//         },
//         [fetchDocuments.fulfilled]: (state, action) => {
//           state.loading = false;
//           state.entities = [...state.entities, ...action.payload];
//         },
//         [fetchDocuments.rejected]: (state, action) => {
//           state.loading = false;
//         },
//       },
// });

// export const { documentAdded, documentUpated, documentDeleted } = documentsSlice.actions;

// export default documentsSlice.reducer;
