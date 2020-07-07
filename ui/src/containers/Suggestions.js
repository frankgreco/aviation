import React, { Component } from 'react';
import Registration from '../components/Registration.js';
import PropTypes from 'prop-types'
import { connect } from 'react-redux'

class Suggestions extends Component {

    static propTypes = {
        registrations: PropTypes.array,
        isFetching: PropTypes.bool
    }

    shouldComponentUpdate = props => !props.isFetching

    // use https://www.npmjs.com/package/react-infinite-scroll-component 
    render = () => this.props.registrations === undefined ? null : (this.props.registrations.slice(0, 10).map((r, i) => <Registration key={i} registration={r}/>))
}

const mapStateToProps = state => {
    const { searchQuery, registrationsByQuery } = state
    const { isFetching, items: registrations } = registrationsByQuery[searchQuery] || {
        isFetching: searchQuery !== '',
        items: []
    }
    return {
        registrations,
        isFetching,
    }
}

export default connect(mapStateToProps)(Suggestions)
