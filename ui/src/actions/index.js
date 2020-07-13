import axios from 'axios'

export const HIDE_CODE_VIEW = "HIDE_CODE_VIEW"
export function hideCodeView(value) {
    return { type: HIDE_CODE_VIEW, value }
}

export const ENABLE_SEARCH_FILTER = "ENABLE_SEARCH_FILTER"
export function enableSearchFilter(filter, query = '') {
    return { type: ENABLE_SEARCH_FILTER, filter, query }
}

export const DISABLE_SEARCH_FILTER = "DISABLE_SEARCH_FILTER"
export function disableSearchFilter(filter, query = '') {
    return { type: DISABLE_SEARCH_FILTER, filter, query }
}

export const SEARCH_QUERY = "SEARCH_QUERY"
export function searchQuery(searchQuery) {
    return { type: SEARCH_QUERY, searchQuery }
}

export const REQUEST_REGISTRATIONS = "REQUEST_REGISTRATIONS"
export function requestRegistrations(searchQuery) {
    return { type: REQUEST_REGISTRATIONS, searchQuery }
}

export const RECEIVE_REGISTRATIONS = "RECEIVE_REGISTRATIONS"
export function receiveRegistrations(searchQuery, json) {
    return {
        type: RECEIVE_REGISTRATIONS,
        searchQuery,
        registrations: json.items,
        receivedAt: Date.now()
    }
}

export function fetchRegistrations(searchQuery) {
    return dispatch => {
        dispatch(requestRegistrations(searchQuery))
        return axios.get(`https://5sh4xcm7m8.execute-api.us-west-2.amazonaws.com/Prod/search`, {
            params: {
                q: btoa(`tail_number: ${searchQuery}`),
                limit: 10
            }
        }).then(response => dispatch(receiveRegistrations(searchQuery, response.data)))
    }
}

function shouldFetchRegistrations(state, searchQuery) {
    const registrations = state.registrationsByQuery[searchQuery]
    if (!registrations) {
        return true
    } else {
        return registrations.isFetching
    }
}

export function fetchRegistrationsIfNeeded(q) {
    return (dispatch, getState) => {
        dispatch(searchQuery(q))
        shouldFetchRegistrations(getState(), q) ? dispatch(fetchRegistrations(q)) : Promise.resolve()
    }
}
