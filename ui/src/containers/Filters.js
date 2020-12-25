import React from 'react';
import { connect } from 'react-redux';
import Filter from './Filter';
import { searchFilters as searchFiltersProp } from '../common/global_types';

const Filters = ({ searchFilters }) => Object.keys(searchFilters).map((f) => ( // eslint-disable-line max-len
  <Filter
    key={`${searchFilters[f].key}-${searchFilters[f].value}-${searchFilters[f].when}`}
    name={searchFilters[f].key}
    value={searchFilters[f].value}
    when={f}
  />
));

Filters.propTypes = {
  searchFilters: searchFiltersProp.isRequired,
};

const mapStateToProps = (state) => {
  const { searchFilters } = state;

  return {
    searchFilters,
  };
};

export default connect(mapStateToProps)(Filters);
