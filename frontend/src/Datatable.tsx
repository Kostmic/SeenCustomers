import { useState, useEffect } from "react";
import axios from "axios";

const DataTable = ({ fetchUrl, columns }) => {
  const [tableData, setTableData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const data = await axios.get(fetchUrl);
      setTableData(data.data);
    };

    fetchData();
  });

  return (
    <table className="table">
      <thead className="table-dark">
        <tr>
          {columns.map((column) => {
            return (
              <th key={column}>{column.toUpperCase().replace("_", " ")}</th>
            );
          })}
        </tr>
      </thead>
      <tbody>
        {tableData.length == 0 ? (
          <tr>
            <td colSpan={columns.length}>No Items found</td>
          </tr>
        ) : null}
        {(
          tableData.map((data) => {
            return (
              <tr key={data}>
                {columns.map((column) => {
                  return <td key={column}>{data[column]}</td>;
                })}
              </tr>
            );
          })
        )}
      </tbody>
    </table>
  );
};

export default DataTable;
