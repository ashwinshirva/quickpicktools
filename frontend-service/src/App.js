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

    const response = await fetch("/to-jpg/png-to-jpg", {
      method: 'POST',
      body: formData,
    });

    const json = await response.json();
    console.log("json: ", json)
    
    const dataURI = `data:image/jpeg;base64,${json.image}`;
    const byteString = atob(json.image);
    const arrayBuffer = new ArrayBuffer(byteString.length);
    const uint8Array = new Uint8Array(arrayBuffer);
    
    for (let i = 0; i < byteString.length; i++) {
      uint8Array[i] = byteString.charCodeAt(i);
    }

    const blob = new Blob([uint8Array], { type: 'image/jpeg' });

    console.log("blob: ", blob)
    setConvertedFile(URL.createObjectURL(blob));
      
    console.log("convertedFile: ", convertedFile)

    const name = json.data.name;
    console.log(name); // prints "pngimage.png_converted.jpg"

    // download the image immediately after setting the source
    const link = document.createElement("a");
    link.href = URL.createObjectURL(blob);
    link.download = name;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

    //URL.revokeObjectURL(url);
  };

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
