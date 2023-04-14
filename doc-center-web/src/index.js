import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import ErrorPage from "./error-page";
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import Authentication from './pages/Authentication';
import Database from  './pages/Database'

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },

  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
  },

  {
    path: "authentication/",
    element: <Authentication />,
  },

  {
    path: "/",
    element: <Database/>
  }
]);

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

  // If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
