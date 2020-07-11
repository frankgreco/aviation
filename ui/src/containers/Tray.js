import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { enableSearchFilter } from '../actions'
import TrayComponent from '../components/Tray.js';

class Tray extends Component {

    static propTypes = {
        searchFilter: PropTypes.func,
        searchFilters: PropTypes.object
    }

    handleClick = f => () => this.props.searchFilter(f)

    render = () => <TrayComponent 
        onClick={this.handleClick} 
        searchFilters={this.props.searchFilters}
    />
}

const mapStateToProps = state => {
    const { searchFilters } = state

    return {
        searchFilters
    }
}

const mapDispatchToProps = dispatch => {
    return {
        searchFilter: f => dispatch(enableSearchFilter(f))
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Tray)
