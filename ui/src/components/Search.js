import React from 'react';
import PropTypes from 'prop-types';
import '../index.css';
import SearchIcon from '@material-ui/icons/Search';
import ClearIcon from '@material-ui/icons/Clear';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import LinearProgress from '@material-ui/core/LinearProgress';
import CodeIcon from '@material-ui/icons/Code';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ExpandLessIcon from '@material-ui/icons/ExpandLess';
import Suggestions from '../containers/Suggestions';
import Result from '../containers/Result';
import Filters from '../containers/Filters';
import Tray from '../containers/Tray';
import { searchFilters as searchFiltersProp, registration } from '../common/global_types';

export default function Search({
  query,
  handleChange,
  onClick,
  isFetching,
  input,
  hasFilters,
  toggleCodeView,
  hideCodeView,
  searchFilters,
  selectedRegistration,
}) {
  const theme = createMuiTheme({
    overrides: {
      MuiLinearProgress: {
        root: {
          height: '1px',
        },
        colorPrimary: {
          backgroundColor: '#999999',
        },
        barColorPrimary: {
          backgroundColor: '#EEEEEE',
        },
      },
    },
  });

  return (
    <div className="input-and-results">
      { Object.keys(searchFilters).length === 0 ? null : (
        <div className="toggleCodeViewRoot">
          <span role="button" tabIndex={0} onKeyDown={toggleCodeView} className="toggleCodeViewIcon" onClick={toggleCodeView}>
            { hideCodeView ? <ExpandMoreIcon fontSize="small" /> : <ExpandLessIcon fontSize="small" />}
          </span>
          <span className="toggleCodeViewPadding" />
        </div>
      )}
      { hideCodeView ? null : (
        <div className="search-input light no-padding-top no-padding-bottom">
          <div className="search">
            <span className="search-icon light-color">
              <CodeIcon fontSize="small" />
            </span>
            <div className="search-parent no-padding-top">
              <span className="bar">
                <input
                  readOnly
                  className="raw-input light-color"
                  value={query}
                  onChange={handleChange}
                  ref={input}
                />
              </span>
            </div>
            <span fontSize="small" className="search-icon" />
          </div>
        </div>
      )}
      { Object.keys(searchFilters).length === 0 ? null : (
        <div className="search-input">
          <div className="search">
            <div className="filters-item">
              <div className="search-icon-parent foobar">
                <SearchIcon fontSize="small" />
              </div>
            </div>
            <div className="search-parent">
              <Filters />
            </div>
            <span fontSize="small" className="search-icon">
              {query.length > 0 || hasFilters ? <ClearIcon onClick={onClick} fontSize="inherit" /> : null}
            </span>
          </div>
        </div>
      )}
      <Tray />
      {
        isFetching ? (
          <ThemeProvider theme={theme}>
            <LinearProgress color="primary" />
          </ThemeProvider>
        ) : null
      }
      <div className="results-container-parent">
        <div className="results-container">
          { Object.entries(selectedRegistration).length === 0 ? (
            <Suggestions />
          ) : (
            <Result registration={selectedRegistration} />
          )}
        </div>
      </div>
    </div>
  );
}

Search.propTypes = {
  query: PropTypes.string.isRequired,
  handleChange: PropTypes.func.isRequired,
  onClick: PropTypes.func.isRequired,
  isFetching: PropTypes.bool.isRequired,
  input: PropTypes.func.isRequired,
  hasFilters: PropTypes.bool.isRequired,
  toggleCodeView: PropTypes.func.isRequired,
  hideCodeView: PropTypes.bool.isRequired,
  searchFilters: searchFiltersProp.isRequired,
  selectedRegistration: registration.isRequired,
};
