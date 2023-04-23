import React, { useState, useEffect } from "react";
import DocumentDataService from "../services/DocumentService";

const Document = props => {
  const initialDocumentState = {
    id: null,
    title: "",
    description: "",
    published: false
  };
  const [currentDocument, setCurrentDocument] = useState(initialDocumentState);
  const [message, setMessage] = useState("");

  const getDocument = id => {
    DocumentDataService.get(id)
      .then(response => {
        setCurrentDocument(response.data);
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  };

  useEffect(() => {
    getDocument(props.match.params.id);
  }, [props.match.params.id]);

  const handleInputChange = event => {
    const { name, value } = event.target;
    setCurrentDocument({ ...currentDocument, [name]: value });
  };

  const updatePublished = status => {
    var data = {
      id: currentDocument.id,
      title: currentDocument.title,
      description: currentDocument.description,
      published: status
    };

    DocumentDataService.update(currentDocument.id, data)
      .then(response => {
        setCurrentDocument({ ...currentDocument, published: status });
        console.log(response.data);
        setMessage("The status was updated successfully!");
      })
      .catch(e => {
        console.log(e);
      });
  };

  const updateDocument = () => {
    DocumentDataService.update(currentDocument.id, currentDocument)
      .then(response => {
        console.log(response.data);
        setMessage("The tutorial was updated successfully!");
      })
      .catch(e => {
        console.log(e);
      });
  };

  const deleteDocument = () => {
    DocumentDataService.remove(currentDocument.id)
      .then(response => {
        console.log(response.data);
        props.history.push("/tutorials");
      })
      .catch(e => {
        console.log(e);
      });
  };

  return (
    <div>
      {currentDocument ? (
        <div className="edit-form">
          <h4>Document</h4>
          <form>
            <div className="form-group">
              <label htmlFor="title">Title</label>
              <input
                type="text"
                className="form-control"
                id="title"
                name="title"
                value={currentDocument.title}
                onChange={handleInputChange}
              />
            </div>
            <div className="form-group">
              <label htmlFor="description">Description</label>
              <input
                type="text"
                className="form-control"
                id="description"
                name="description"
                value={currentDocument.description}
                onChange={handleInputChange}
              />
            </div>

            <div className="form-group">
              <label>
                <strong>Status:</strong>
              </label>
              {currentDocument.published ? "Published" : "Pending"}
            </div>
          </form>

          {currentDocument.published ? (
            <button
              className="badge badge-primary mr-2"
              onClick={() => updatePublished(false)}
            >
              UnPublish
            </button>
          ) : (
            <button
              className="badge badge-primary mr-2"
              onClick={() => updatePublished(true)}
            >
              Publish
            </button>
          )}

          <button className="badge badge-danger mr-2" onClick={deleteDocument}>
            Delete
          </button>

          <button
            type="submit"
            className="badge badge-success"
            onClick={updateDocument}
          >
            Update
          </button>
          <p>{message}</p>
        </div>
      ) : (
        <div>
          <br />
          <p>Please click on a Document...</p>
        </div>
      )}
    </div>
  );
};

export default Document;
