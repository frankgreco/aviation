import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import {
  searchQuery as searchQueryAction,
  disableSearchFilter,
  hideCodeView as hideCodeViewAction,
  clearSearchFilters as clearSearchFiltersAction,
} from '../actions';
import { buildQuery } from '../utils/timer';
import SearchComponent from '../components/Search';
import { registration as registrationProp, searchFilters as searchFiltersProp } from '../common/global_types';

class Search extends Component {
  hasFilters = (f) => Object.keys(f).length > 0

  handleClear = () => {
    const { searchQueryProp, clearSearchFilters } = this.props;

    searchQueryProp('');
    clearSearchFilters();
  }

  toggleCodeView = () => {
    const { toggleCodeView, hideCodeView } = this.props;
    toggleCodeView(!hideCodeView);
  }

  render = () => {
    const {
      searchFilters,
      isFetching,
      hideCodeView,
      selectedRegistration,
    } = this.props;

    return (
      <SearchComponent
        query={buildQuery(searchFilters)}
        handleChange={(e) => { this.handleChange(e); }}
        onClick={this.handleClear}
        isFetching={isFetching}
        input={(input) => { this.input = input; }}
        hasFilters={this.hasFilters(searchFilters)}
        toggleCodeView={this.toggleCodeView}
        hideCodeView={hideCodeView}
        searchFilters={searchFilters}
        selectedRegistration={selectedRegistration}
      />
    );
  }
}

Search.propTypes = {
  searchQueryProp: PropTypes.func.isRequired,
  clearSearchFilters: PropTypes.func.isRequired,
  isFetching: PropTypes.bool.isRequired,
  searchFilters: searchFiltersProp.isRequired,
  hideCodeView: PropTypes.bool.isRequired,
  toggleCodeView: PropTypes.func.isRequired,
  selectedRegistration: registrationProp.isRequired,
};

const mapStateToProps = (state) => {
  const {
    searchQuery,
    registrationsByQuery,
    searchFilters,
    hideCodeView,
    selectedRegistration,
  } = state;
  const { isFetching } = registrationsByQuery[searchQuery] || {
    isFetching: false,
  };

  return {
    isFetching,
    searchFilters,
    hideCodeView,
    selectedRegistration,
  };
};

const mapDispatchToProps = (dispatch) => ({
  searchQueryProp: (q) => dispatch(searchQueryAction(q)),
  disableFilter: (f) => dispatch(disableSearchFilter(f)),
  toggleCodeView: (f) => dispatch(hideCodeViewAction(f)),
  clearSearchFilters: (f) => dispatch(clearSearchFiltersAction(f)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Search);
