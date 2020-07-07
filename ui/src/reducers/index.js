import { combineReducers } from 'redux'
import { 
    SEARCH_QUERY,
    REQUEST_REGISTRATIONS,
    RECEIVE_REGISTRATIONS,
} from '../actions'

function searchQuery(state = '', action) {
  switch (action.type) {
    case SEARCH_QUERY:
        return action.searchQuery
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
    searchQuery
})

export default rootReducer