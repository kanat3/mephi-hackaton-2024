import { useEffect, useState } from 'react'
import FileUploadComponent from './components/FilesLoader';

function App() {

  // useEffect(() => {
  //   async function fetchData() {
  //     try {
  //       const req = await fetch("http://localhost:8081/status");
  //       const jsonData = await req.json();
  //       setData(jsonData);
  //     } catch (error) {
  //       console.error("Error fetching data:", error);
  //     }
  //   }

  //   fetchData();
  // }, []);

  return (
    <>
      <FileUploadComponent />
    </>
  );
}


export default App
