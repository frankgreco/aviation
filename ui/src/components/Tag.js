import React from 'react';
import PropTypes from 'prop-types';
import '../index.css';

export default function Tag({ v, onClick, invert = false }) {
  if (v === undefined || v.length === 0) {
    return null;
  }
  return (
    <span role="button" tabIndex={0} onClick={onClick} onKeyDown={onClick} className="label-parent">
      <span className={`label ${invert ? 'invert' : ''}`}>{v}</span>
    </span>
  );
}

Tag.propTypes = {
  v: PropTypes.string.isRequired,
  onClick: PropTypes.func.isRequired,
  invert: PropTypes.bool.isRequired,
};
