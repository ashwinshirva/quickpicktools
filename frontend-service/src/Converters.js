import React from "react";
import { Link } from 'react-router-dom';

const Converters = () => {
  return (
    <div className="converters-container">
      <h1>Converters</h1>
      <div className="converter-card">
        <h2>PNG-JPG Converter</h2>
        <p>Convert PNG image to JPG image.</p>
        {/* <a href="/png-to-jpg">Try it now</a> */}
        <li><Link to="/png-to-jpg">Try it now</Link></li>
      </div>
      <div className="converter-card">
        <h2>Temperature Converter</h2>
        <p>Convert between Celsius, Fahrenheit, and Kelvin.</p>
        <a href="#">Try it now</a>
      </div>
      <div className="converter-card">
        <h2>Weight Converter</h2>
        <p>Convert between different units of weight.</p>
        <a href="#">Try it now</a>
      </div>
    </div>
  );
};

export default Converters;
