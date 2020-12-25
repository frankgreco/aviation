import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import FilterComponent from '../components/Filter';
import { disableSearchFilter, enableSearchFilter, fetchRegistrationsIfNeeded } from '../actions';
import { reset, buildQuery } from '../utils/timer';
import { searchFilters as searchFiltersProp } from '../common/global_types';

class Filter extends Component {
  constructor(props) {
    super(props);
    this.state = {
      isFocused: true,
    };
  }

    componentDidMount = () => this.input.focus()

    handleKeyDown = (f) => (e) => {
      const { searchFilters, disableFilter, when } = this.props;

      switch (e.key) {
        case 'Backspace':
          if (searchFilters[when].value === '') {
            disableFilter(f, when);
          }
          break;
        default:
          break;
      }
    }

    handleChange = (e, id) => {
      const {
        fetchRegistrations,
        searchFilters,
        name,
        queryFilter,
      } = this.props;

      queryFilter(name, id, e.target.value);

      const query = buildQuery(searchFilters);

      reset(1000, query).then((q) => {
        if (q.length > 0) {
          fetchRegistrations(q);
        }
      });
    }

    handleKeyPress = (e) => {
      e.target.style.width = `${(4 + (e.target.value.length + 1) * 1)}ch`;
    }

    handleOnBlur = () => {
      this.setState({ isFocused: false });
    }

    handleOnFocus = (e) => {
      e.target.style.width = `${(4 + (e.target.value.length + 1) * 1)}ch`;
      this.setState({ isFocused: true });
    }

    render = () => {
      const {
        name,
        value,
        when,
      } = this.props;

      const {
        isFocused,
      } = this.state;

      return (
        <FilterComponent
          name={name}
          value={value}
          when={when}
          onBlur={(e) => { this.handleOnBlur(e); }}
          onFocus={(e) => { this.handleOnFocus(e); }}
          isFocused={isFocused}
          onKeyDown={this.handleKeyDown(name)}
          onChange={(e) => { this.handleChange(e, when); }}
          input={(input) => { this.input = input; }}
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
  value: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  when: PropTypes.string.isRequired,
};

const mapDispatchToProps = (dispatch) => ({
  disableFilter: (f, w) => dispatch(disableSearchFilter(f, w)),
  queryFilter: (f, w, q) => dispatch(enableSearchFilter(f, w, q)),
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
