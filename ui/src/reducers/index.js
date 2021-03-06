import { combineReducers } from 'redux';
import {
  SEARCH_QUERY,
  REQUEST_REGISTRATIONS,
  RECEIVE_REGISTRATIONS,
  ENABLE_SEARCH_FILTER,
  DISABLE_SEARCH_FILTER,
  HIDE_CODE_VIEW,
  SELECTED_REGISTRATION,
  CLEAR_SELECTED_REGISTRATION,
  CLEAR_SEARCH_FILTERS,
} from '../actions';

function selectedRegistration(state = {}, action) {
  switch (action.type) {
    case SELECTED_REGISTRATION:
      return { ...state, ...action.value };
    case CLEAR_SELECTED_REGISTRATION:
      return {};
    default:
      return state;
  }
}

function searchQuery(state = '', action) {
  switch (action.type) {
    case SEARCH_QUERY:
      return action.searchQuery;
    default:
      return state;
  }
}

function hideCodeView(state = true, action) {
  switch (action.type) {
    case HIDE_CODE_VIEW:
      return action.value;
    default:
      return state;
  }
}

function searchFilters(state = {}, action) {
  switch (action.type) {
    case CLEAR_SEARCH_FILTERS:
      return {};
    case ENABLE_SEARCH_FILTER:
      return {
        ...state,
        [action.when]: {
          key: action.filter,
          value: action.query,
        },
      };
    case DISABLE_SEARCH_FILTER: {
      const { [action.when]: _, ...trimmed } = state;
      return trimmed;
    }
    default:
      return state;
  }
}

function registrations(state = {
  isFetching: false,
  items: [],
}, action) {
  switch (action.type) {
    case REQUEST_REGISTRATIONS:
      return { ...state, isFetching: true };
    case RECEIVE_REGISTRATIONS:
      return {
        ...state,
        isFetching: false,
        items: action.registrations,
        lastUpdated: action.receivedAt,
      };
    default:
      return state;
  }
}

function registrationsByQuery(state = {}, action) {
  switch (action.type) {
    case REQUEST_REGISTRATIONS:
    case RECEIVE_REGISTRATIONS:
      return {
        ...state,
        [action.searchQuery]: registrations(state[action.searchQuery], action),
      };
    default:
      return state;
  }
}

const rootReducer = combineReducers({
  registrationsByQuery,
  searchQuery,
  searchFilters,
  hideCodeView,
  selectedRegistration,
});

export default rootReducer;
