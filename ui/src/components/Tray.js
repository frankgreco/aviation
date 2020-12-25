import React from 'react';
import PropTypes from 'prop-types';
import '../index.css';
import LabelOutlinedIcon from '@material-ui/icons/LabelOutlined';
import SearchIcon from '@material-ui/icons/Search';
import Tag from './Tag';
import { searchFilters as searchFiltersProp } from '../common/global_types';

export default function Tray({ onClick, searchFilters }) {
  return (
    <div className="filters-parent">
      <div className="filters-item">
        <div className="search-icon-parent foobar">
          { Object.keys(searchFilters).length === 0 ? <SearchIcon fontSize="small" /> : <LabelOutlinedIcon fontSize="small" /> }
        </div>
      </div>
      <div className="filters">
        <Tag k="tail number" v="tail number" onClick={onClick('tail number')} invert />
        <Tag k="make" v="make" onClick={onClick('make')} invert />
        <Tag k="model" v="model" onClick={onClick('model')} invert />
        <Tag k="airline" v="airline" onClick={onClick('airline')} invert />
      </div>
      <div className="filters-item" />
    </div>
  );
}

Tray.propTypes = {
  onClick: PropTypes.func.isRequired,
  searchFilters: searchFiltersProp.isRequired,
};
