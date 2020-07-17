import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import FilterComponent from '../components/Filter';
import { disableSearchFilter, enableSearchFilter, fetchRegistrationsIfNeeded } from '../actions';
import { reset, buildQuery } from '../utils/timer';
import { searchFilters as searchFiltersProp } from '../common/global_types';

class Filter extends Component {
    componentDidMount = () => this.input.focus()

    handleKeyDown = (f) => (e) => {
      const { searchFilters, disableFilter, name } = this.props;

      switch (e.key) {
        case 'Backspace':
          if (searchFilters[name].value === '') {
            disableFilter(f);
          }
          break;
        default:
          break;
      }
    }

    shouldFetchRegistrations = (q) => {
      switch (q.length) {
        case 0:
          return false;
        default:
          return true;
      }
    }

    handleChange = (e) => {
      const {
        fetchRegistrations,
        searchFilters,
        name,
        queryFilter,
      } = this.props;

      queryFilter(name, e.target.value);

      const query = buildQuery(searchFilters);

      // searchQuery(query)
      reset(1000, query).then((q) => {
        if (this.shouldFetchRegistrations(q)) {
          fetchRegistrations(q);
        }
      });
    }

    handleKeyPress = (e) => {
      e.target.style.width = `${(4 + (e.target.value.length + 1) * 1)}ch`;
    }

    render = () => {
      const {
        key,
        name,
        includeConj,
        searchFilters,
      } = this.props;

      return (
        <FilterComponent
          key={key}
          name={name}
          onKeyDown={this.handleKeyDown(name)}
          onChange={(e) => { this.handleChange(e); }}
          value={searchFilters[name].value}
          input={(input) => { this.input = input; }}
          includeConj={includeConj}
          onKeyPress={(e) => { this.handleKeyPress(e); }}
        />
      );
    }
}

Filter.propTypes = {
  searchFilters: searchFiltersProp.isRequired,
  fetchRegistrations: PropTypes.func.isRequired,
  disableFilter: PropTypes.func.isRequired,
  queryFilter: PropTypes.func.isRequired,
  name: PropTypes.string.isRequired,
  includeConj: PropTypes.bool.isRequired,
  key: PropTypes.string.isRequired,
};

const mapDispatchToProps = (dispatch) => ({
  disableFilter: (f) => dispatch(disableSearchFilter(f)),
  queryFilter: (f, q) => dispatch(enableSearchFilter(f, q)),
  fetchRegistrations: (q) => dispatch(fetchRegistrationsIfNeeded(q)),
});

const mapStateToProps = (state) => {
  const { searchFilters, searchQuery } = state;

  return {
    searchFilters,
    searchQuery,
  };
};

export default connect(mapStateToProps, mapDispatchToProps)(Filter);
