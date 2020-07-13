import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Filter from './Filter.js';

class Filters extends Component {

    static propTypes = {
        searchFilters: PropTypes.object,
    }

    render = () => {
        const { searchFilters } = this.props 

        let enabledFilters = Object.keys(searchFilters).filter(k => searchFilters[k].enabled).reduce((obj, key) => {
            obj[key] = searchFilters[key]
            return obj
        }, {}) 
        
        return Object.keys(enabledFilters).map((k, i) => enabledFilters[k].enabled ? <Filter
            key={i}
            name={k}
            includeConj={i > 0}
        /> : null)
    }
}

const mapStateToProps = state => {
    const { searchFilters } = state

    return {
        searchFilters,
    }
}

export default connect(mapStateToProps)(Filters)
