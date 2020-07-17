import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { enableSearchFilter } from '../actions';
import TrayComponent from '../components/Tray';

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
  searchFilters: PropTypes.object.isRequired, // eslint-disable-line react/forbid-prop-types
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
