import axios from 'axios'

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
                q: btoa(`tail_number: ${searchQuery}`)
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
