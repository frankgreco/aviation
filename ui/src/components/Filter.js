import React from 'react';
import PropTypes from 'prop-types';
import { filterColorMap } from '../common/global_types';
import '../index.css';

export default function Filter({
  onKeyDown,
  onKeyPress,
  onChange,
  onBlur,
  onFocus,
  name,
  value,
  when,
  input,
  isFocused,
}) {
  return (
    <span className="test">
      <span style={{ backgroundColor: filterColorMap[name] }} className={!isFocused ? 'testing blurred' : `testing focused ${name === 'tail number' ? 'n-number' : name}`}>
        <span className="test-input-parent">
          <input
            id={when} // might as well use the time the filter was created as the unique id
            className="test-input"
            onKeyDown={onKeyDown}
            onKeyPress={onKeyPress}
            onKeyUp={onKeyPress} // onKeyPress doesn't work for all keys (i.e. backspace)
            onChange={onChange}
            onBlur={onBlur} // same as onFocusOut
            onFocus={onFocus}
            value={value}
            ref={input}
          />
        </span>
      </span>
    </span>
  );
}

Filter.propTypes = {
  onKeyDown: PropTypes.func.isRequired,
  onBlur: PropTypes.func.isRequired,
  onFocus: PropTypes.func.isRequired,
  onKeyPress: PropTypes.func.isRequired,
  onChange: PropTypes.func.isRequired,
  name: PropTypes.string.isRequired,
  value: PropTypes.string.isRequired,
  when: PropTypes.string.isRequired,
  input: PropTypes.func.isRequired,
  isFocused: PropTypes.bool.isRequired,
};
