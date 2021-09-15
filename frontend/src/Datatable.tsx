import React, { useState, useEffect } from "react";
import axios from "axios";

const SORT_ASC = "asc";

const DataTable = ({ fetchUrl, columns }) => {
  const [tableData, setTableData] = useState([]);
  const [loading, setLoading] = useState(false);
  const loaderStyle = { width: "4rem", height: "4rem", color: "red" };
  const [sortField, setSortField] = useState(columns[0]);
  const [sortOrder, setSortOrder] = useState(SORT_ASC);
  const [currentPage, setCurrentPage] = useState(1);
  const [perPage, setPerPage] = useState(10);

  useEffect(() => {
    const fetchData = async () => {
      const params = {
        sort_field: sortField,
        sort_order: sortOrder,
        per_page: perPage,
        page: currentPage,
      };

      const data = await axios.get(fetchUrl);
      setTableData(data.data);
    };

    fetchData();
  }, [sortField, sortOrder, perPage, currentPage]);

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
        {loading ? (
          <tr>
            <td colSpan={columns.length + 1}>
              <div className="spinner-border" role="status">
                <span className="visually-hidden">Loading...</span>
              </div>
            </td>
          </tr>
        ) : (
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
