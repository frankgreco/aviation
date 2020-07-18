import React from 'react';
import PropTypes from 'prop-types';
import '../index.css';
import LabelOutlinedIcon from '@material-ui/icons/LabelOutlined';
import SearchIcon from '@material-ui/icons/Search';
import Tag from './Tag';
import { allEnabled, allDisabled } from '../utils/timer';
import { searchFilters as searchFiltersProp } from '../common/global_types';

const renderTagIfNeeded = (searchFilters, key, onClick) => (searchFilters[key].enabled ? null : (
  <Tag v={key} onClick={onClick(key)} invert />
));

export default function Tray({ onClick, searchFilters }) {
  return allEnabled(searchFilters) ? null : (
    <div className="filters-parent">
      <div className="filters-item">
        { allDisabled(searchFilters) ? <SearchIcon fontSize="small" /> : <LabelOutlinedIcon fontSize="small" /> }
      </div>
      <div className="filters">
        {renderTagIfNeeded(searchFilters, 'tail number', onClick)}
        {renderTagIfNeeded(searchFilters, 'make', onClick)}
        {renderTagIfNeeded(searchFilters, 'model', onClick)}
        {renderTagIfNeeded(searchFilters, 'airline', onClick)}
      </div>
      <div className="filters-item" />
    </div>
  );
}

Tray.propTypes = {
  onClick: PropTypes.func.isRequired,
  searchFilters: searchFiltersProp.isRequired,
};
