import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Registration from '../components/Registration';
import { selectedRegistration } from '../actions';
import { registration as registrationProp } from '../common/global_types';

class Suggestions extends Component {
  shouldComponentUpdate = (props) => !props.isFetching

  handleClick = (r) => () => {
    const { selectRegistration } = this.props;
    selectRegistration(r);
  }

  render = () => {
    const { registrations } = this.props;
    return registrations === undefined ? null : (
      registrations.map((r) => (
        <Registration
          key={r.tailNumber}
          onClick={this.handleClick(r)}
          registration={r}
        />
      ))
    );
  }
}

Suggestions.propTypes = {
  registrations: PropTypes.arrayOf(registrationProp).isRequired,
  isFetching: PropTypes.bool.isRequired,
  selectRegistration: PropTypes.func.isRequired,
};

const mapStateToProps = (state) => {
  const { searchQuery, registrationsByQuery } = state;
  const { isFetching, items: registrations } = registrationsByQuery[searchQuery] || {
    isFetching: searchQuery !== '',
    items: [],
  };
  return {
    registrations,
    isFetching,
  };
};

const mapDispatchToProps = (dispatch) => ({
  selectRegistration: (q) => dispatch(selectedRegistration(q)),
});

export default connect(mapStateToProps, mapDispatchToProps)(Suggestions);
