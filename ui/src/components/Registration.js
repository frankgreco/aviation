import React from 'react';
import PropTypes from 'prop-types';
import AirplanemodeActiveIcon from '@material-ui/icons/AirplanemodeActive';
import Tag from './Tag';
import '../index.css';

export default function Registration({ registration, onClick }) {
  return (
    <div role="button" tabIndex={0} onClick={onClick} onKeyDown={onClick} className="container">
      <span className="plane">
        <AirplanemodeActiveIcon fontSize="small" />
      </span>
      <div className="tags">
        <Tag k="n-number" v={registration.n_number} />
        <Tag k="make" v={registration.make} />
        <Tag k="model" v={registration.Model} />
        <Tag k="year" v={registration.year_manufactured} />
      </div>
    </div>
  );
}

Registration.propTypes = {
  registration: PropTypes.object.isRequired, // eslint-disable-line react/forbid-prop-types
  onClick: PropTypes.func.isRequired,
};
