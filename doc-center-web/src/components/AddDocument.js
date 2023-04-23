import React, { useState } from "react";
import DocumentDataService from "../services/DocumentService";

const AddDocument = () => {
  const initialDocumentState = {
    id: null,
    title: "",
    description: "",
    published: false
  };
  const [tutorial, setDocument] = useState(initialDocumentState);
  const [submitted, setSubmitted] = useState(false);

  const handleInputChange = event => {
    const { name, value } = event.target;
    setDocument({ ...tutorial, [name]: value });
  };

  const saveDocument = () => {
    var data = {
      title: tutorial.title,
      description: tutorial.description
    };

    DocumentDataService.create(data)
      .then(response => {
        setDocument({
          id: response.data.id,
          title: response.data.title,
          description: response.data.description,
          published: response.data.published
        });
        setSubmitted(true);
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  };

  const newDocument = () => {
    setDocument(initialDocumentState);
    setSubmitted(false);
  };

  return (
    <div className="submit-form">
      {submitted ? (
        <div>
          <h4>You submitted successfully!</h4>
          <button className="btn btn-success" onClick={newDocument}>
            Add
          </button>
        </div>
      ) : (
        <div>
          <div className="form-group">
            <label htmlFor="title">Title</label>
            <input
              type="text"
              className="form-control"
              id="title"
              required
              value={tutorial.title}
              onChange={handleInputChange}
              name="title"
            />
          </div>

          <div className="form-group">
            <label htmlFor="description" >Description</label>
            <input
              type="text"
              className="form-control"
              id="description"
              required
              value={tutorial.description}
              onChange={handleInputChange}
              name="description"
            />
          </div>

          <button onClick={saveDocument} className="btn btn-success">
            Submit
          </button>
        </div>
      )}
    </div>
  );
};

export default AddDocument;