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

    render = () => this.props.registrations === undefined ? null : (this.props.registrations.map((r, i) => <Registration key={i} registration={r}/>))
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
