import React from 'react';
import '../index.css';
import SearchIcon from '@material-ui/icons/Search';
import ClearIcon from '@material-ui/icons/Clear';
import Suggestions from '../containers/Suggestions.js';
import Filters from '../containers/Filters.js';
import Tray from '../containers/Tray.js';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import LinearProgress from '@material-ui/core/LinearProgress';
import CodeIcon from '@material-ui/icons/Code';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ExpandLessIcon from '@material-ui/icons/ExpandLess';

export default ({ 
    query,
    registrations,
    handleChange,
    onClick,
    isFetching,
    input,
    hasFilters,
    toggleCodeView,
    hideCodeView 
}) => {

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
            <div className="toggleCodeViewRoot">
                <span className="toggleCodeViewIcon" onClick={toggleCodeView}>
                    { hideCodeView ? <ExpandMoreIcon fontSize="small" /> : <ExpandLessIcon fontSize="small" />}
                </span>
                <span className="toggleCodeViewPadding"/>
            </div>
            { hideCodeView ? null : <div className="search-input light no-padding-top no-padding-bottom ">
                    <div className="search">
                        <span className="search-icon light-color">
                            <CodeIcon fontSize="small" />
                        </span>
                        <div className="search-parent">
                            {/* <Filters /> */}
                            <span className="bar">
                                <input
                                    readOnly
                                    className="raw-input light-color"
                                    value={query}
                                    onChange={handleChange}
                                    autoFocus="autofocus"
                                    ref={input}
                                />
                            </span>
                        </div>
                        <span fontSize="small" className="search-icon" />
                    </div>
                </div>
            }
            <div className="search-input">
                <div className="search">
                    <span className="search-icon">
                        <SearchIcon fontSize="small" />
                    </span>
                    <div className="search-parent">
                        <Filters />
                    </div>
                    <span fontSize="small" className="search-icon">
                        {query.length > 0 || hasFilters ? <ClearIcon onClick={onClick} fontSize="inherit"/> : null}
                    </span>
                </div>
            </div>
            <Tray />
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
