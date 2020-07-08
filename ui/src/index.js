import React from 'react';
import { render } from 'react-dom';
import App from './containers/App.js';
import thunkMiddleware from 'redux-thunk'
import { createLogger } from 'redux-logger'
import { createStore, applyMiddleware } from 'redux'
import rootReducer from './reducers'

const store = createStore(
    rootReducer,
    {},
    process.env.NODE_ENV === 'production' ? applyMiddleware(
        thunkMiddleware
    ) : applyMiddleware(
        thunkMiddleware,
        createLogger()
    )
)

render(<App store={store}/>, document.getElementById('root'))