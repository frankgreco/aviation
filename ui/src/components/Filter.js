import React from 'react';
import PropTypes from 'prop-types';
import '../index.css';
import Tag from './Tag';

export default function Filter({
  name,
  onKeyDown,
  onKeyPress,
  onChange,
  value,
  input,
  includeConj,
}) {
  return (
    <div className="test">
      {includeConj ? <span className="frank"><Tag v="AND" /></span> : null}
      <div className="tag-input">
        <span className="test-tag">
          <Tag v={name} invert />
        </span>
        <span className="test-input-parent">
          <input
            className="test-input"
            onKeyDown={onKeyDown}
            onKeyPress={onKeyPress}
            // onKeyPress doesn't work for all keys (i.e. backspace)
            onKeyUp={onKeyPress}
            onChange={onChange}
            value={value}
            ref={input}
          />
        </span>
      </div>
    </div>
  );
}

Filter.propTypes = {
  name: PropTypes.string.isRequired,
  onKeyDown: PropTypes.func.isRequired,
  onKeyPress: PropTypes.func.isRequired,
  onChange: PropTypes.func.isRequired,
  value: PropTypes.string.isRequired,
  input: PropTypes.func.isRequired,
  includeConj: PropTypes.bool.isRequired,
};
