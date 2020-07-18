import React from 'react';
import '../index.css';
import Search from '../containers/Search';

export default () => (
  <div className="home">
    <div className="home-item top">
      <span className="title">AIR</span>
    </div>
    <div className="home-item middle search-container">
      <Search />
    </div>
    <div className="home-item bottom graph-container" />
  </div>
);
