import { Link } from "react-router-dom";
import "./Home.css";

function Home() {
  return (
    <div className="home-container">
      <nav className="topbar">
        <div className="logo">
          <Link to="/" img="" style={{ textDecoration: "none" }}>
            <h1 className="logo-header">QuickPickTools</h1>
          </Link>
        </div>
        <ul className="nav-links">
          <li>
            <div className="dropdown">
              <Link to="/converters">Converters</Link>
              <div className="dropdown-content">
                <Link to="/pdf-to-word">PDF to Word</Link>
                <Link to="/image-converter">Image Converter</Link>
                <Link to="/video-converter">Video Converter</Link>
                <Link to="/audio-converter">Audio Converter</Link>
              </div>
            </div>
          </li>
          <li>
            <Link to="/tools">Tools</Link>
          </li>
          <li>
            <Link to="/about">About Us</Link>
          </li>
          <li className="search-bar">
            <form>
              <input type="text" placeholder="Search" />
              <button type="submit">
                <i className="fa fa-search"></i>
              </button>
            </form>
          </li>
          <li className="user-account">
            <Link to="/login">
              <i className="fa fa-user"></i>
            </Link>
          </li>
        </ul>
      </nav>
      <div className="hero-section">
        <h1>Convert Anything You Want</h1>
        <p>
          Use our powerful online tools to convert your files and documents
          instantly.
        </p>
        <Link to="/converters" className="cta-btn">
          Get Started
        </Link>
      </div>
      <div className="slider-section">
        <h2>Popular Conversion Tools</h2>
        <div className="slider">
          <div className="slide">
            <img
              src="https://via.placeholder.com/150x150.png?text=Image+1"
              alt="Slide 1"
            />
            <h3>PDF to Word</h3>
          </div>
          <div className="slide">
            <img
              src="https://via.placeholder.com/150x150.png?text=Image+2"
              alt="Slide 2"
            />
            <h3>Image Converter</h3>
          </div>
          <div className="slide">
            <img
              src="https://via.placeholder.com/150x150.png?text=Image+3"
              alt="Slide 3"
            />
            <h3>Video Converter</h3>
          </div>
          <div className="slide">
            <img
              src="https://via.placeholder.com/150x150.png?text=Image+4"
              alt="Slide 4"
            />
            <h3>Audio Converter</h3>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Home;
