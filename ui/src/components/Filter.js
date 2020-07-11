import React from 'react';
import '../index.css';
import Tag from '../components/Tag.js';

export default ({ name, onKeyDown, onChange, value }) => (
    <div className="test  no-border">
        <span className="test-tag">
            <Tag v={name} invert={true}/>
        </span>
        <span className="test-input-parent">
            <input 
                className="test-input" 
                onKeyDown={onKeyDown}
                onChange={onChange}
                value={value}
            />
        </span>
    </div>
)
