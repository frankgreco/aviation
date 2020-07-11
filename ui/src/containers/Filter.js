import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import FilterComponent from '../components/Filter.js';
import { disableSearchFilter } from '../actions'

class Filter extends Component {

    static propTypes = {
        searchFilters: PropTypes.object,
        searchQuery: PropTypes.string,
        disableFilter: PropTypes.func,
        name: PropTypes.string
    }

    constructor(props) {
        super(props)

        this.state = {
            value: ''
        }
    }

    componentDidMount = () => this.input.focus()

    handleKeyDown = f => e => {
        switch(e.key) {
            case 'Backspace':
                if (this.state.value === '') {
                    this.props.disableFilter(f)
                }
        }
    }

    handleChange = e => {
        this.setState({
            value: e.target.value
        })
    }

    render = () => <FilterComponent
        key={this.props.key}
        name={this.props.name}
        onKeyDown={this.handleKeyDown(this.props.name)}
        onChange={e => {this.handleChange(e)}}
        value={this.state.value}
        input={input => { this.input = input; }}
    />
}

const mapDispatchToProps = dispatch => {
    return {
        disableFilter: f => dispatch(disableSearchFilter(f))
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
