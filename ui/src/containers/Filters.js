import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Filter from './Filter';

const Filters = ({ searchFilters }) => {
  const enabledFilters = Object.keys(searchFilters)
    .filter((k) => searchFilters[k].enabled)
    .reduce((obj, key) => {
      obj[key] = searchFilters[key]; // eslint-disable-line no-param-reassign
      return obj;
    }, {});

  return Object.keys(enabledFilters).map((k, i) => (enabledFilters[k].enabled ? (
    <Filter
      key={k}
      name={k}
      includeConj={i > 0}
    />
  ) : null));
};

Filters.propTypes = {
  searchFilters: PropTypes.object.isRequired, // eslint-disable-line react/forbid-prop-types
};

const mapStateToProps = (state) => {
  const { searchFilters } = state;

  return {
    searchFilters,
  };
};

export default connect(mapStateToProps)(Filters);
