import React from 'react';
import '../index.css';
import { registration as registrationProp } from '../common/global_types';

export default function Result({ registration }) {
  return (
    <div className="result-container">
      <div className="result-card">
        <div>{registration.n_number}</div>
        <div>{registration.make}</div>
        <div>{registration.Model}</div>
        <div>{registration.year_manufactured}</div>
        <div>{registration.num_seats}</div>
        <div>{registration.num_engines}</div>
      </div>
      <div className="result-card">
        <div>{registration.n_number}</div>
        <div>{registration.make}</div>
        <div>{registration.Model}</div>
        <div>{registration.year_manufactured}</div>
        <div>{registration.num_seats}</div>
        <div>{registration.num_engines}</div>
      </div>
    </div>
  );
}

Result.propTypes = {
  registration: registrationProp.isRequired,
};
