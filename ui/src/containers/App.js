import React from 'react';
import { Provider } from 'react-redux'
import PropTypes from 'prop-types'
import Home from '../components/Home';

const App = ({ store }) => (
    <Provider store={store}>
        <Home />
    </Provider>
)
  
App.propTypes = {
    store: PropTypes.object.isRequired,
}
export default App
  