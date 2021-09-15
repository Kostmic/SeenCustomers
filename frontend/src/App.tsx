import React, { useEffect } from "react";
import logo from "./logo.svg";
import "./App.css";
import axios from "axios";
import DataTable from "./Datatable";

function App() {
  useEffect(() => {
    const url = "http://localhost:8080/users";
    fetch(url)
      .then((response) => response.json())
      .then((data) => console.log(data));
  });
  return (
    <div className="App">
      <header className="App-header">
        <DataTable
          fetchUrl="http://localhost:8080/users"
          columns={["id", "email", "phone", "views"]}
        />
      </header>
    </div>
  );
}

export default App;
