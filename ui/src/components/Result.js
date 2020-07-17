import React from 'react';
import PropTypes from 'prop-types';
import '../index.css';

export default function Result({ registration }) {
  return (
    <div className="result">
      <span>{registration}</span>
    </div>
  );
}

Result.propTypes = {
  registration: PropTypes.object.isRequired, // eslint-disable-line react/forbid-prop-types
};
