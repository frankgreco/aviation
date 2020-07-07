import React from 'react';
import { render } from 'react-dom';
import App from './containers/App.js';
import thunkMiddleware from 'redux-thunk'
import { createLogger } from 'redux-logger'
import { createStore, applyMiddleware } from 'redux'
import rootReducer from './reducers'

const loggerMiddleware = createLogger()

const store = createStore(
    rootReducer,
    {},
    applyMiddleware(
        thunkMiddleware,
        loggerMiddleware 
    )
)

render(<App store={store}/>, document.getElementById('root'))