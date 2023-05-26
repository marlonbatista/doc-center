export const ADD_DOC = "ADD_DOC";
export const DELETE_DOC = "DELETE_DOC";
export const UPDATE_DOC = "UPDATE_DOC";

export function addDOC(DOC) {
	return {
		type: ADD_DOC,
		payload: DOC,
	}
};

export function deleteDOC(DOC) {
	return {
		type: DELETE_DOC,
		payload: DOC
	}
};

export function updateDOC(DOC) {
	return {
		type: UPDATE_DOC,
		payload: DOC
	}
};
