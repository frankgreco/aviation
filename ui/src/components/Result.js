import React from 'react';
import AirplanemodeActiveIcon from '@material-ui/icons/AirplanemodeActive';
import '../index.css';
import { registration as registrationProp } from '../common/global_types';

export default function Result({ registration }) {
  return (
    <div className="result-container">
      <div className="result-card first-line">
        <span className="first-line-plane">
          <AirplanemodeActiveIcon fontSize="small" />
        </span>
        <span className="result-item">
          {`${registration.make} ${registration.Model}`}
        </span>
      </div>
    </div>
  );
}

Result.propTypes = {
  registration: registrationProp.isRequired,
};
