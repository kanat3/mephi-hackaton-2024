import React, { useState, useRef } from 'react';
import axios from 'axios';
import styled from "@emotion/styled";

const DrugAndDropStyled = styled.div`
  border: 2px dashed #ccc;
  padding: 20px;
  border-radius: 15px;
  text-align: center;
  cursor: pointer;
`;

const Container = styled.div`
  display: flex;
  flex-direction: column;
`;

function FileUploadComponent() {
  const [file, setFile] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [mediaUrl, setMediaUrl] = useState(null);
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

    setIsLoading(true);

    const formData = new FormData();
    formData.append('file', file);

    try {
      const response = await axios.post('http://localhost:8081/video', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      });
      console.log('File uploaded successfully:', response.data);
      // Проверяем тип загруженного файла
      if (file.type === 'audio/mp3' || file.type === 'video/mp4') {
        setMediaUrl(response.data.url);
      }
    } catch (error) {
      console.error('Error uploading file:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };


  return (
    <Container>
      <DrugAndDropStyled
        onDragOver={(e) => e.preventDefault()}
        onDrop={handleDrop}
      >
        <input
          type="file"
          onChange={handleFileChange}
          style={{ display: 'none' }}
          ref={fileInputRef}
        />
        <p>Перетащите сюда файл</p>
        <button onClick={handleClick}>Выбрать файл</button>
        {file && <p>Выбранный файл: {file.name}</p>}
      </DrugAndDropStyled>
      <button onClick={handleSubmit} disabled={isLoading}>
        {isLoading ? 'Загрузка...' : 'Отправить'}
      </button>
      {mediaUrl && (
        file.type === 'audio/mp3' ? (
          <audio controls>
            <source src={mediaUrl} type="audio/mp3" />
            Your browser does not support the audio element.
          </audio>
        ) : (
          <video controls>
            <source src={mediaUrl} type="video/mp4" />
            Your browser does not support the video element.
          </video>
        )
      )}
    </Container>
  );
}

export default FileUploadComponent;
