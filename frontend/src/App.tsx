import "./App.css";
import DataTable from "./Datatable";

function App() {
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
