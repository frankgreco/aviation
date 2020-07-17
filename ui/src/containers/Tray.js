import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { enableSearchFilter } from '../actions';
import TrayComponent from '../components/Tray';
import { searchFilters as searchFiltersProp } from '../common/global_types';

class Tray extends Component {
  handleClick = (f) => () => {
    const { searchFilter } = this.props;
    searchFilter(f);
  }

  render = () => {
    const { searchFilters } = this.props;
    return (
      <TrayComponent
        onClick={this.handleClick}
        searchFilters={searchFilters}
      />
    );
  }
}

Tray.propTypes = {
  searchFilter: PropTypes.func.isRequired,
  searchFilters: searchFiltersProp.isRequired,
};

const mapStateToProps = (state) => {
  const { searchFilters } = state;

  return {
    searchFilters,
  };
};

const mapDispatchToProps = (dispatch) => ({
  searchFilter: (f) => dispatch(enableSearchFilter(f)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Tray);
