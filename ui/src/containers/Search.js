import React, { Component } from 'react';
import PropTypes from 'prop-types'
import { connect } from 'react-redux'
import { fetchRegistrationsIfNeeded, searchQuery } from '../actions'
import { reset } from '../utils/timer';
import SearchComponent from '../components/Search.js';

class Search extends Component {

    static propTypes = {
        query: PropTypes.string,
        searchQuery: PropTypes.func,
        fetchRegistrations: PropTypes.func,
        isFetching: PropTypes.bool,
        registrations: PropTypes.array
    }

    // componentDidMount = () => {
    //     this.input.focus();
    // }

    shouldFetchRegistrations = q => {
        switch(q.length) {
            case 0:
                return false
            case 1:
                if (q.charAt(0).toUpperCase() === 'N') {
                    return false
                }
                return true
            default: 
                return true
        }
    }
    
    handleChange = e => {
        const { searchQuery, fetchRegistrations }  = this.props

        searchQuery(e.target.value)
        reset(1000, e.target.value).then((q) => {
            if (this.shouldFetchRegistrations(q)) {
                fetchRegistrations(q)
            }
        })
    }

    handleClear = () => {
        this.props.searchQuery('')
        this.input.focus()
    }
      
    render = () => (
        <SearchComponent 
            query={this.props.query}
            registrations={this.props.registrations}
            handleChange={e => {this.handleChange(e)}}
            onClick={this.handleClear}
            isFetching={this.props.isFetching}
            input={(input) => { this.input = input; }}
        />
    )
}

const mapStateToProps = state => {
    const { searchQuery, registrationsByQuery } = state
    const { isFetching, items: registrations } = registrationsByQuery[searchQuery] || {
        isFetching: false
    }

    return {
        query: state.searchQuery,
        isFetching,
        registrations
    }
}

const mapDispatchToProps = dispatch => {
    return {
      fetchRegistrations: q => dispatch(fetchRegistrationsIfNeeded(q)),
      searchQuery: q => dispatch(searchQuery(q))
    }
  }

export default connect(mapStateToProps, mapDispatchToProps)(Search)
