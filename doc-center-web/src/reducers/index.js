import { combineReducers } from "@reduxjs/toolkit";

import documents from './documents'

const rootReducer = combineReducers({
    documents
})

export default rootReducer;