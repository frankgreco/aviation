import axios from 'axios';

export const CLEAR_SELECTED_REGISTRATION = 'CLEAR_SELECTED_REGISTRATION';
export function clearSelectedRegistration() {
  return { type: CLEAR_SELECTED_REGISTRATION };
}

export const SELECTED_REGISTRATION = 'SELECTED_REGISTRATION';
export function selectedRegistration(value) {
  return { type: SELECTED_REGISTRATION, value };
}

export const HIDE_CODE_VIEW = 'HIDE_CODE_VIEW';
export function hideCodeView(value) {
  return { type: HIDE_CODE_VIEW, value };
}

export const ENABLE_SEARCH_FILTER = 'ENABLE_SEARCH_FILTER';
export function enableSearchFilter(filter, query = '') {
  return { type: ENABLE_SEARCH_FILTER, filter, query };
}

export const DISABLE_SEARCH_FILTER = 'DISABLE_SEARCH_FILTER';
export function disableSearchFilter(filter, query = '') {
  return { type: DISABLE_SEARCH_FILTER, filter, query };
}

export const SEARCH_QUERY = 'SEARCH_QUERY';
export function searchQuery(value) {
  return { type: SEARCH_QUERY, searchQuery: value };
}

export const REQUEST_REGISTRATIONS = 'REQUEST_REGISTRATIONS';
export function requestRegistrations(value) {
  return { type: REQUEST_REGISTRATIONS, searchQuery: value };
}

export const RECEIVE_REGISTRATIONS = 'RECEIVE_REGISTRATIONS';
export function receiveRegistrations(value, json) {
  return {
    type: RECEIVE_REGISTRATIONS,
    searchQuery: value,
    registrations: json.items,
    receivedAt: Date.now(),
  };
}

export function fetchRegistrations(value) {
  return (dispatch) => {
    dispatch(requestRegistrations(value));
    return axios.get('https://5sh4xcm7m8.execute-api.us-west-2.amazonaws.com/Prod/search', {
      params: {
        q: btoa(value),
        limit: 10,
      },
    }).then((response) => dispatch(receiveRegistrations(value, response.data)));
  };
}

function shouldFetchRegistrations(state, value) {
  const registrations = state.registrationsByQuery[value];
  if (!registrations) {
    return true;
  }
  return registrations.isFetching;
}

export function fetchRegistrationsIfNeeded(q) {
  return (dispatch, getState) => {
    dispatch(searchQuery(q));
    if (shouldFetchRegistrations(getState(), q)) {
      dispatch(fetchRegistrations(q));
    } else {
      Promise.resolve();
    }
  };
}
