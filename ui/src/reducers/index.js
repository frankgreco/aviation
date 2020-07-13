import { combineReducers } from 'redux'
import { 
    SEARCH_QUERY,
    REQUEST_REGISTRATIONS,
    RECEIVE_REGISTRATIONS,
    ENABLE_SEARCH_FILTER,
    DISABLE_SEARCH_FILTER,
    HIDE_CODE_VIEW
} from '../actions'

function searchQuery(state = '', action) {
  switch (action.type) {
    case SEARCH_QUERY:
        return action.searchQuery
    default:
        return state
  }
}

function hideCodeView(state= true, action) {
  switch (action.type) {
    case HIDE_CODE_VIEW:
        return action.value
    default:
        return state
  }
}

function searchFilters(state = {
  enabled: false,
  value: ''
}, action) {
  switch (action.type) {
    case ENABLE_SEARCH_FILTER:
      return Object.assign({}, state, {
        [action.filter]: Object.assign({}, state[action.filter], {
          enabled: true,
          value: action.query
        })
      })
    case DISABLE_SEARCH_FILTER:
      return Object.assign({}, state, {
        [action.filter]: Object.assign({}, state[action.filter], {
          enabled: false,
          value: ''
        })
      })
    default:
        return state
  }
}

function registrations(
  state = {
    isFetching: false,
    items: []
  },
  action
) {
  switch (action.type) {
    case REQUEST_REGISTRATIONS:
      return Object.assign({}, state, {
        isFetching: true,
      })
    case RECEIVE_REGISTRATIONS:
      return Object.assign({}, state, {
        isFetching: false,
        items: action.registrations,
        lastUpdated: action.receivedAt
      })
    default:
      return state
  }
}

function registrationsByQuery(state = {}, action) {
  switch (action.type) {
    case REQUEST_REGISTRATIONS:
    case RECEIVE_REGISTRATIONS:
      return Object.assign({}, state, {
        [action.searchQuery]: registrations(state[action.searchQuery], action)
      })
    default:
      return state
  }
}

const rootReducer = combineReducers({
    registrationsByQuery,
    searchQuery,
    searchFilters,
    hideCodeView
})

export default rootReducer