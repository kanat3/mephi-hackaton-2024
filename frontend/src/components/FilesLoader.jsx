import React, { useState, useRef } from 'react';
import axios, { AxiosError } from 'axios';
import styled from "@emotion/styled";
import { ErrorComponent } from './ErrorPage';
import logo from "../assets/logo.png";
import drugAndDrop from "../assets/dragAndDrop.png";

const Wrapper = styled.div`
  border: none;
  padding: 30px;
  border-radius: 15px;
  text-align: center;
  width: 50%;
  width: 700px;
  background: white;
`;

const DrugAndDrop = styled.div`
  border: 2px dashed gray;
  border-radius: 10px;
  padding: 20px;
  cursor: pointer;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
`;

const Container = styled.div`
  display: flex;
  flex-direction: column;
  gap: 30px;
`;

function FileUploadComponent() {
  const [file, setFile] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [mediaUrl, setMediaUrl] = useState(null);
  const [errorMessage, setErrorMessage] = useState(null);
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
      alert('Выберите файл');
      return;
    }

    setIsLoading(true);

    const formData = new FormData();
    formData.append('file', file);

    try {
      const response = await axios.post('http://localhost:8081/video', formData, {
        headers: {
          'Content-Type': 'multipart/form-data'
        },
        responseType: 'blob'
      });
      console.log('File uploaded successfully:', response.data);

      const mediaBlobUrl = URL.createObjectURL(response.data);
      console.log(mediaBlobUrl);

      // if (file.type === 'audio/mp3' || file.type === 'video/mp4') {
        console.log("change");
        setMediaUrl(mediaBlobUrl);
      // }
    } catch (error) {
      console.error('Error uploading file:', error);
      setErrorMessage(error);
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
    !errorMessage ?
      <Container>
        <Wrapper
          onDragOver={(e) => e.preventDefault()}
          onDrop={handleDrop}
        >
          <img src={logo} width={300} />
          <p style={{ fontSize: "18px" }}>Сервис для определения эмоций людей в видеоконференциях</p>

          <DrugAndDrop>
            <input
              type="file"
              onChange={handleFileChange}
              style={{ display: 'none' }}
              ref={fileInputRef}
            />
            <img src={drugAndDrop} width={60} />
            <span>Перетащите сюда файл</span>
            <button onClick={handleClick}>Выбрать файл</button>
            {file && <span>Выбранный файл: {file.name}</span>}
          </DrugAndDrop>

        </Wrapper>
        <button onClick={handleSubmit} disabled={isLoading}>
          {isLoading ? 'Загрузка...' : 'Отправить'}
        </button>
        {mediaUrl && (
          file.type === 'audio/mp3' ? (
            <audio controls autoPlay>
              <source src={mediaUrl} type="audio/mp3"/>
            </audio>
          ) : (
            <video controls autoPlay>
              <source src={mediaUrl} type="video/mp4" />
            </video>
          )
        )}
      </Container> : <ErrorComponent errorObject={errorMessage} />
  );
}

export default FileUploadComponent;
