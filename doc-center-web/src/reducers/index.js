import { combineReducers } from "@reduxjs/toolkit";

import documents from './actions/documents'

const rootReducer = combineReducers({
    documents
})

export default rootReducer;