import React from 'react';
import '../index.css';
import { FaGithub } from 'react-icons/fa';
import Search from '../containers/Search';

export default () => (
  <div className="home">
    <div className="home-item top">
      <span className="title">AIR</span>
    </div>
    <div className="home-item middle search-container">
      <Search />
    </div>
    <div className="home-item bottom graph-container">
      <span className="github">
        <a href="https://github.com/frankgreco/aviation" target="_blank" rel="noopener noreferrer">
          <FaGithub size={16} />
        </a>
      </span>
    </div>
  </div>
);
