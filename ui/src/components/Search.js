import React from 'react';
import '../index.css';
import SearchIcon from '@material-ui/icons/Search';
import ClearIcon from '@material-ui/icons/Clear';
import Suggestions from '../containers/Suggestions.js';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import LinearProgress from '@material-ui/core/LinearProgress';
import Tag from './Tag.js';

export default ({ query, registrations, handleChange, onClick, isFetching, input }) => {

    const theme = createMuiTheme({
        overrides: {
            MuiLinearProgress: {
                root: {
                    height: '1px'
                },
                colorPrimary: {
                    backgroundColor: '#999999'
                },
                barColorPrimary: {
                    backgroundColor: '#EEEEEE'
                }
            },
        },
    });

    return (
        <div className="input-and-results">
            <div className="search-input">
                <div className="search">
                    <span className="search-icon">
                        <SearchIcon/>
                    </span>
                    <span className="bar">
                        <input
                            placeholder="Search by tail number..."
                            value={query}
                            onChange={handleChange}
                            autoFocus="autofocus"
                            ref={input}
                        />
                    </span>
                    <span fontSize="small" className="search-icon">
                        {query.length > 0 ? <ClearIcon onClick={onClick} fontSize="inherit"/> : null}
                    </span>
                </div>
            </div>
            {
                query.length > 0 ? 
                    <div className="filters-parent">
                        <div className="filters-item" />
                        <div className="filters">
                            <Tag v="tail number" invert={true} />
                            <Tag v="make" invert={true}/>
                            <Tag v="model" invert={true}/>
                            <Tag v="airline" invert={true}/>
                        </div>
                        <div className="filters-item"/>
                    </div>
                : null
            }
            {
                isFetching ? 
                    <ThemeProvider theme={theme}>
                        <LinearProgress color="primary"/>
                    </ThemeProvider>
                : null
            }
            <div className="seperator-parent">
                { registrations !== undefined && registrations.length > 0  && !isFetching ? <div className="seperator" /> : null }
            </div>
            <div className="results-container-parent">
                <div className="results-container">
                    <Suggestions />
                </div>
            </div>
        </div>
    )
}
