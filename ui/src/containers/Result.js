import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { clearSelectedRegistration } from '../actions';
import ResultComponent from '../components/Result';
import { registration } from '../common/global_types';

class Result extends Component {
  componentDidMount() {
    document.addEventListener('keydown', this.esc, false);
  }

  componentWillUnmount() {
    document.removeEventListener('keydown', this.esc, false);
  }

  esc = (e) => {
    const { clearRegistration } = this.props;

    switch (e.keyCode) {
      case 27: // escape key
        clearRegistration();
        break;
      default:
        break;
    }
  }

  render = () => {
    const { selectedRegistration } = this.props;
    return (
      <ResultComponent
        registration={selectedRegistration}
      />
    );
  }
}

Result.propTypes = {
  selectedRegistration: registration.isRequired,
  clearRegistration: PropTypes.func.isRequired,
};

const mapStateToProps = (state) => {
  const { selectedRegistration } = state;

  return {
    selectedRegistration,
  };
};

const mapDispatchToProps = (dispatch) => ({
  clearRegistration: () => dispatch(clearSelectedRegistration()),
});

export default connect(mapStateToProps, mapDispatchToProps)(Result);
