import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Filter from './Filter.js';

class Filters extends Component {

    static propTypes = {
        searchFilters: PropTypes.object,
    }

    render = () => (Object.keys(this.props.searchFilters).map((k, i) => this.props.searchFilters[k].enabled ? <Filter
        key={i}
        name={k}
    /> : null))
}

const mapStateToProps = state => {
    const { searchFilters } = state

    return {
        searchFilters,
    }
}

export default connect(mapStateToProps)(Filters)
