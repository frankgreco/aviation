import React from 'react';
import '../index.css';
import LabelOutlinedIcon from '@material-ui/icons/LabelOutlined';
import Tag from '../components/Tag.js';
import { allEnabled } from '../utils/timer.js' 

const renderTagIfNeeded = (searchFilters, key, onClick) => searchFilters[key].enabled ? null : <Tag v={key} onClick={onClick(key) } invert={true} />

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
