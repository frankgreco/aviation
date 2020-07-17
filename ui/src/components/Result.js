import React from 'react';
import '../index.css';
import { registration as registrationProp } from '../common/global_types';

export default function Result({ registration }) {
  return (
    <div className="result">
      <span>{registration.tailNumber}</span>
    </div>
  );
}

Result.propTypes = {
  registration: registrationProp.isRequired,
};
