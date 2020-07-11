import React from 'react';
import Tag from './Tag.js'
import '../index.css';
import AirplanemodeActiveIcon from '@material-ui/icons/AirplanemodeActive';

export default ({ registration }) => (
    <div className="container">
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
)
