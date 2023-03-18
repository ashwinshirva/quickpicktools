import React, { useState } from 'react';
import './App.css';

function App() {
  const [selectedFile, setSelectedFile] = useState(null);
  const [convertedFile, setConvertedFile] = useState(null);

  const handleFileChange = (event) => {
    const file = event.target.files[0];
    console.log("Read File: ", file)
    console.dir("Read File Dir: ", file)
    if (file && file.type === 'image/png') {
      console.log("Calling setSelectedFile...")
      setSelectedFile(file);
    }
  };

  const handleConvertClick = async () => {
    if (!selectedFile) return;
    
    const formData = new FormData();
    console.log("selectedFile: ", formData)
    console.dir("selectedFile dir: ", formData)
    formData.append('image', selectedFile);
    console.log("FormData: ", Array.from(formData.entries()))

    const response = fetch('/to-jpg/png-to-jpg', {
      method: 'POST',
      body: formData
    })
      .then(response => {
        if (!response.ok) {
          console.log("PNG to JPG conversion error...")
          console.log("PNG to JPG Response: ", response)
          throw new Error('Failed to upload image');
        }
        // Handle successful response
        const blob = response.blob();
        const url = URL.createObjectURL(blob);

        setConvertedFile(url);
      })
      .catch(error => {
        // Handle error
      });
  };
//);
    //======================
  //};

  return (
    <div className="App">
      <header className="App-header">
        <h1>PNG to JPG Converter</h1>
      </header>
      <main>
        <h2>Select a PNG file to convert to JPG</h2>
        <input type="file" accept=".png" onChange={handleFileChange} />
        <button onClick={handleConvertClick}>Convert</button>
        {convertedFile && (
          <div>
            <h2>Converted JPG file:</h2>
            <img src={convertedFile} alt="Converted PNG file" />
          </div>
        )}
      </main>
    </div>
  );
}

export default App;
