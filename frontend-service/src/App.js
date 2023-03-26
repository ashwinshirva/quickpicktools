// import React, { useState } from 'react';
// import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
// import { BrowserRouter, Route, Switch } from 'react-router-dom';
import Home from './Home';
import PngToJpg from './PngToJpg';
import Converters from './Converters';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<Home />} />
        <Route path="/png-to-jpg" element={<PngToJpg />} />
        <Route path="/converters" element={<Converters />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
