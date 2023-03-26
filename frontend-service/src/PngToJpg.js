import React, { useState } from 'react';
import './App.css';

function PngToJpg() {
    const [selectedFile, setSelectedFile] = useState(null);
    const [convertedFile, setConvertedFile] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
  
    const handleFileChange = (event) => {
      const file = event.target.files[0];
      if (file && file.type === 'image/png') {
        setSelectedFile(file);
      } else {
        setError("Please select a PNG file");
      }
    };
  
    const handleConvertClick = async () => {
      if (!selectedFile) {
        setError("Please select a PNG file");
        return;
      }
  
      setLoading(true);
      setError(null);
  
      const formData = new FormData();
      formData.append('image', selectedFile);
  
      try {
        const response = await fetch("/to-jpg/png-to-jpg", {
          method: 'POST',
          body: formData,
        });
  
        const json = await response.json();
        const dataURI = `data:image/jpeg;base64,${json.image}`;
        const byteString = atob(json.image);
        const arrayBuffer = new ArrayBuffer(byteString.length);
        const uint8Array = new Uint8Array(arrayBuffer);
  
        for (let i = 0; i < byteString.length; i++) {
          uint8Array[i] = byteString.charCodeAt(i);
        }
  
        const blob = new Blob([uint8Array], { type: 'image/jpeg' });
  
        setConvertedFile(URL.createObjectURL(blob));
  
        const name = json.data.name;
  
        const link = document.createElement("a");
        link.href = URL.createObjectURL(blob);
        link.download = name;
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
  
        // Update the download link with the dynamic file name
        const downloadLink = document.getElementById("download-link");
        downloadLink.download = `${name}`;
  
      } catch (error) {
        setError("An error occurred during conversion. Please try again.");
      } finally {
        setLoading(false);
      }
    };
  
    return (
      <div className="App">
        <header className="App-header">
          <h1>PNG to JPG Converter</h1>
        </header>
        <main>
          <h2>Select a PNG file to convert to JPG</h2>
          <div className="file-upload-container">
            <label htmlFor="file-upload" className="file-upload-label">
              <i className="fas fa-upload"></i> Choose File
            </label>
            <input type="file" accept=".png" id="file-upload" onChange={handleFileChange} />
            <button onClick={handleConvertClick} disabled={loading} className="convert-btn">
              {loading ? "Converting..." : "Convert"}
            </button>
          </div>
          {error && <div className="error">{error}</div>}
          {convertedFile && (
            <div className="converted-file">
              <h2>Converted JPG file:</h2>
              <img src={convertedFile} alt="Converted PNG file" />
              <a href={convertedFile} download="converted-image.jpg" id="download-link" className="download-link">
                <i className="fas fa-download"></i> Download
              </a>
            </div>
          )}
        </main>
        <footer>
          <p>Created by Ashwin Shirva</p>
        </footer>
      </div>
    );
}
export default PngToJpg;