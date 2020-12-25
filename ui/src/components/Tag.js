import React from 'react';
import PropTypes from 'prop-types';
import '../index.css';
import { filterColorMap } from '../common/global_types';

export default function Tag({ k, v, onClick }) {
  if (v === undefined || v.length === 0) {
    return null;
  }
  return (
    <span role="button" tabIndex={0} onClick={onClick} onKeyDown={onClick} className="label-parent">
      <span style={{ backgroundColor: filterColorMap[k] }} className="label">{v}</span>
    </span>
  );
}

Tag.propTypes = {
  k: PropTypes.string.isRequired,
  v: PropTypes.string.isRequired,
  onClick: PropTypes.func.isRequired,
};
