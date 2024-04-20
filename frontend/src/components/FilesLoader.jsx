import React, { useState, useRef } from 'react';
import axios from 'axios';

function FileUploadComponent() {
  const [file, setFile] = useState(null);
  const fileInputRef = useRef(null);

  const handleFileChange = (event) => {
    setFile(event.target.files[0]);
  };

  const handleDrop = (event) => {
    event.preventDefault();
    const droppedFile = event.dataTransfer.files[0];
    setFile(droppedFile);
  };

  const handleSubmit = async () => {
    if (!file) {
      console.error('Please select a file.');
      return;
    }

    const formData = new FormData();
    formData.append('file', file);

    try {
      const response = await axios.post('/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      });
      console.log('File uploaded successfully:', response.data);
    } catch (error) {
      console.error('Error uploading file:', error);
    }
  };

  const handleClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  return (
    <div
      onDragOver={(e) => e.preventDefault()}
      onDrop={handleDrop}
      style={{ border: '2px dashed #ccc', padding: '20px', borderRadius: "15px", textAlign: 'center', cursor: 'pointer' }}
    >
      <input
        type="file"
        onChange={handleFileChange}
        style={{ display: 'none' }}
        ref={fileInputRef}
      />
      <p>Drag & Drop a file here or click to upload</p>
      <button onClick={handleClick}>Select File</button>
      <button onClick={handleSubmit}>Upload</button>
      {file && <p>Selected file: {file.name}</p>}
    </div>
  );
}

export default FileUploadComponent;
