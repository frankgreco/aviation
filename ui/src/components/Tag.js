import React from 'react';
import '../index.css';

export default ({ k, v, onClick, invert=false }) => {
    if (v === undefined || v.length === 0) {
        return null
    }
    return (
        <span onClick={onClick} className="label-parent">
            <span className={`label ${invert ? 'invert' : ''}`}>{v}</span>
        </span>
    )
}
