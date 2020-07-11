import React from 'react';
import '../index.css';
import LabelOutlinedIcon from '@material-ui/icons/LabelOutlined';
import Tag from '../components/Tag.js';

const renderTagIfNeeded = (searchFilters, key, onClick) => searchFilters[key].enabled ? null : <Tag v={key} onClick={onClick(key) } invert={true} />

const allEnabled = searchFilters => Object.keys(searchFilters).filter(k => !searchFilters[k].enabled).length === 0

export default ({ onClick, searchFilters }) => allEnabled(searchFilters) ? null : 
    <div className="filters-parent">
        <div className="filters-item">
            <LabelOutlinedIcon fontSize="small" />
        </div>
        <div className="filters">
            {renderTagIfNeeded(searchFilters, 'tail number', onClick)}
            {renderTagIfNeeded(searchFilters, 'make', onClick)}
            {renderTagIfNeeded(searchFilters, 'model', onClick)}
            {renderTagIfNeeded(searchFilters, 'airline', onClick)}
        </div>
        <div className="filters-item"/>
    </div>
