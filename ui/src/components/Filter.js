import React from 'react';
import '../index.css';
import Tag from '../components/Tag.js';

export default ({ name, onKeyDown, onChange, value, input, includeConj }) => (
    <div className="test">
        {includeConj ? <span className="frank"><Tag v="AND" /></span> : null}
        <div className="tag-input">
            <span className="test-tag">
                <Tag v={name} invert={true}/>
            </span>
            <span className="test-input-parent">
                <input 
                    className="test-input" 
                    onKeyDown={onKeyDown}
                    onChange={onChange}
                    value={value}
                    ref={input}
                />
            </span>
        </div>
    </div>
)
