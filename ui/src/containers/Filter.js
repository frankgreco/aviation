import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import FilterComponent from '../components/Filter.js';
import { disableSearchFilter, enableSearchFilter } from '../actions'

class Filter extends Component {

    static propTypes = {
        searchFilters: PropTypes.object,
        searchQuery: PropTypes.string,
        disableFilter: PropTypes.func,
        queryFilter: PropTypes.func,
        name: PropTypes.string,
        includeConj: PropTypes.bool
    }

    componentDidMount = () => this.input.focus()

    handleKeyDown = f => e => {
        switch(e.key) {
            case 'Backspace':
                if (this.props.searchFilters[this.props.name].value === '') {
                    this.props.disableFilter(f)
                }
        }
    }

    handleChange = e => this.props.queryFilter(this.props.name, e.target.value)

    render = () => <FilterComponent
        key={this.props.key}
        name={this.props.name}
        onKeyDown={this.handleKeyDown(this.props.name)}
        onChange={e => {this.handleChange(e)}}
        value={this.props.searchFilters[this.props.name].value}
        input={input => { this.input = input; }}
        includeConj={this.props.includeConj}
    />
}

const mapDispatchToProps = dispatch => {
    return {
        disableFilter: f => dispatch(disableSearchFilter(f)),
        queryFilter: (f, q) => dispatch(enableSearchFilter(f, q))
    }
}

const mapStateToProps = state => {
    const { searchFilters, searchQuery } = state

    return {
        searchFilters,
        searchQuery
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Filter)
