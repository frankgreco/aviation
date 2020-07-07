import React from 'react';
import '../index.css';

export default ({ k, v }) => {
    if (v === undefined || v.length === 0) {
        return null
    }
    return (
        <span className="label-parent">
            <span className={`label ${k}`}>{v}</span>
        </span>
    )
}
